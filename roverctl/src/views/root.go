package views

import (
	"os"

	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pageModel interface {
	// implement everything from tea.Model
	Init() tea.Cmd                           // Initialize the model
	Update(msg tea.Msg) (pageModel, tea.Cmd) // Update the model based on a message
	View() string                            // Render the model

	isQuitable() bool          // can be quit on "q" entry
	keys() utils.GeneralKeyMap // to be shown at the bottom of the screen, nil if no help is available
	previousPage() *pageModel  // if "back" is pressed, return to the previous page. nil if no previous page is available (will quit)
}

// This is the main model that will be used to render all sub-models
type MainModel struct {
	help    help.Model
	current pageModel
	cli     bool
}

func RootScreen(s *state.AppState) MainModel {
	var start pageModel
	start = NewStartPage()

	// Roverctl was never initialized before, force the user to pick an author name
	if s.Config.Author == "" {
		start = NewFirstOpenPage()
	}

	h := help.New()
	h.Styles.ShortKey = lipgloss.NewStyle().Foreground(style.AsePrimary).Bold(true)
	h.Styles.ShortKey = h.Styles.ShortKey.Bold(true)

	return MainModel{
		current: start, // needs to be a pointer so that the model state can be modified (see https://shi.foo/weblog/multi-view-interfaces-in-bubble-tea)
		help:    h,
		cli:     false,
	}
}

func CliRootScreen(s *state.AppState, page pageModel) MainModel {
	h := help.New()
	h.Styles.ShortKey = lipgloss.NewStyle().Foreground(style.AsePrimary).Bold(true)
	h.Styles.ShortKey = h.Styles.ShortKey.Bold(true)

	return MainModel{
		current: page, // needs to be a pointer so that the model state can be modified (see https://shi.foo/weblog/multi-view-interfaces-in-bubble-tea)
		help:    h,
		cli:     true,
	}
}

func (m MainModel) Init() tea.Cmd {
	return m.current.Init()
}

func (m MainModel) View() string {
	bg := style.AsePrimary
	if state.Get().VersionMismatch {
		bg = style.WarningPrimary
	}

	s := state.Get()
	// Define the header style
	headerStyle := lipgloss.NewStyle().
		Width(s.WindowWidth).   // Set the width of the header to the window width
		Align(lipgloss.Center). // Center-align the text
		Background(bg)          // Set the background color

	con := ""
	if s.RoverConnections.Active != "" {
		con = " | " + s.RoverConnections.Active
	}

	header := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true).Padding(0, 0).Render("VU ASE") + lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(bg).Bold(false).Padding(0, 0).Render(", "+s.Quote+" | "+utils.Version(version)+con)
	fullScreen := lipgloss.NewStyle().Padding(1, 2).Width(s.WindowWidth).Height(s.WindowHeight - 3) // leave room for the header and help keys

	keys := m.current.keys()
	if m.cli {
		// Nothing to go back to
		keys.Back.SetEnabled(false)
	}
	helpView := m.help.View(keys)

	if state.Get().Interactive {
		return fullScreen.Render(m.current.View()) + "\n" + " " + helpView + "\n\n" + headerStyle.Render(header)
	} else {
		return m.current.View()
	}
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle global messages first
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Store the window dimensions
		state.Get().WindowWidth = msg.Width
		state.Get().WindowHeight = msg.Height

		passedMsg := tea.WindowSizeMsg{
			Width:  msg.Width,
			Height: msg.Height - 4, // leave room for the header
		}

		// Forward the message to the current sub-model
		updatedModel, cmd := m.current.Update(passedMsg)
		m.current = updatedModel
		return m, cmd
	case tea.KeyMsg:
		switch {
		// Back
		case key.Matches(msg, m.current.keys().Back), msg.String() == "ctrl+b":
			// These pages expect user input, which might contain a "q", so don't quit if they say that they are not quitable
			if !m.current.isQuitable() && msg.String() != "ctrl+b" {
				// Delegate other messages to the current sub-model
				updatedModel, cmd := m.current.Update(msg)
				m.current = updatedModel
				return m, cmd
			}

			// Roverctl was not "opened" from another screen
			argv := os.Args[1:]
			if len(argv) > 0 {
				return m, tea.Quit
			}

			// Return to a route based on the current route
			returnTo := m.current.previousPage()
			if returnTo == nil {
				return m, tea.Quit
			}

			var cmd tea.Cmd
			m.current, cmd = RootScreen(state.Get()).SwitchScreen(*returnTo)
			return m, cmd
		// Hard quit
		case msg.String() == "ctrl+c", msg.String() == "esc":
			return m, tea.Quit
		}
	}

	// Delegate other messages to the current sub-model
	updatedModel, cmd := m.current.Update(msg)
	m.current = updatedModel
	return m, cmd
}

// This function is used to switch between screens, the caller should supply the "route" taken so far to get to this screen, so that users can return to the previous screen
func (m MainModel) SwitchScreen(model pageModel) (pageModel, tea.Cmd) {
	m.current = model

	// Roverctl was never initialized before, force the user to pick an author name
	if state.Get().Config.Author == "" {
		m.current = NewFirstOpenPage()
	}

	// Notify the new model of the current window size
	windowSizeMsg := tea.WindowSizeMsg{
		Width:  state.Get().WindowWidth,
		Height: state.Get().WindowHeight,
	}

	// Initialize the new model and send the size message
	initCmd := m.current.Init()
	sizeCmd := func() tea.Cmd {
		return func() tea.Msg {
			return windowSizeMsg
		}
	}

	return m.current, tea.Sequence(initCmd, sizeCmd())
}
