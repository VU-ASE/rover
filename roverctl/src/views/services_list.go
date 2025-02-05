package views

import (
	"context"
	"fmt"
	"time"

	"github.com/VU-ASE/roverctl/src/openapi"
	"github.com/VU-ASE/roverctl/src/state"
	"github.com/VU-ASE/roverctl/src/style"
	"github.com/VU-ASE/roverctl/src/tui"
	"github.com/VU-ASE/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//
// The page model
//

type ServicesListPage struct {
	spinner spinner.Model
	// actions
	authors        tui.Action[[]string]
	services       tui.Action[[]string]
	versions       tui.Action[[]string]
	versionDetails tui.Action[openapi.ServicesAuthorServiceVersionGet200Response]
	delete         tui.Action[openapi.ServicesAuthorServiceVersionDelete200Response]
	// keep level state
	selectedAuthor  string
	selectedService string
	selectedVersion string
	// table to display the services
	table table.Model
}

func NewServicesListPage() ServicesListPage {
	t := table.New()

	return ServicesListPage{
		spinner: spinner.New(),
		// actions
		authors:        tui.NewAction[[]string]("fetchAuthors"),
		services:       tui.NewAction[[]string]("fetchServices"),
		versions:       tui.NewAction[[]string]("fetchVersions"),
		versionDetails: tui.NewAction[openapi.ServicesAuthorServiceVersionGet200Response]("fetchVersionDetails"),
		delete:         tui.NewAction[openapi.ServicesAuthorServiceVersionDelete200Response]("deleteVersion"),
		// keep level state
		selectedAuthor:  "",
		selectedService: "",
		selectedVersion: "",
		// table to display the services
		table: t,
	}
}

//
// Page model methods
//

func (m ServicesListPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	case tui.ActionInit[[]string]:
		m.authors.ProcessInit(msg)
		m.services.ProcessInit(msg)
		m.versions.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[[]string]:
		m.authors.ProcessResult(msg)
		m.services.ProcessResult(msg)
		m.versions.ProcessResult(msg)
		m.table = m.createTable()
		return m, nil
	case tui.ActionInit[openapi.ServicesAuthorServiceVersionGet200Response]:
		m.versionDetails.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[openapi.ServicesAuthorServiceVersionGet200Response]:
		m.versionDetails.ProcessResult(msg)
		return m, nil
	case tui.ActionInit[openapi.ServicesAuthorServiceVersionDelete200Response]:
		m.delete.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[openapi.ServicesAuthorServiceVersionDelete200Response]:
		m.delete.ProcessResult(msg)
		if m.delete.IsSuccess() {
			m.selectedVersion = ""
			cmd = m.fetchVersions()
			m.delete = tui.NewAction[openapi.ServicesAuthorServiceVersionDelete200Response]("deleteVersion")
		}
		return m, cmd
	case tea.KeyMsg:
		switch {
		// case key.Matches(msg, servicesListKeysRegular.Quit):
		// 	return m, tea.Quit
		// case key.Matches(msg, servicesListKeysRegular.Retry):
		// 	// reset all state,
		// 	m.selectedAuthor = ""
		// 	m.selectedService = ""
		// 	m.selectedVersion = ""
		// 	return m, m.fetchAuthors()
		// case key.Matches(msg, servicesListKeysRegular.Confirm):
		// 	sel := m.table.SelectedRow()
		// 	if sel == nil {
		// 		return m, nil
		// 	}

		// 	if m.selectedAuthor == "" {
		// 		m.selectedAuthor = sel[0]
		// 		cmd = m.fetchServices()
		// 	} else if m.selectedService == "" {
		// 		m.selectedService = sel[0]
		// 		cmd = m.fetchVersions()
		// 	} else {
		// 		m.selectedVersion = sel[0]
		// 		cmd = m.fetchVersionDetails()
		// 	}

		// 	m.table = m.createTable()
		// 	return m, cmd
		// case key.Matches(msg, servicesListKeysRegular.Back):
		// 	if m.selectedVersion != "" {
		// 		m.selectedVersion = ""
		// 	} else if m.selectedService != "" {
		// 		m.selectedService = ""
		// 	} else if m.selectedAuthor != "" {
		// 		m.selectedAuthor = ""
		// 	}
		// 	m.table = m.createTable()
		// 	return m, nil
		// case key.Matches(msg, servicesListKeysVersionDetails.Delete):
		// 	return m, m.deleteVersion()
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m ServicesListPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchAuthors())
}

func (m ServicesListPage) versionDetailsView() string {
	s := m.selectedAuthor + "/" + m.selectedService + "/" + m.selectedVersion + "\n"

	if m.versionDetails.IsSuccess() {
		s += "\n"
		s += style.Gray.Render("Last build: ")
		if m.versionDetails.Data.BuiltAt != nil {
			s += time.Unix(*m.versionDetails.Data.BuiltAt/1000, 0).String()
		} else {
			s += "unknown"
		}
		s += "\n" + style.Gray.Render("Inputs: ")
		for _, input := range m.versionDetails.Data.Inputs {
			for _, stream := range input.Streams {
				s += "\n- " + stream + style.Gray.Render(" from ") + input.Service
			}
		}
		if len(m.versionDetails.Data.Inputs) == 0 {
			s += "none"
		}
		s += "\n" + style.Gray.Render("Outputs: ")
		for _, output := range m.versionDetails.Data.Outputs {
			s += "\n- " + output
		}
		if len(m.versionDetails.Data.Outputs) == 0 {
			s += "none"
		}

	} else if m.versionDetails.IsError() {
		s += m.versionDetails.Error.Error()
	} else {
		s += m.spinner.View() + " Loading..."
	}

	if m.delete.IsLoading() {
		s += "\n\n" + m.spinner.View() + " Deleting version..."
	}

	return s + "\n\n"
}

func (m ServicesListPage) View() string {
	s := style.Title.Render("Installed services") + "\n\n"

	if m.selectedVersion != "" {
		return s + m.versionDetailsView()
	}

	h := m.table.HelpView()
	if m.selectedService != "" && m.selectedAuthor != "" {
		if m.versions.IsSuccess() {
			if len(*m.versions.Data) <= 0 {
				s += style.Gray.Render("No installed versions found for service ") + m.selectedAuthor + "/" + m.selectedService
			} else {
				s += m.table.View()
			}
		} else if m.versions.IsError() {
			s += "Could not fetch versions (" + m.versions.Error.Error() + ")"
		} else {
			s += m.spinner.View() + " Fetching versions for " + m.selectedAuthor + "/" + m.selectedService
		}

	} else if m.selectedAuthor != "" {
		if m.services.IsSuccess() {
			if len(*m.services.Data) <= 0 {
				s += style.Gray.Render("No services found for author ") + m.selectedAuthor
			} else {
				s += m.table.View()
			}
		} else if m.services.IsError() {
			s += "Could not fetch services (" + m.services.Error.Error() + ")"
		} else {
			s += m.spinner.View() + " Fetching services for " + m.selectedAuthor
		}

	} else {
		if m.authors.IsSuccess() {
			if len(*m.authors.Data) <= 0 {
				s += style.Gray.Render("No authors found. Go ahead and create a service!")
			} else {
				s += m.table.View()
			}
		} else if m.authors.IsError() {
			s += style.Error.Render("Could not fetch authors") + style.Gray.Render(" ("+m.authors.Error.Error()+")")
		} else {
			s += m.spinner.View() + " Fetching authors"
		}

	}

	if m.delete.IsLoading() {
		s += "\n\n" + m.spinner.View() + " Deleting..."
	}

	return s + "\n\n" + h
}

//
// Actions
//

func (m ServicesListPage) fetchAuthors() tea.Cmd {
	return tui.PerformAction(&m.authors, func() (*[]string, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesGet(
			context.Background(),
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return &res, err
	})
}

func (m ServicesListPage) fetchServices() tea.Cmd {
	return tui.PerformAction(&m.services, func() (*[]string, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorGet(
			context.Background(),
			m.selectedAuthor,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return &res, err
	})
}

func (m ServicesListPage) fetchVersions() tea.Cmd {
	return tui.PerformAction(&m.versions, func() (*[]string, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorServiceGet(
			context.Background(),
			m.selectedAuthor,
			m.selectedService,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		sorted := utils.SortByVersion(res)

		return &sorted, err
	})
}

func (m ServicesListPage) fetchVersionDetails() tea.Cmd {
	return tui.PerformAction(&m.versionDetails, func() (*openapi.ServicesAuthorServiceVersionGet200Response, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorServiceVersionGet(
			context.Background(),
			m.selectedAuthor,
			m.selectedService,
			m.selectedVersion,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return res, err
	})
}

func (m ServicesListPage) deleteVersion() tea.Cmd {
	return tui.PerformAction(&m.delete, func() (*openapi.ServicesAuthorServiceVersionDelete200Response, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorServiceVersionDelete(
			context.Background(),
			m.selectedAuthor,
			m.selectedService,
			m.selectedVersion,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return res, err
	})
}

//
// Table rendering
//

func (m ServicesListPage) getColWidth() int {
	return (state.Get().WindowWidth - 4 - 6) / 1 // Adjust for padding and borders
}

func (m ServicesListPage) colPct(pct int) int {
	total := m.getColWidth() - 2
	return (total*pct)/100 - 1
}

func (m ServicesListPage) createTable() table.Model {
	var columns []table.Column
	rows := make([]table.Row, 0)

	if m.selectedService != "" && m.selectedAuthor != "" {
		columns = []table.Column{
			{Title: "Versions (" + m.selectedAuthor + "/" + m.selectedService + ")", Width: m.colPct(100)},
		}

		if m.versions.HasData() {
			for _, version := range *m.versions.Data {
				rows = append(rows, table.Row{version})
			}
		}
	} else if m.selectedAuthor != "" {
		columns = []table.Column{
			{Title: "Services (" + m.selectedAuthor + ")", Width: m.colPct(100)},
		}

		if m.services.HasData() {
			for _, service := range *m.services.Data {
				rows = append(rows, table.Row{service})
			}
		}
	} else {
		columns = []table.Column{
			{Title: "Authors", Width: m.colPct(100)},
		}

		if m.authors.HasData() {
			for _, author := range *m.authors.Data {
				rows = append(rows, table.Row{author})
			}
		}
	}

	cursor := 0

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
	)
	if len(rows) < 7 {
		t.SetHeight(len(rows) + 1)
	}
	t.SetCursor(cursor)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("FFFFFF")).
		Background(style.AsePrimary).
		Bold(false)

	t.SetStyles(s)

	return t
}

func (m ServicesListPage) isQuitable() bool {
	return true
}

func (m ServicesListPage) keys() utils.GeneralKeyMap {
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
	return kb
}

func (m ServicesListPage) previousPage() *pageModel {
	return nil
}
