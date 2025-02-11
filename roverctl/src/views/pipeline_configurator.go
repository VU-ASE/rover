package views

import (
	"context"
	"fmt"
	"regexp"

	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lempiy/dgraph"
	"github.com/lempiy/dgraph/core"
)

//
// All keys
//

// Keys to navigate
type PipelineConfiguratorKeyMap struct {
	Retry   key.Binding
	Confirm key.Binding
	Switch  key.Binding // switch table focus
	Remove  key.Binding // remove service from pipeline
	Back    key.Binding // go back one level
	Save    key.Binding // save the pipeline
	Quit    key.Binding
	Refetch key.Binding
}

// Shown when the services are being updated
var pipelineConfiguratorKeysRegular = PipelineConfiguratorKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "refetch"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),
	Back: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "go back one level"),
	),
	Switch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch table"),
	),
	Remove: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "remove"),
	),
	Save: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "save"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

// When the "active" table is focussed
var pipelineConfiguratorKeysActiveTable = PipelineConfiguratorKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "refetch"),
	),
	Switch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch table"),
	),
	Remove: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "remove from pipeline"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

// When the "remote" table is focussed on the "select an author" level
var pipelineConfiguratorKeysRemoteTableAuthor = PipelineConfiguratorKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "refetch"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view services"),
	),
	Switch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch table"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

// When the "remote" table is focussed on the "select a service" level
var pipelineConfiguratorKeysRemoteTableService = PipelineConfiguratorKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "refetch"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view versions"),
	),
	Back: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "view all authors"),
	),
	Switch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch table"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

// When the "remote" table is focussed on the "select a version" level
var pipelineConfiguratorKeysRemoteTableVersion = PipelineConfiguratorKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "refetch"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "add to pipeline"),
	),
	Back: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "view all services"),
	),
	Switch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch table"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

func (k PipelineConfiguratorKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Retry, k.Confirm, k.Remove, k.Back, k.Switch, k.Quit}
}

func (k PipelineConfiguratorKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

//
// The page model
//

type PipelineConfiguratorPage struct {
	help    help.Model
	spinner spinner.Model
	// tables that are shown next to each other
	tableActive table.Model // the pipeline as it is configured now, with the enabled services
	tableRemote table.Model // all remote services that still can be enabled
	// actions to fetch data to populate the tables
	pipeline         tui.Action[PipelineOverviewSummary]
	pipelineGraph    string               // preserved in the model to avoid re-rendering in .View()
	dependencyErrors []error              // errors in the pipeline configuration
	authors          tui.Action[[]string] // first part of FQN
	services         tui.Action[[]string] // second part of FQN
	versions         tui.Action[[]string] // third part of FQN
	savePipeline     tui.Action[bool]     // save the pipeline
	// Keep track of the focussed table (left/right)
	focussed int // 0 = active, 1 = remote
	// For querying remote services
	remoteAuthor  string
	remoteService string
}

func NewPipelineConfiguratorPage() PipelineConfiguratorPage {
	// todo

	return PipelineConfiguratorPage{
		spinner:       spinner.New(),
		help:          help.New(),
		tableActive:   table.New(),
		tableRemote:   table.New(),
		pipeline:      tui.NewAction[PipelineOverviewSummary]("fetchActive"),
		authors:       tui.NewAction[[]string]("fetchAuthors"),
		services:      tui.NewAction[[]string]("fetchServices"),
		versions:      tui.NewAction[[]string]("fetchVersions"),
		savePipeline:  tui.NewAction[bool]("savePipeline"),
		focussed:      1,
		remoteAuthor:  "",
		remoteService: "",
	}
}

//
// Page model methods
//

func (m PipelineConfiguratorPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.pipeline.HasData() {
			m.tableActive = m.createActiveTable(*m.pipeline.Data)
		} else {
			m.tableActive = m.createActiveTable(PipelineOverviewSummary{})
		}
		m.tableRemote = m.createRemoteTable()
		return m, nil
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tui.ActionInit[bool]:
		m.savePipeline.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[bool]:
		m.savePipeline.ProcessResult(msg)
		return m, nil
	case tui.ActionInit[PipelineOverviewSummary]:
		m.pipeline.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[PipelineOverviewSummary]:
		m.pipeline.ProcessResult(msg)
		if m.pipeline.IsSuccess() {
			if len(m.pipeline.Data.Pipeline.Enabled) <= 0 {
				m.focussed = 1 // nothing to navigate in the active table
				m.tableRemote = m.createRemoteTable()
			}

			// Create the pipeline graph based on enabled services
			nodes := make([]core.NodeInput, 0)
			for _, service := range m.pipeline.Data.Pipeline.Enabled {
				// Check if the service is selected, in this case unselect it
				if m.remoteService == service.Service.Fq.Name {
					m.remoteService = ""
				}

				nodes = append(nodes, core.NodeInput{
					Id: service.Service.Fq.Name,
					Next: func() []string {
						// Find services that depend on an output of this service
						found := make([]string, 0)
						for _, s := range m.pipeline.Data.Services {
							if s.Name != service.Service.Fq.Name {
								for _, input := range s.Configuration.Inputs {
									if input.Service == service.Service.Fq.Name {
										found = append(found, s.Name)
									}
								}
							}
						}

						return found
					}(),
				})
			}
			canvas, err := dgraph.DrawGraph(nodes)
			if len(nodes) <= 0 {
				m.pipelineGraph = style.Gray.Render("")
			} else if err != nil {
				m.pipelineGraph = "Failed to draw pipeline\n"
			} else {
				m.pipelineGraph = fmt.Sprintf("%s\n", canvas)
			}
			m.dependencyErrors = m.findDependencyErrors()
			m.tableActive = m.createActiveTable(*m.pipeline.Data)
			m.tableRemote = m.createRemoteTable()
		}
		m.savePipeline = tui.NewAction[bool]("savePipeline")
		return m, nil
	case tui.ActionInit[[]string]:
		m.authors.ProcessInit(msg)
		m.services.ProcessInit(msg)
		m.versions.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[[]string]:
		m.authors.ProcessResult(msg)
		m.services.ProcessResult(msg)
		m.versions.ProcessResult(msg)
		m.tableRemote = m.createRemoteTable()
		return m, nil
	case tea.KeyMsg:
		switch {
		// case key.Matches(msg, pipelineConfiguratorKeysRegular.Quit):
		// return m, tea.Quit
		case key.Matches(msg, pipelineConfiguratorKeysRegular.Remove):
			if m.focussed == 0 {
				return m.onActiveTableNavigation(msg)
			}

			// todo:
			return m, nil
		case key.Matches(msg, pipelineConfiguratorKeysRegular.Confirm):
			if m.focussed == 1 {
				return m.onRemoteTableNavigation(msg)
			}

			// todo:
			return m, nil
		case key.Matches(msg, pipelineConfiguratorKeysRegular.Back):
			if m.focussed == 1 {
				return m.onRemoteTableNavigation(msg)
			}

			// todo:
			return m, nil
		case key.Matches(msg, pipelineConfiguratorKeysRegular.Retry):
			// Redo all actions, reset to the initial state
			cmds := tea.Batch(
				m.fetchAllAuthors(),
				m.fetchPipeline(),
			)
			return m, cmds
		case key.Matches(msg, pipelineConfiguratorKeysRegular.Save):
			if len(m.dependencyErrors) <= 0 {
				return m, m.savePipelineRemote()
			}
		case key.Matches(msg, pipelineConfiguratorKeysRegular.Switch):
			if m.focussed == 1 && len(m.tableActive.Rows()) <= 0 {
				return m, nil
			}

			m.focussed = (m.focussed + 1) % 2
			m.tableActive = m.createActiveTable(*m.pipeline.Data)
			m.tableRemote = m.createRemoteTable()
			return m, nil
		}
	}

	if m.focussed == 0 {
		m.tableActive, cmd = m.tableActive.Update(msg)
	} else if m.focussed == 1 {
		m.tableRemote, cmd = m.tableRemote.Update(msg)
	}

	return m, cmd
}

func (m PipelineConfiguratorPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchPipeline(), m.fetchAllAuthors())
}

func (m PipelineConfiguratorPage) remoteTableView() string {
	note := ""
	if m.remoteAuthor == "" {
		if m.authors.IsSuccess() {
			if len(m.tableRemote.Rows()) <= 0 {
				note = style.Gray.Render(" No remote authors available. Upload a service first.")
			}

			return m.tableRemote.View() + note
		} else if m.authors.IsError() {
			return style.Error.Render("Error loading available authors") + style.Gray.Render(" "+m.authors.Error.Error())
		} else {
			return m.spinner.View() + style.Gray.Render(" Loading authors")
		}
	} else if m.remoteService == "" {
		if m.services.IsSuccess() {
			if len(m.tableRemote.Rows()) <= 0 {
				note = style.Gray.Render(" No unused services available for " + m.remoteAuthor)
			}

			return m.tableRemote.View() + note
		} else if m.services.IsError() {
			return style.Error.Render("Error loading services for "+m.remoteAuthor) + style.Gray.Render(" "+m.services.Error.Error())
		} else {
			return m.spinner.View() + style.Gray.Render(" Loading services for "+m.remoteAuthor)
		}
	} else if m.versions.IsSuccess() {
		if len(m.tableRemote.Rows()) <= 0 {
			note = style.Gray.Render(" No unused versions available for " + m.remoteAuthor + "/" + m.remoteService)
		}

		return m.tableRemote.View() + note
	} else if m.versions.IsError() {
		return style.Error.Render("Error loading versions for "+m.remoteAuthor+"/"+m.remoteService) + style.Gray.Render(" "+m.versions.Error.Error())
	} else {
		return m.spinner.View() + style.Gray.Render(" Loading versions for "+m.remoteAuthor+"/"+m.remoteService)
	}
}

func (m PipelineConfiguratorPage) View() string {
	s := style.Title.Render("Pipeline configurator") + "\n\n"

	s += style.Primary.Render("A pipeline is composed of ") + "services" + style.Primary.Render(" which can be enabled to run when the pipeline is started.") + "\n" + style.Primary.Render("By moving services between ") + "available" + style.Primary.Render(" and ") + "enabled" + style.Primary.Render(", you can modify the pipeline.") + "\n\n"

	// Calculate column width (subtract padding and borders)
	columnWidth := m.getColWidth()

	// Define styles for each column
	columnStyle := lipgloss.NewStyle().Width(columnWidth)

	// Define the columns
	columnActive := m.spinner.View() + style.Gray.Render(" Loading active services...")
	if m.pipeline.IsSuccess() {
		note := ""
		if len(m.tableActive.Rows()) <= 0 {
			note = style.Gray.Render("This pipeline is empty. Start by enabling a service.")
		}

		columnActive = m.tableActive.View() + note
	} else if m.pipeline.IsError() {
		columnActive = style.Error.Render("Error loading active services") + style.Gray.Render(" "+m.pipeline.Error.Error())
	}

	columnRemote := m.remoteTableView()

	row := lipgloss.JoinHorizontal(lipgloss.Top,
		columnStyle.Render(columnRemote),
		" ",
		columnStyle.Render(columnActive),
	)

	h := ""
	if m.focussed == 0 {
		h = m.tableActive.HelpView() + style.Gray.Render(" • ") + m.help.View(pipelineConfiguratorKeysActiveTable)
	} else {
		h = m.tableRemote.HelpView() + style.Gray.Render(" • ")
		if m.remoteService != "" {
			h += m.help.View(pipelineConfiguratorKeysRemoteTableVersion)
		} else if m.remoteAuthor != "" {
			h += m.help.View(pipelineConfiguratorKeysRemoteTableService)
		} else {
			h += m.help.View(pipelineConfiguratorKeysRemoteTableAuthor)
		}
	}

	// You might be tempted to extract the "\n\n" but we don't want to render an empty space if there is no graph
	graph := ""
	if m.pipeline.IsSuccess() {
		if m.pipelineGraph != "" {
			graph = "\n\n" + m.postProcessGraph(m.pipelineGraph)
		}
	} else if m.pipeline.IsError() {
		graph = "\n\n" + style.Error.Render("Could not fetch pipeline: ") + style.Gray.Render(m.pipeline.Error.Error())
	} else {
		graph = "\n\n" + m.spinner.View() + " Loading pipeline"
	}

	status := ""
	if m.pipeline.HasData() {
		if len(m.pipeline.Data.Services) > 0 && len(m.dependencyErrors) <= 0 {
			status += "\n" + style.Success.Render("✓ This pipeline is valid") + " " + style.Gray.Render("- save it to your Rover with ") + "s"
		} else if len(m.pipeline.Data.Services) > 0 {
			for _, err := range m.dependencyErrors {
				status += "\n" + style.Error.Render("✗ "+err.Error())
			}
		}
	}
	if m.savePipeline.IsLoading() {
		status += "\n" + m.spinner.View() + " Saving pipeline\n\n"
	} else if m.savePipeline.IsSuccess() {
		status += "\n" + style.Success.Render("✓ Pipeline saved to Rover successfully") + "\n\n"
	} else if m.savePipeline.IsError() {
		status += "\n" + style.Error.Render("✗ Error saving pipeline") + style.Gray.Render(" "+m.savePipeline.Error.Error()) + "\n\n"
	} else {
		status += "\n\n"
	}

	return s + row + graph + status + h
}

//
// Aux for table and col rendering
//

func (m PipelineConfiguratorPage) getColWidth() int {
	return (state.Get().WindowWidth - 4 - 6) / 2 // Adjust for padding and borders
}

func (m PipelineConfiguratorPage) colPct(pct int) int {
	total := m.getColWidth() - 2
	return (total*pct)/100 - 1
}

//
// Actions
//

func (m PipelineConfiguratorPage) fetchPipeline() tea.Cmd {
	return tui.PerformAction(&m.pipeline, func() (*PipelineOverviewSummary, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()

		// First, fetch all services and the status of the current pipeline
		pipeline, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		// Then, for each service, we need to query the service for its actual configuration (inputs, outputs)
		services := make([]PipelineOverviewServiceInfo, 0)
		for _, enabled := range pipeline.Enabled {
			configuration, htt, err := api.ServicesAPI.ServicesAuthorServiceVersionGet(
				context.Background(),
				enabled.Service.Fq.Author,
				enabled.Service.Fq.Name,
				enabled.Service.Fq.Version,
			).Execute()

			if err != nil && htt != nil {
				return nil, utils.ParseHTTPError(err, htt)
			}

			services = append(services, PipelineOverviewServiceInfo{
				Name:          enabled.Service.Fq.Name,
				Version:       enabled.Service.Fq.Version,
				Author:        enabled.Service.Fq.Author,
				Configuration: *configuration,
			})
		}

		// Then the Rover status
		status, htt, err := api.HealthAPI.StatusGet(
			context.Background(),
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		// Combined response
		res := PipelineOverviewSummary{
			Pipeline: *pipeline,
			Services: services,
			Status:   *status,
		}

		return &res, nil
	})
}

//
// Tables
//

func (m PipelineConfiguratorPage) createActiveTable(res PipelineOverviewSummary) table.Model {
	// Retrieve the previously selected entry
	prev := m.tableActive.SelectedRow()

	columns := []table.Column{
		{Title: "Enabled services in this pipeline", Width: m.colPct(99)},
	}

	rows := make([]table.Row, 0)
	for _, _ = range res.Services {
		rows = append(rows, table.Row{
			// utils.ServiceFqn(service.Author, service.Name, service.Version),
			"",
		})
	}

	// Find the previously selected row, if it exists
	cursor := 0
	if prev != nil {
		for i, row := range rows {
			if len(row) >= 3 && len(prev) >= 3 && row[0] == prev[0] && row[1] == prev[1] && row[2] == prev[2] {
				cursor = i
				break
			}
		}
	}

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

	if m.focussed == 0 {
		t.Focus()
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("FFF")).
			Background(style.AsePrimary).
			Bold(false)
	} else {
		t.Blur()
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("FFF")).
			Background(style.GrayPrimary).
			Bold(false)
	}

	t.SetStyles(s)

	return t
}

func (m PipelineConfiguratorPage) createRemoteTable() table.Model {
	// Retrieve the previously selected entry
	prev := m.tableRemote.SelectedRow()

	columns := []table.Column{
		{Title: "Available authors", Width: m.colPct(100)},
	}
	rows := make([]table.Row, 0)
	// Go from most fine-grained to least fine-grained
	if m.remoteService != "" {
		columns = []table.Column{
			{Title: fmt.Sprintf("Available versions for %s/%s", m.remoteAuthor, m.remoteService), Width: m.colPct(100)},
		}
		if m.versions.HasData() {
			for _, version := range *m.versions.Data {
				// Does this version already exist in the pipeline?
				exists := false
				if m.pipeline.HasData() {
					for _, enabled := range m.pipeline.Data.Pipeline.Enabled {
						if enabled.Service.Fq.Name == m.remoteService && enabled.Service.Fq.Version == version {
							exists = true
							break
						}
					}
				}

				if !exists {
					rows = append(rows, table.Row{
						version,
					})
				}
			}
		}
	} else if m.remoteAuthor != "" {
		columns = []table.Column{
			{Title: fmt.Sprintf("Available services from author %s", m.remoteAuthor), Width: m.colPct(100)},
		}
		if m.services.HasData() {
			for _, service := range *m.services.Data {
				// Does this service already exist in the pipeline?
				exists := false
				if m.pipeline.HasData() {
					for _, enabled := range m.pipeline.Data.Pipeline.Enabled {
						if enabled.Service.Fq.Name == service {
							exists = true
							break
						}
					}
				}

				if !exists {
					rows = append(rows, table.Row{
						service,
					})
				}
			}
		}

	} else {
		if m.authors.HasData() {
			for _, author := range *m.authors.Data {
				rows = append(rows, table.Row{
					author,
				})
			}
		}
	}

	// Find the previously selected row, if it exists
	cursor := 0
	if prev != nil {
		for i, row := range rows {
			if len(row) >= 1 && len(prev) >= 1 && row[0] == prev[0] {
				cursor = i
				break
			}
		}
	}

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

	if m.focussed == 1 {
		t.Focus()
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("FFF")).
			Background(style.AsePrimary).
			Bold(false)
	} else {
		t.Blur()
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("FFF")).
			Background(style.GrayPrimary).
			Bold(false)
	}

	t.SetStyles(s)

	return t
}

//
// Actions
//

func (m PipelineConfiguratorPage) fetchAllAuthors() tea.Cmd {
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

func (m PipelineConfiguratorPage) fetchServicesForAuthor(author string) tea.Cmd {
	return tui.PerformAction(&m.services, func() (*[]string, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorGet(
			context.Background(),
			author,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return &res, err
	})
}

func (m PipelineConfiguratorPage) fetchVersionsForService(author string, service string) tea.Cmd {
	return tui.PerformAction(&m.versions, func() (*[]string, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorServiceGet(
			context.Background(),
			author,
			service,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		sorted := utils.SortByVersion(res)

		return &sorted, err
	})
}

func (m PipelineConfiguratorPage) onActiveTableNavigation(pressedKey tea.KeyMsg) (pageModel, tea.Cmd) {
	if m.focussed != 0 {
		return m, nil
	}

	if key.Matches(pressedKey, pipelineConfiguratorKeysRegular.Remove) {
		// If there is no value, nothing we can do
		sel := m.tableActive.SelectedRow()
		if len(sel) <= 0 {
			return m, nil
		}

		// Remove service from pipeline
		return m, m.removeServiceFromPipeline(sel[0])
	}

	return m, nil
}

func (m PipelineConfiguratorPage) onRemoteTableNavigation(pressedKey tea.KeyMsg) (pageModel, tea.Cmd) {
	if m.focussed != 1 {
		return m, nil
	}

	// Hitting "enter" makes the lookup go one level deeper, unless we are at the deepest level (a specific version), which will then insert the service into the pipeline
	if key.Matches(pressedKey, pipelineConfiguratorKeysRegular.Confirm) {
		// If there is no value, nothing we can do
		sel := m.tableRemote.SelectedRow()
		if len(sel) <= 0 {
			return m, nil
		}

		if m.remoteAuthor == "" {
			m.remoteAuthor = sel[0]
			return m, m.fetchServicesForAuthor(m.remoteAuthor)
		} else if m.remoteService == "" {
			m.remoteService = sel[0]
			return m, m.fetchVersionsForService(m.remoteAuthor, m.remoteService)
		} else {
			// Insert service into pipeline
			return m, m.addServiceToPipeline(m.remoteAuthor, m.remoteService, sel[0])
		}
	}

	// Hitting "backspace" goes one level up
	if key.Matches(pressedKey, pipelineConfiguratorKeysRegular.Back) {
		if m.remoteService != "" {
			m.remoteService = ""
		} else if m.remoteAuthor != "" {
			m.remoteAuthor = ""
		}
	}

	m.tableRemote = m.createRemoteTable()
	return m, nil
}

// This adds a service to a pipeline *locally*. It will only be checked by the server when the pipeline is saved.
func (m PipelineConfiguratorPage) addServiceToPipeline(author string, service string, version string) tea.Cmd {
	return tui.PerformAction(&m.pipeline, func() (*PipelineOverviewSummary, error) {
		// There should already be a pipeline in the model
		if !m.pipeline.IsSuccess() {
			return nil, fmt.Errorf("Cannot add a service to a non-fetched pipeline")
		}

		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorServiceVersionGet(
			context.Background(),
			author,
			service,
			version,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		// Add this service to the pipeline
		pipeline := *m.pipeline.Data

		// Check if the service is already in the pipeline
		for _, enabled := range pipeline.Pipeline.Enabled {
			if enabled.Service.Fq.Name == service {
				return &pipeline, nil
			}
		}

		pipeline.Services = append(pipeline.Services, PipelineOverviewServiceInfo{
			Name:          service,
			Author:        author,
			Version:       version,
			Configuration: *res,
		})
		pipeline.Pipeline.Enabled = append(pipeline.Pipeline.Enabled, openapi.PipelineGet200ResponseEnabledInner{
			Service: openapi.PipelineGet200ResponseEnabledInnerService{
				Fq: openapi.FullyQualifiedService{
					Name:    service,
					Version: version,
					Author:  author,
				},
			},
		})

		return &pipeline, nil
	})
}

// This removes a service from a pipeline *locally*. It will only be checked by the server when the pipeline is saved.
func (m PipelineConfiguratorPage) removeServiceFromPipeline(fqn string) tea.Cmd {
	return tui.PerformAction(&m.pipeline, func() (*PipelineOverviewSummary, error) {
		// There should already be a pipeline in the model
		if !m.pipeline.IsSuccess() {
			return nil, fmt.Errorf("Cannot remove a service from a non-fetched pipeline")
		}

		// Remove this service from the pipeline
		pipeline := *m.pipeline.Data
		newServices := make([]PipelineOverviewServiceInfo, 0)
		for _, s := range pipeline.Services {
			if "" != fqn {
				newServices = append(newServices, s)
			}
		}
		pipeline.Services = newServices
		newEnabled := make([]openapi.PipelineGet200ResponseEnabledInner, 0)
		for _, enabled := range pipeline.Pipeline.Enabled {
			if "" != fqn {
				newEnabled = append(newEnabled, enabled)
			}
		}
		pipeline.Pipeline.Enabled = newEnabled
		return &pipeline, nil
	})
}

func (m PipelineConfiguratorPage) findDependencyErrors() []error {
	errors := make([]error, 0)
	if !m.pipeline.IsSuccess() {
		return errors
	}

	// For each service, check if it has unmet dependencies with other services
	for _, service := range m.pipeline.Data.Services {
		for _, input := range service.Configuration.Inputs {
			for _, stream := range input.Streams {
				found := false
				for _, other := range m.pipeline.Data.Services {
					if other.Name == input.Service {
						for _, output := range other.Configuration.Outputs {
							if output == stream {
								found = true
								break
							}
						}
					}
				}

				if !found {
					errors = append(errors, fmt.Errorf("Service '%s' depends on unresolved stream '%s' from service '%s'", service.Name, stream, input.Service))
				}

			}
		}
	}

	return errors
}

func (m PipelineConfiguratorPage) savePipelineRemote() tea.Cmd {
	return tui.PerformAction(&m.savePipeline, func() (*bool, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		req := api.PipelineAPI.PipelinePost(
			context.Background(),
		)

		pipelineReq := []openapi.PipelinePostRequestInner{}
		if m.pipeline.HasData() {
			for _, service := range m.pipeline.Data.Services {
				pipelineReq = append(pipelineReq, openapi.PipelinePostRequestInner{
					Fq: openapi.FullyQualifiedService{
						Name:    service.Name,
						Version: service.Version,
						Author:  service.Author,
					},
				})
			}
		}
		req = req.PipelinePostRequestInner(pipelineReq)
		htt, err := req.Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return openapi.PtrBool(true), err
	})
}

//
// Aux methods for views
//

// Clean up the graph to make it a bit more readable and compressed
func (m PipelineConfiguratorPage) postProcessGraph(s string) string {
	n := s

	// Remove empty lines
	n = regexp.MustCompile(`\n\s*\n`).ReplaceAllString(n, "\n")

	// Highlight the currently selected service
	sel := m.tableActive.SelectedRow()
	if sel != nil {
		// The first item is always the service name
		name := sel[0]

		// Find the service in the graph
		if m.focussed == 0 {
			n = regexp.MustCompile(fmt.Sprintf(`\b%s\b`, name)).ReplaceAllString(n, style.Primary.Bold(true).Render(name))
		} else {
			n = regexp.MustCompile(fmt.Sprintf(`\b%s\b`, name)).ReplaceAllString(n, lipgloss.NewStyle().Bold(true).Render(name))
		}
	}

	return n
}

func (m PipelineConfiguratorPage) isQuitable() bool {
	return true
}

func (m PipelineConfiguratorPage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m PipelineConfiguratorPage) previousPage() *pageModel {
	return nil
}
