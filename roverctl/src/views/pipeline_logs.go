package views

import (
	"context"
	"fmt"
	"strings"

	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

//
// All keys
//

// Keys to navigate
type PipelineLogsKeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Retry   key.Binding
	Confirm key.Binding
	Quit    key.Binding
}

// Shown when the services are being updated
var PipelineLogsKeysRegular = PipelineLogsKeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "down"),
	),
	Retry: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refetch"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

func (k PipelineLogsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Retry, k.Confirm, k.Quit}
}

func (k PipelineLogsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

//
// The page model
//

type PipelineLogsPage struct {
	help     help.Model
	spinner  spinner.Model
	viewport viewport.Model
	logs     tui.Action[[]string]
	// Service FQN
	service string // name
	author  string
	version string
}

func NewPipelineLogsPage(service string, author string, version string) PipelineLogsPage {
	return PipelineLogsPage{
		spinner: spinner.New(),
		help:    help.New(),
		logs:    tui.NewAction[[]string]("fetchLogs"),
		viewport: viewport.Model{
			Width:  state.Get().WindowWidth,
			Height: state.Get().WindowHeight - 4,
		},
		service: service,
		author:  author,
		version: version,
	}
}

//
// Page model methods
//

func (m PipelineLogsPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tui.ActionInit[[]string]:
		m.logs.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[[]string]:
		m.logs.ProcessResult(msg)
		if m.logs.IsSuccess() {
			m.viewport.YOffset = 100
			m.viewport.SetContent(strings.Join(*m.logs.Data, "\n"))
		}
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, PipelineLogsKeysRegular.Quit):
			return m, tea.Quit
		case key.Matches(msg, PipelineLogsKeysRegular.Retry):
			if !m.logs.IsLoading() {
				return m, m.fetchLogs()
			}
		case key.Matches(msg, PipelineLogsKeysRegular.Confirm):
			// todo:
			return m, nil
		}
	case tea.WindowSizeMsg:
		offset := m.viewport.YOffset
		m.viewport = viewport.New(msg.Width, msg.Height-2-2-2) // -2 for the header, -2 for the padding, -2 for the help
		if m.logs.IsSuccess() {
			m.viewport.YOffset = offset
			m.viewport.SetContent(strings.Join(*m.logs.Data, "\n"))
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m PipelineLogsPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchLogs())
}

func (m PipelineLogsPage) View() string {
	// We use hasData, so that we can do optimistic refetching, without blocking the entire view
	if m.logs.IsSuccess() || m.logs.HasData() {
		s := ""
		if m.logs.IsLoading() && m.logs.HasData() { // if we are loading optimistically, show the spinner so that there is still an indication that something is happening
			s = " " + m.spinner.View()
		}
		return style.Title.Render("Logs for "+m.author+"/"+m.service+":"+m.version) + style.Gray.Render(fmt.Sprintf(" %3.f%%", 100.0*m.viewport.ScrollPercent())) + s + "\n\n" + m.viewport.View() + "\n\n" + m.help.View(PipelineLogsKeysRegular)
	} else if m.logs.IsError() {
		return style.Error.Render("Failed to load logs for " + m.service + ".")
	} else {
		return m.spinner.View() + " Fetching logs..."
	}
}

func (m PipelineLogsPage) fetchLogs() tea.Cmd {
	return tui.PerformAction(&m.logs, func() (*[]string, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.PipelineAPI.LogsAuthorNameVersionGet(
			context.Background(),
			m.author,
			m.service,
			m.version,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return &res, err
	})
}

func (m PipelineLogsPage) isQuitable() bool {
	return true
}

func (m PipelineLogsPage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m PipelineLogsPage) previousPage() *pageModel {
	var pageModel pageModel = NewPipelineManagerPage()
	return &pageModel
}
