package views

import (
	"github.com/VU-ASE/roverctl/src/style"
	"github.com/VU-ASE/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

//
// The page model
//

type TemplatePage struct {
	spinner spinner.Model
}

func NewTemplatePage() TemplatePage {
	// todo

	return TemplatePage{
		spinner: spinner.New(),
	}
}

//
// Page model methods
//

func (m TemplatePage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
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

func (m TemplatePage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick)
}

func (m TemplatePage) View() string {
	return style.Title.Render("Template page")
}

func (m TemplatePage) isQuitable() bool {
	return true
}

func (m TemplatePage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m TemplatePage) previousPage() *pageModel {
	return nil
}
