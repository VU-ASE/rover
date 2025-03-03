package views

import (
	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

//
// The page model
//

type ShutdownRoverPage struct {
	spinner spinner.Model

	// action
	shutdown tui.Action[bool]
}

func NewShutdownRoverPage() ShutdownRoverPage {
	s := tui.NewAction[bool]("installed")
	return ShutdownRoverPage{
		spinner:  spinner.New(),
		shutdown: s,
	}
}

//
// Page model methods
//

func (m ShutdownRoverPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys().Back):
			return m, tea.Quit
		}
	case tui.ActionInit[bool]:
		m.shutdown.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[bool]:
		m.shutdown.ProcessResult(msg)
		return m, nil
	}

	return m, nil
}

func (m ShutdownRoverPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.shutdownRover())
}

func (m ShutdownRoverPage) View() string {
	s := style.Title.Render("Download default pipeline page") + "\n\n"

	if m.shutdown.IsError() {
		s += style.Error.Render("✗ Could not shutdown rover") + style.Gray.Render(" ("+m.shutdown.Error.Error()+")") + "\n\n"
	} else if m.shutdown.IsSuccess() {
		s += style.Success.Render("✓ Rover was shut down") + "\n\n"
		s += "    It might take a few seconds for the rover to fully shut down." + "\n"
		s += style.Warning.Render("    Once shut down, unplug battery power IMMEDIATELY") + "\n\n"

	} else {
		s += m.spinner.View() + " Shutting down rover..." + "\n\n"
	}

	return s
}

func (m ShutdownRoverPage) isQuitable() bool {
	return true
}

func (m ShutdownRoverPage) keys() utils.GeneralKeyMap {
	kb := utils.NewGeneralKeyMap()
	if m.shutdown.IsLoading() {
		kb.Back.SetEnabled(false)
	}
	return kb
}

func (m ShutdownRoverPage) previousPage() *tea.Model {
	var tea.Model tea.Model = NewStartPage()
	return &tea.Model
}

//
// Actions
//

func (m ShutdownRoverPage) shutdownRover() tea.Cmd {
	return tui.PerformAction(&m.shutdown, func() (*bool, error) {
		// remote := state.Get().RoverConnections.GetActive()
		// if remote == nil {
		// 	return nil, fmt.Errorf("No active rover connection")
		// }

		// // First, save the pipeline
		// api := remote.ToApiClient()
		// req := api.HealthAPI.ShutdownPost(
		// 	context.Background(),
		// )

		// htt, err := req.Execute()
		// if err != nil {
		// 	return nil, utils.ParseHTTPError(err, htt)
		// }

		return openapi.PtrBool(true), nil
	})
}
