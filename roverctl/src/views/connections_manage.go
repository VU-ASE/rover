package views

import (
	"github.com/VU-ASE/roverctl/src/configuration"
	"github.com/VU-ASE/roverctl/src/state"
	"github.com/VU-ASE/roverctl/src/style"
	"github.com/VU-ASE/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ConnectionsManagePage struct {
	selectedIndex int
}

func NewConnectionsManagePage() ConnectionsManagePage {
	return ConnectionsManagePage{
		selectedIndex: 0,
	}
}

func (m ConnectionsManagePage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys().Up):
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case key.Matches(msg, m.keys().Down):
			if m.selectedIndex < len(m.GetNonActiveConnections())-1 {
				m.selectedIndex++
			}
		case key.Matches(msg, m.keys().Confirm):
			items := m.GetNonActiveConnections()
			if m.selectedIndex >= 0 && m.selectedIndex < len(items) {
				state.Get().RoverConnections.Active = items[m.selectedIndex].Name
			}
			return m, nil
		}
	}

	// switch _ = msg.(type) {

	// case tea.KeyMsg:
	// 	switch {
	// 	// case key.Matches(msg, connectionsManageKeys.MarkActive):

	// 	// 	}
	// 	// case key.Matches(msg, connectionsManageKeys.New):
	// 	// 	return RootScreen(state.Get()).SwitchScreen(NewConnectionsInitPage(nil))
	// 	// case key.Matches(msg, connectionsManageKeys.Delete):
	// 	// 	if len(m.list.Items()) > 1 && m.list.Index() >= 0 && m.list.Index() < len(m.list.Items()) {
	// 	// 		item := m.list.Items()[m.list.Index()].(item)
	// 	// 		state.Get().RoverConnections = state.Get().RoverConnections.Remove(item.connection.Name)
	// 	// 		m.list.SetItems(connectionsToListItems())
	// 	// 		m.list.ResetSelected()
	// 	// 		return m, nil
	// 	// 	}
	// 	// }
	// 	}
	// }

	return m, nil
}

func (m ConnectionsManagePage) Init() tea.Cmd {
	return nil
}

func (m ConnectionsManagePage) View() string {
	s := style.Title.Render("Manage connections") + "\n\n"

	connections := state.Get().RoverConnections.Available

	if connections == nil {
		s += style.Gray.Render("No connections available")
		return s
	}

	active := m.GetActiveConnection()
	if active != nil {
		s += "Active:\n"
		s += "   " + style.Gray.Render(active.Name+"\n"+"   "+active.Host) + "\n"
	} else {
		s += style.Gray.Render("No active connection")
	}

	available := m.GetNonActiveConnections()
	if len(available) > 0 {
		s += "\n" + ("Available") + ":\n"

		for i, connection := range available {
			if connection.Name == active.Name {
				continue
			}

			item := connection.Name + " " + connection.Host

			if i == m.selectedIndex {
				s += style.Primary.Render(" > " + item)
			} else {
				s += style.Gray.Render(" • " + item)
			}
			s += "\n"
		}
	}

	return s
}

func (m ConnectionsManagePage) GetActiveConnection() *configuration.RoverConnection {
	connections := state.Get().RoverConnections.Available
	activeName := state.Get().RoverConnections.Active

	for _, connection := range connections {
		if connection.Name == activeName {
			return &connection
		}
	}

	return nil
}

func (m ConnectionsManagePage) GetNonActiveConnections() []configuration.RoverConnection {
	connections := state.Get().RoverConnections.Available
	activeName := state.Get().RoverConnections.Active

	nonActive := []configuration.RoverConnection{}

	for _, connection := range connections {
		if connection.Name != activeName {
			nonActive = append(nonActive, connection)
		}
	}

	return nonActive
}

func (m ConnectionsManagePage) isQuitable() bool {
	return true
}

func (m ConnectionsManagePage) keys() utils.GeneralKeyMap {
	kb := utils.NewGeneralKeyMap()
	items := m.GetNonActiveConnections()
	if len(items) > 1 {
		kb.Up = key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		)
		kb.Down = key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		)
		kb.Confirm = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "set active"),
		)
	} else if len(items) == 1 {
		kb.Confirm = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "change active"),
		)
	}
	return kb
}

func (m ConnectionsManagePage) previousPage() *pageModel {
	var pageModel pageModel = NewStartPage()
	return &pageModel
}
