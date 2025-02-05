package views

import (
	"fmt"
	"strings"
	"time"

	"github.com/VU-ASE/roverctl/src/state"
	"github.com/VU-ASE/roverctl/src/style"
	"github.com/VU-ASE/roverctl/src/tui"
	"github.com/VU-ASE/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StartPage struct {
	// To select an action to perform with this utility
	help      help.Model // to display a help footer
	listItems []Category // list of lists (one list per category)
	listIndex int        // index of the current list
	itemIndex int        // index of the current item in the current list

	filterInput textinput.Model

	forceOnline bool // override the rover connection state

	// Actions
	updateAvailable tui.Action[utils.UpdateAvailable] // preserved in the model to avoid re-rendering in .View(), based on the latest service information
	roverOnline     tui.ActionV2[any, bool]
}

type CategoryItem struct {
	label string
	key   string
}

type Category struct {
	name  string
	kind  string
	items []CategoryItem
}

func NewStartPage() StartPage {
	ti := textinput.New()
	ti.Placeholder = "Type to filter..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 100

	return StartPage{
		listItems:   initialCategories(),
		help:        help.New(),
		filterInput: ti,
		roverOnline: tui.NewActionV2[any, bool](),
		forceOnline: false,
	}
}

func (m StartPage) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.checkUpdate(), m.checkRoverOnline(false))
}

func (m StartPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tui.ActionUpdate[any, bool]:
		m.roverOnline.ProcessUpdate(msg)
		if m.roverOnline.IsDone() {
			return m, m.checkRoverOnline(true) // keep checking
		}
		return m, nil
	case tui.ActionInit[utils.UpdateAvailable]:
		m.updateAvailable.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[utils.UpdateAvailable]:
		m.updateAvailable.ProcessResult(msg)
		return m, nil
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch {
		case key.Matches(msg, m.keys().Confirm):
			value := m.getSelectedItem().key
			if value != "" {
				switch value {
				case "pipeline-manage":
					return RootScreen(state.Get()).SwitchScreen(NewPipelineManagerPage())
				case "pipeline-configure":
					return RootScreen(state.Get()).SwitchScreen(NewPipelineConfiguratorPage())
				case "pipeline-install":
					return RootScreen(state.Get()).SwitchScreen(NewPipelineDownloadDefaultPage())
				case "connections-new":
					return RootScreen(state.Get()).SwitchScreen(NewConnectionsInitPage(nil))
				case "connections-manage":
					return RootScreen(state.Get()).SwitchScreen(NewConnectionsManagePage())
				case "services-init":
					return RootScreen(state.Get()).SwitchScreen(NewServiceInitPage())
				case "calibrate":
					return RootScreen(state.Get()).SwitchScreen(NewRoverCalibrationPage())
				case "services-upload":
					{
						path := "."
						_, err := utils.GetServiceInformation(path)
						if err != nil {
							return RootScreen(state.Get()).SwitchScreen(NewServiceSyncInstructionsPage())
						} else {
							return RootScreen(state.Get()).SwitchScreen(NewServicesSyncPage([]string{path}))
						}

					}
				case "services-list":
					return RootScreen(state.Get()).SwitchScreen(NewServicesListPage())
				case "info":
					return RootScreen(state.Get()).SwitchScreen(NewInfoPage())
				case "shutdown":
					return RootScreen(state.Get()).SwitchScreen(NewShutdownRoverPage())
				case "update-roverctl":
					state.Get().QuitCommand = "curl -fsSL https://raw.githubusercontent.com/VU-ASE/roverctl/main/install.sh | bash"
					return m, tea.Quit
				case "ssh":
					active := state.Get().RoverConnections.GetActive()
					if active != nil {
						state.Get().QuitCommand = "ssh " + active.Username + "@" + active.Host
					}
					return m, tea.Quit
				default:
					// state.Get().Route.Push(value)
				}
				return m, tea.Quit
			}
		case key.Matches(msg, m.keys().Down):
			if m.listIndex < len(m.listItems) {
				if len(m.listItems) > 0 && m.itemIndex < len(m.listItems[m.listIndex].items)-1 {
					m.itemIndex++
				} else if m.listIndex < len(m.listItems)-1 {
					m.listIndex++
					m.itemIndex = 0
				}
			}
		case key.Matches(msg, m.keys().Up):
			if m.itemIndex > 0 {
				m.itemIndex--
			} else if m.listIndex > 0 && len(m.listItems) > 0 {
				m.listIndex--
				m.itemIndex = len(m.listItems[m.listIndex].items) - 1
			}
		// Allow overriding rover connection state (force online)
		case msg.String() == "ctrl+o":
			m.forceOnline = !m.forceOnline
		}
	}

	var cmd tea.Cmd
	m.filterInput, cmd = m.filterInput.Update(msg)
	m.listItems = m.filterCategories()

	if m.listIndex >= len(m.listItems) {
		m.listIndex = len(m.listItems) - 1
	}
	// Adjust the list index to be in bounds
	if m.listIndex >= len(m.listItems) {
		m.listIndex = len(m.listItems) - 1
	}
	// Set the item to the last item in the list if it is out of bounds
	if len(m.listItems) > 0 && m.itemIndex >= len(m.listItems[m.listIndex].items) {
		m.itemIndex = len(m.listItems[m.listIndex].items) - 1
	}

	if m.listIndex < 0 {
		m.listIndex = 0
	}
	if m.itemIndex < 0 {
		m.itemIndex = 0
	}

	// If the listIndex is currently in a list without items, move it to the first list with items (if any)
	if len(m.listItems) > 0 && len(m.listItems[m.listIndex].items) == 0 {
		for i, category := range m.listItems {
			if len(category.items) > 0 {
				m.listIndex = i
				m.itemIndex = 0
				break
			}
		}
	}

	return m, cmd
}

func (m StartPage) View() string {
	s := m.filterInput.View()[2:]
	s += "\n\n"

	if len(m.listItems) == 0 {
		s += style.Gray.Render(" No commands found") + "\n"
	}

	for i, category := range m.listItems {
		n := lipgloss.NewStyle().Bold(false).Render(category.name)
		if i == m.listIndex {
			n = lipgloss.NewStyle().Bold(false).Underline(false).Render(category.name)
		}
		s += n + "\n"
		if category.kind == "rover" && m.roverOnline.HasResult() && !m.roverOnline.Result() && !m.forceOnline {
			s += style.Warning.Render(" ! ") + style.Gray.Render("Powered off") + "\n\n"
			continue
		}

		rightItems := []string{}
		for j, item := range category.items {
			label := item.label
			if item.key == "update-roverctl" {
				if m.updateAvailable.IsSuccess() && m.updateAvailable.Data.Available {
					label = "Update roverctl to " + style.Success.Render("v"+m.updateAvailable.Data.LatestVersion)
				} else {
					label = "Update roverctl"
				}
			}

			if i == m.listIndex && j == m.itemIndex {
				rightItems = append(rightItems, style.Primary.Bold(true).Render(" > "+label))
			} else {
				rightItems = append(rightItems, style.Gray.Render(" • "+label))
			}
		}
		rightCol := strings.Join(rightItems, "\n")

		s += rightCol + "\n\n"
	}

	s += "\n"

	return s
}

func (m StartPage) isQuitable() bool {
	return true
}

func (m StartPage) keys() utils.GeneralKeyMap {
	kb := utils.NewGeneralKeyMap()
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
		key.WithHelp("enter", "select"),
	)
	kb.Back.SetEnabled(false)
	return kb
}

func (m StartPage) previousPage() *pageModel {
	return nil
}

func (m StartPage) getSelectedItem() CategoryItem {
	if m.listIndex < len(m.listItems) && m.itemIndex < len(m.listItems[m.listIndex].items) {
		return m.listItems[m.listIndex].items[m.itemIndex]
	}
	return CategoryItem{}
}

func (m StartPage) filterCategories() []Category {
	oldCategories := initialCategories()

	newCategories := make([]Category, 0)
	for _, category := range oldCategories {
		// Filter out values if the rover is offline
		if category.kind == "rover" && m.roverOnline.HasResult() && !m.roverOnline.Result() && !m.forceOnline {
			category.items = []CategoryItem{}
		}

		// Add this category if its name (not items) matches the filter
		if strings.Contains(strings.ToLower(category.name), strings.ToLower(m.filterInput.Value())) {
			newCategories = append(newCategories, category)
			continue
		}

		// Add all items that match the filter
		newItems := make([]CategoryItem, 0)
		for _, item := range category.items {
			if m.filterInput.Value() == "" || strings.Contains(strings.ToLower(item.label), strings.ToLower(m.filterInput.Value())) {
				newItems = append(newItems, item)
			}
		}
		if len(newItems) > 0 {
			newCategories = append(newCategories, Category{
				name:  category.name,
				kind:  category.kind,
				items: newItems,
			})
		}
	}

	return newCategories
}

func initialCategories() []Category {
	// Default state (i.e. not connected to a Rover)
	roverCategory := Category{
		name: "Rover",
		items: []CategoryItem{
			{
				label: "Connect",
				key:   "connections-new",
			},
		},
	}
	localServiceCategory := Category{
		name: "Local Services",
		items: []CategoryItem{
			{
				label: "Create",
				key:   "services-init",
			},
		},
	}
	miscCategory := Category{
		name: "Misc",
		items: []CategoryItem{
			{
				label: "About",
				key:   "info",
			},
			{
				label: "Update roverctl",
				key:   "update-roverctl",
			},
		},
	}

	s := state.Get()
	if len(s.RoverConnections.Available) > 0 {
		roverCategory = Category{
			name: s.RoverConnections.Active,
			kind: "rover",
			items: []CategoryItem{
				{
					label: "Pipeline",
					key:   "pipeline-manage",
				},
				{
					label: "Status",
					key:   "info",
				},
				{
					label: "SSH",
					key:   "ssh",
				},
				{
					label: "Calibrate",
					key:   "calibrate",
				},
				{
					label: "Shutdown",
					key:   "shutdown",
				},
			},
		}
		localServiceCategory = Category{
			name: "Local Services",
			items: []CategoryItem{
				{
					label: "Upload to Rover",
					key:   "services-upload",
				},
				{
					label: "Create",
					key:   "services-init",
				},
			},
		}
		connectionsCategory := Category{
			name: "Connections",
			items: []CategoryItem{
				{
					label: "New",
					key:   "connections-new",
				},
			},
		}
		if len(s.RoverConnections.Available) > 1 {
			connectionsCategory.items = append(connectionsCategory.items, CategoryItem{
				label: "Switch active Rover",
				key:   "connections-manage",
			})
		}
		miscCategory = Category{
			name: "Misc",
			items: []CategoryItem{
				{
					label: "About",
					key:   "info",
				},
				{
					label: "Update roverctl",
					key:   "update-roverctl",
				},
			},
		}

		return []Category{
			roverCategory,
			localServiceCategory,
			connectionsCategory,
			miscCategory,
		}
	}

	return []Category{
		roverCategory,
		localServiceCategory,
		miscCategory,
	}
}

//
// Actions
//

func (m StartPage) checkUpdate() tea.Cmd {
	return tui.PerformAction(&m.updateAvailable, func() (*utils.UpdateAvailable, error) {
		return utils.CheckForGithubUpdate("roverctl", "vu-ase", version)
	})
}

func (m StartPage) checkRoverOnline(wait bool) tea.Cmd {
	return tui.PerformActionV2(&m.roverOnline, nil, func() (*bool, []error) {
		if wait {
			time.Sleep(time.Second * 1) // so that we don't poll all the time
		}

		res := false
		rover := state.Get().RoverConnections.GetActive()
		if rover == nil {
			return &res, []error{fmt.Errorf("No active rover connection")}
		}

		res = utils.IsHostOnline(rover.Host, "80", time.Millisecond*500)
		return &res, nil
	})
}
