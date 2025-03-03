package views

import (
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

//
// The page model
//

type ServiceSyncInstructionsPage struct {
	spinner spinner.Model
}

func NewServiceSyncInstructionsPage() ServiceSyncInstructionsPage {
	// todo

	return ServiceSyncInstructionsPage{
		spinner: spinner.New(),
	}
}

//
// Page model methods
//

func (m ServiceSyncInstructionsPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	}

	return m, nil
}

func (m ServiceSyncInstructionsPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick)
}

func (m ServiceSyncInstructionsPage) View() string {
	s := style.Title.Render("Uploading services") + "\n\n"

	s += "By default, roverctl will look at the current directory for a " + style.Gray.Render("service.yaml") + " file. But none was found.\n"
	s += "To upload one or more services, you can run:\n\n"
	s += style.Primary.Render("   roverctl upload <paths> [--watch]") + "\n\n"
	s += "Example usage:\n\n"
	s += style.Gray.Render("   Upload current working directory") + "\n"
	s += style.Primary.Render("   roverctl upload .") + "\n\n"

	s += style.Gray.Render("   Upload multiple directories") + "\n"
	s += style.Primary.Render("   roverctl upload ./controller ./imaging") + "\n\n"

	s += style.Gray.Render("   Upload multiple directories with file watching (hot reload)") + "\n"
	s += style.Primary.Render("   roverctl upload ./controller ./imaging --watch") + "\n\n"

	return s
}

func (m ServiceSyncInstructionsPage) isQuitable() bool {
	return true
}

func (m ServiceSyncInstructionsPage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m ServiceSyncInstructionsPage) previousPage() *tea.Model {
	return nil
}
