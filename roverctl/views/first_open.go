package views

// import (
// 	"fmt"
// 	"regexp"

// 	"github.com/VU-ASE/rover/roverctl/src/state"
// 	"github.com/VU-ASE/rover/roverctl/src/style"
// 	"github.com/VU-ASE/rover/roverctl/src/utils"
// 	"github.com/charmbracelet/bubbles/key"
// 	"github.com/charmbracelet/bubbles/spinner"
// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// )

// //
// // The page model
// //

// type FirstOpenPage struct {
// 	spinner     spinner.Model
// 	authorInput textinput.Model
// 	error       error
// }

// func NewFirstOpenPage() FirstOpenPage {
// 	ti := textinput.New()
// 	ti.Placeholder = "Start typing..."
// 	ti.Focus()
// 	ti.CharLimit = 156
// 	ti.Width = 100

// 	return FirstOpenPage{
// 		spinner:     spinner.New(),
// 		authorInput: ti,
// 	}
// }

// //
// // Page model methods
// //

// func (m FirstOpenPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	switch msg := msg.(type) {
// 	case spinner.TickMsg:
// 		m.spinner, cmd = m.spinner.Update(msg)
// 		return m, cmd
// 	case tea.KeyMsg:
// 		switch {
// 		case key.Matches(msg, m.keys().Confirm):
// 			// Can only contain lowercase letters and hyphens
// 			valid := regexp.MustCompile(`^[a-z0-9-]*$`).MatchString(m.authorInput.Value())
// 			if !valid {
// 				m.error = fmt.Errorf("Author names can only contain lowercase letters and hyphens")
// 			} else if len(m.authorInput.Value()) < 3 {
// 				m.error = fmt.Errorf("Author name must be at least 3 characters long")
// 			} else {
// 				s := state.Get()
// 				s.Config.Author = m.authorInput.Value()
// 				return RootScreen(s).SwitchScreen(NewStartPage())
// 			}
// 			return m, nil
// 		}
// 		m.error = nil
// 	}

// 	m.authorInput, cmd = m.authorInput.Update(msg)
// 	return m, cmd
// }

// func (m FirstOpenPage) Init() tea.Cmd {
// 	return tea.Batch(m.spinner.Tick, textinput.Blink)
// }

// func (m FirstOpenPage) View() string {
// 	s := style.Success.Bold(true).Render("Welcome to Roverctl!") +
// 		"\n\nEnter your author name to get started.\n" +
// 		"(It will be attached to all code you write)\n\n"
// 	s += m.authorInput.View()

// 	if m.error != nil {
// 		s += "\n\n" + style.Error.Render("✘ "+m.error.Error())
// 	}

// 	return s
// }

// func (m FirstOpenPage) isQuitable() bool {
// 	return true
// }

// func (m FirstOpenPage) keys() utils.GeneralKeyMap {
// 	kb := utils.NewGeneralKeyMap()

// 	if len(m.authorInput.Value()) > 3 {
// 		kb.Confirm = key.NewBinding(
// 			key.WithKeys("enter"),
// 			key.WithHelp("enter", "confirm"),
// 		)
// 	}
// 	kb.Back.SetEnabled(false)
// 	return kb
// }

// func (m FirstOpenPage) previousPage() *tea.Model {
// 	return nil
// }
