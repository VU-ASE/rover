package view_upload

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/radovskyb/watcher"
)

var mutex sync.Mutex = sync.Mutex{}

type model struct {
	rover         configuration.RoverConnection                        // The rover to upload to
	watch         bool                                                 // Whether to watch for changes
	isServiceDir  bool                                                 // Is the cwd a service directory?
	uploading     map[string]*tui.Action[openapi.FetchPost200Response] // map of file paths to upload responses
	watchers      map[string]*tui.Action[bool]                         // map of file paths to watcher actions
	services      map[string]utils.ServiceInformation                  // map of file paths to service information
	paths         []string                                             // because maps don't preserve order
	watchDebounce time.Duration                                        // debounce time for collecting changes
	spinner       spinner.Model
	cwd           string
}

func New(rover configuration.RoverConnection, paths []string, watch bool) model {
	// Actions
	uploading := make(map[string]*tui.Action[openapi.FetchPost200Response])
	services := make(map[string]utils.ServiceInformation)
	watchers := make(map[string]*tui.Action[bool])

	for _, path := range paths {
		act := tui.NewAction[openapi.FetchPost200Response](path)
		uploading[path] = &act

		// Try to parse the yaml for this path
		info, err := utils.GetServiceInformation(path)
		if err != nil || info.Name == "" || info.Version == "" {
			continue
		}
		services[path] = *info

		// Create a watcher for this path
		watchAct := tui.NewAction[bool](path)
		watchers[path] = &watchAct
	}

	// Get cwd for the watcher
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	sp := spinner.New()
	model := model{
		isServiceDir:  true,
		watchDebounce: 500 * time.Millisecond,
		spinner:       sp,
		uploading:     uploading,
		services:      services,
		cwd:           cwd,
		paths:         paths,
		watchers:      watchers,
		watch:         watch,
		rover:         rover,
	}

	return model
}

func (m model) Init() tea.Cmd {
	sequence := make([]tea.Cmd, 0)
	sequence = append(sequence, m.spinner.Tick)
	for path := range m.uploading {
		sequence = append(sequence, m.uploadChanges(path))
	}

	return tea.Sequence(
		sequence...,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tui.ActionInit[openapi.FetchPost200Response]:
		if _, ok := m.uploading[msg.Name]; !ok {
			return m, nil
		}

		act := m.uploading[msg.Name]
		act.ProcessInit(msg)

		return m, nil
	case tui.ActionResult[openapi.FetchPost200Response]:
		if _, ok := m.uploading[msg.Name]; !ok {
			return m, nil
		}

		act := m.uploading[msg.Name]
		act.ProcessResult(msg)

		if m.allUploadsDone() && !m.watch {
			return m, tea.Quit
		}

		// Start watcher
		if act.IsSuccess() && m.watch {
			return m, m.watchChanges(msg.Name)
		}

		return m, nil
	case tui.ActionInit[bool]:
		if _, ok := m.watchers[msg.Name]; !ok {
			return m, nil
		}

		act := m.watchers[msg.Name]
		act.ProcessInit(msg)

		return m, nil
	case tui.ActionResult[bool]:
		if _, ok := m.watchers[msg.Name]; !ok {
			return m, nil
		}

		act := m.watchers[msg.Name]
		act.ProcessResult(msg)

		// Find the uploader action
		uploader, ok := m.uploading[msg.Name]
		if !ok {
			return m, nil
		}

		if act.IsSuccess() && uploader.IsDone() {
			return m, m.uploadChanges(msg.Name)
		}

		return m, nil
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	st := state.Get()
	s := ""

	if st.Config.Author == "" {
		s += "Uh oh, " + style.Error.Render("roverctl does not know who you are yet") + "!\n"
		s += "To get started, just run " + style.Primary.Render("roverctl author --set <NAME>") + " and try again.\n"
		return s
	}

	for _, path := range m.paths {
		act, ok := m.uploading[path]
		if !ok {
			continue
		}

		pathStr := path
		if pathStr == "." {
			pathStr = m.cwd
		}
		pathStr = strings.TrimRight(pathStr, "/")

		// Is this service defined?
		info, ok := m.services[path]

		if !ok {
			s += (pathStr) + " -> " + m.rover.Identifier + style.Gray.Render(" (unknown)") + "\n"
			s += style.Warning.Render("✗ No valid service.yaml file found in this directory") + "\n" + style.Gray.Render("Check out ") + style.Primary.Render("https://ase.vu.nl/docs/framework/glossary/service") + style.Gray.Render(" to understand service.yaml requirements") + "\n\n"
			continue
		}
		s += (pathStr) + " -> " + (m.rover.Identifier + style.Gray.Render(" (/home/debix/.rover/"+st.Config.Author+"/"+info.Name+"/"+info.Version+")")) + "\n"

		if act.IsLoading() {
			s += style.Primary.Render(m.spinner.View()+" Uploading...") + "\n\n"
		} else if act.IsError() {
			msg := ""
			if act.Error != nil {
				msg = act.Error.Error()
				// Remove trailing newline (if there)
				msg = strings.TrimRight(msg, "\n")
			}

			s += style.Error.Render("✗ Upload failed") + style.Gray.Render(" ("+msg+")") + "\n\n"
		} else if act.IsSuccess() {
			// Is there a watcher loading?
			msg := ""
			watcher, ok := m.watchers[path]
			if ok && watcher.IsLoading() {
				msg = " " + m.spinner.View() + " watching for new changes to upload..."
			}

			s += style.Success.Render("✓ Upload successful") + msg + "\n\n"
		}
	}

	return s
}

func createZipFromDirectory(zipPath, sourceDir string) error {
	// Create the zip file
	tmpZip, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create temp zip file: %v", err)
	}
	defer tmpZip.Close()

	// Create a zip writer
	zipWriter := zip.NewWriter(tmpZip)
	defer func() {
		if closeErr := zipWriter.Close(); closeErr != nil {
			fmt.Printf("Error closing zip writer: %v\n", closeErr)
		}
	}()

	// Walk through the source directory
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking the path %q: %v", path, err)
		}

		// Get the relative path to store in the zip archive
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %v", err)
		}

		// Skip the root directory itself
		if relPath == "." {
			return nil
		}

		// Skip the "roverctl" binary
		if relPath == "roverctl" {
			return nil
		}

		if info.IsDir() {
			// Add a directory entry to the zip file
			_, err := zipWriter.Create(relPath + "/")
			if err != nil {
				return fmt.Errorf("failed to create directory entry: %v", err)
			}
			return nil
		}

		// Add a file entry to the zip file
		fileWriter, err := zipWriter.Create(relPath)
		if err != nil {
			return fmt.Errorf("failed to create file entry: %v", err)
		}

		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %v", err)
		}
		defer file.Close()

		// Copy the file content to the zip file
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return fmt.Errorf("failed to write file content: %v", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Upload all collected changes to the Rover
func (m model) uploadChanges(path string) tea.Cmd {
	act := m.uploading[path]
	return tui.PerformAction(act, func() (*openapi.FetchPost200Response, error) {
		mutex.Lock()
		defer mutex.Unlock()
		sourceDir := path

		// Copy all files to a temp directory
		copyDir, err := os.MkdirTemp("", "rover-sync-"+time.Now().GoString())
		if err != nil {
			return nil, fmt.Errorf("Error creating temp dir: %v\n", err)
		}
		defer os.RemoveAll(copyDir)
		err = copy(sourceDir, copyDir)
		if err != nil {
			return nil, fmt.Errorf("Error copying files: %v\n", err)
		}
		// Replace the author field in service.yaml
		err = replaceAuthor(filepath.Join(copyDir, "service.yaml"), state.Get().Config.Author)
		if err != nil {
			return nil, fmt.Errorf("Error replacing author: %v\n", err)
		}

		// Create a temp zip file
		zipPath := os.TempDir() + "/rover-sync-" + time.Now().Format("20060102150405") + ".zip"

		err = createZipFromDirectory(zipPath, copyDir)
		if err != nil {
			return nil, fmt.Errorf("Error creating zip: %v\n", err)
		}

		// Open the zip file
		zipFile, err := os.Open(zipPath)
		if err != nil {
			return nil, fmt.Errorf("Failed to open temp zip file: %v", err)
		}
		defer zipFile.Close()

		// Upload the zip file
		api := m.rover.ToApiClient()
		req := api.ServicesAPI.UploadPost(
			context.Background(),
		)
		req = req.Content(zipFile)

		// req.Content(zipFile)
		res, htt, err := req.Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return res, err
	})
}

// Collect changes for a path, report true if there are changes
// Upload all collected changes to the Rover
func (m model) watchChanges(path string) tea.Cmd {
	act, ok := m.watchers[path]
	if !ok {
		return nil
	}

	return tui.PerformAction(act, func() (*bool, error) {
		w := watcher.New()

		// Only notify rename and move events
		w.FilterOps(watcher.Rename, watcher.Move, watcher.Create, watcher.Remove, watcher.Write)

		// Ignore .git directory
		err := w.Ignore("./.git")
		if err != nil {
			return nil, err
		}

		// Watch recursively for changes
		err = w.AddRecursive(path)
		if err != nil {
			return nil, err
		}

		changed := false
		// Goroutine to monitor events and close the watcher after the first event
		go func() {
			for {
				select {
				case <-w.Event:
					w.Close()
					changed = true
					return
				case <-w.Error:
					w.Close()
					return
				case <-w.Closed:
					return
				}
			}
		}()

		// Start the watching process - check for changes every 500ms
		err = w.Start(time.Millisecond * 500)
		if err != nil {
			return nil, err
		}
		return &changed, nil
	})
}

// Reports if all uploads are done (i.e. successful or failed, but not loading anymore)
func (m model) allUploadsDone() bool {
	for _, act := range m.uploading {
		if !act.IsDone() {
			return false
		}
	}

	return true
}

func (m model) atLeastOneUploadSuccess() bool {
	for _, act := range m.uploading {
		if act.IsSuccess() {
			return true
		}
	}

	return false
}

// Copy recursively copies files and directories from src to dst
func copy(src, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Determine destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		// Create directories if necessary
		if info.IsDir() {
			return os.MkdirAll(dstPath, os.ModePerm)
		}

		// Copy files
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}

// Process service.yaml to replace "author" field with given author
func replaceAuthor(filePath, author string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	found := false

	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "author:") {
			lines[i] = fmt.Sprintf("author: %s", author)
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, fmt.Sprintf("author: %s", author))
	}

	return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), os.ModePerm)
}
