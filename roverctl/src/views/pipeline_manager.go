package views

import (
	"context"
	"fmt"
	"strings"

	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var defaultPipeline = []string{
	"imaging", "controller", "actuator",
}

//
// The page model
//

type PipelineManagerPage struct {
	spinner spinner.Model

	// All fetched data
	pipeline               tui.ActionV2[any, openapi.PipelineGet200Response]  // the current pipeline and its execution, periodically refetched
	services               tui.ActionV2[any, []openapi.FullyQualifiedService] // all services, of which some are part of the pipeline (= enabled)
	locallyEnabledServices []openapi.FullyQualifiedService                    // services that are enabled locally, but not per se in the remote pipeline

	// Filtering through the list of available services
	filterValue   textinput.Model
	filterEnabled bool
	selectedIndex int
	ignoreUpdates bool

	// Actions to install the basic pipeline
	defaultPipelineServices     tui.ActionV2[any, []utils.UpdateAvailable]                           // all latest services in the default autonomous driving pipeline
	missingServicesInstallation []*tui.ActionV2[utils.UpdateAvailable, openapi.FetchPost200Response] // installation if missing services (if started)

	// Actions to get more details on the selected service
	selectedServiceDetails tui.ActionV2[openapi.FullyQualifiedService, openapi.ServicesAuthorServiceVersionGet200Response]
	selectedServiceLogs    tui.ActionV2[openapi.FullyQualifiedService, []string]
}

func NewPipelineManagerPage() PipelineManagerPage {
	// Create actions
	pipeline := tui.NewActionV2[any, openapi.PipelineGet200Response]()
	services := tui.NewActionV2[any, []openapi.FullyQualifiedService]()
	locallyEnabledServices := make([]openapi.FullyQualifiedService, 0)
	selectedServiceDetails := tui.NewActionV2[openapi.FullyQualifiedService, openapi.ServicesAuthorServiceVersionGet200Response]()
	selectedServiceLogs := tui.NewActionV2[openapi.FullyQualifiedService, []string]()

	ti := textinput.New()
	ti.Placeholder = "Type to filter services..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 100

	return PipelineManagerPage{
		spinner:                spinner.New(),
		pipeline:               pipeline,
		services:               services,
		filterValue:            ti,
		locallyEnabledServices: locallyEnabledServices,
		selectedIndex:          0,
		ignoreUpdates:          false,
		selectedServiceDetails: selectedServiceDetails,
		selectedServiceLogs:    selectedServiceLogs,
	}
}

//
// Page model methods
//

func (m PipelineManagerPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch {
		// Install/update the basic pipeline
		// case msg.String() == "y" && m.defaultPipelineServicesMissing() && len(m.missingServicesInstallation) == 0:
		// for _, service := range m.defaultPipelineServices.Result() {
		// 	// Find the installed counterpart
		// 	var installed *utils.ServiceFqn
		// 	for _, i := range *m.availableServices.Data {
		// 		if i.Name == service.Name && i.Author == service.Author {
		// 			installed = &i
		// 		}
		// 	}

		// 	if installed == nil || installed.Version != service.LatestVersion {
		// 		// Create a new action and start it
		// 		action := tui.NewActionV2[utils.UpdateAvailable, openapi.FetchPost200Response]()
		// 		m.missingServicesInstallation = append(m.missingServicesInstallation, &action)
		// 		cmd = tea.Batch(cmd, m.installService(&action, service))
		// 	}
		// }
		// return m, cmd
		// case msg.String() == "r" && m.defaultPipelineServicesMissing() && len(m.missingServicesInstallation) > 0:
		// All actions need to be done
		// for _, action := range m.missingServicesInstallation {
		// 	if !action.IsDone() {
		// 		return m, nil
		// 	}
		// }

		// // Reset the actions
		// m.missingServicesInstallation = make([]*tui.ActionV2[utils.UpdateAvailable, openapi.FetchPost200Response], 0)
		// for _, service := range m.defaultPipelineServices.Result() {
		// 	// Find the installed counterpart
		// 	var installed *utils.ServiceFqn
		// 	for _, i := range *m.availableServices.Data {
		// 		if i.Name == service.Name && i.Author == service.Author {
		// 			installed = &i
		// 		}
		// 	}

		// 	if installed == nil || installed.Version != service.LatestVersion {
		// 		// Create a new action and start it
		// 		action := tui.NewActionV2[utils.UpdateAvailable, openapi.FetchPost200Response]()
		// 		m.missingServicesInstallation = append(m.missingServicesInstallation, &action)
		// 		cmd = tea.Batch(cmd, m.installService(&action, service))
		// 	}
		// }
		// return m, cmd
		// case msg.String() == "n", msg.String() == "i":
		// All actions need to be done
		// for _, action := range m.missingServicesInstallation {
		// 	if !action.IsDone() {
		// 		return m, nil
		// 	}
		// }

		// m.ignoreUpdates = true
		// return m, nil
		case key.Matches(msg, m.keys().Back):
			return m, tea.Quit
		case key.Matches(msg, m.keys().Up):
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
			selected := m.selectedService()
			if selected != nil {
				return m, tea.Batch(m.fetchServiceDetails(*selected), m.fetchServiceLogs(*selected))
			}
		case key.Matches(msg, m.keys().Down):
			if m.selectedIndex < len(m.services.Result())-1 {
				m.selectedIndex++
			}
			selected := m.selectedService()
			if selected != nil {
				return m, tea.Batch(m.fetchServiceDetails(*selected), m.fetchServiceLogs(*selected))
			}
		case key.Matches(msg, m.keys().Toggle):
			// Both pipeline and services need to be available
			if !m.services.HasResult() || !m.pipeline.HasResult() {
				return m, nil
			}

			// Must be in range
			if m.selectedIndex < 0 || m.selectedIndex >= len(m.services.Result()) {
				return m, nil
			}

			// Is this the first custom service that is enabled? Then copy the enabled services from the pipeline
			newEnabled := m.locallyEnabledServices
			if len(newEnabled) == 0 {
				for _, e := range m.pipeline.Result().Enabled {
					newEnabled = append(newEnabled, e.Service.Fq)
				}
			}

			// Find the selected service
			selected := m.services.Result()[m.selectedIndex]
			enabled := false
			for i, s := range newEnabled {
				if s.Name == selected.Name && s.Author == selected.Author && s.Version == selected.Version {
					enabled = true
					newEnabled = append(newEnabled[:i], newEnabled[i+1:]...)
					break
				}
			}
			// If not enabled, add it
			if !enabled {
				newEnabled = append(newEnabled, selected)
			}

			// The resulting list should contain at least one service
			if len(newEnabled) > 0 {
				m.locallyEnabledServices = newEnabled
			}
			return m, nil

		case key.Matches(msg, m.keys().Save):
			// if !m.pipeline.IsLoading() {
			// 	return m, m.togglePipelineExecution()
			// } else {
			// 	return m, nil
			// }
		case key.Matches(msg, m.keys().Details):
			// if m.selectedIndex >= 0 && m.selectedIndex < len(m.processedAvailableServices) {
			// 	s := m.processedAvailableServices[m.selectedIndex]
			// 	return RootScreen(state.Get()).SwitchScreen(NewPipelineLogsPage(s.service.Name, s.service.Author, s.service.Version))
			// }
		case key.Matches(msg, m.keys().Configure):
			// m.filterEnabled = !m.filterEnabled
			// m.processAvailableServices()
			// return m, nil
		}

	// Action catchers
	case tui.ActionUpdate[any, any]:
		m.pipeline.ProcessUpdate(msg)
		m.services.ProcessUpdate(msg)
		m.defaultPipelineServices.ProcessUpdate(msg)
		for _, action := range m.missingServicesInstallation {
			action.ProcessUpdate(msg)
		}
		m.selectedServiceDetails.ProcessUpdate(msg)
		m.selectedServiceLogs.ProcessUpdate(msg)

		// If this was a response to the service request, find out which services are enabled remotely
		// if msg.IsForAction() == m.services.Id() && m.services.IsSuccess() {
		// 	// Find the services that are enabled in the pipeline
		// 	enabledServices := make([]openapi.FullyQualifiedService, 0)
		// 	for _, service := range m.services.Result() {
		// 		enabledServices = append(enabledServices, service)
		// 	}

		// 	m.locallyEnabledServices = enabledServices
		// }

		// // We don't want to keep asking to install the basic pipeline the entire time
		// if m.availableServices.HasData() && len(m.missingServicesInstallation) == 0 {
		// 	missingServices := make([]openapi.FullyQualifiedService, 0)
		// 	for _, official := range defaultPipeline {
		// 		found := false
		// 		for _, installed := range *m.availableServices.Data {
		// 			if installed.Name == official && installed.Author == "vu-ase" {
		// 				found = true
		// 			}
		// 		}

		// 		if !found {
		// 			missingServices = append(missingServices, openapi.FullyQualifiedService{
		// 				Name:   official,
		// 				Author: "vu-ase",
		// 			})
		// 		}
		// 	}
		// }
		// return m, m.renderPipelineGraph()
		return m, nil
	}

	if m.selectedIndex < 0 {
		m.selectedIndex = 0
	}

	return m, cmd
}

func (m PipelineManagerPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchPipeline(), m.fetchServices(), m.fetchDefaultServiceReleases(), textinput.Blink)
}

// status := statusStyle.Background(style.GrayPrimary).Bold(true).Render("unknown")
// if m.pipelineExecution.HasData() {
// 	if m.pipelineExecution.Data.Status == openapi.STARTED {
// 		status = statusStyle.Background(style.SuccessPrimary).Bold(true).Render("running")
// 	} else if m.pipelineExecution.Data.Status == openapi.STARTABLE {
// 		status = statusStyle.BorderForeground(style.WarningPrimary).Bold(true).Foreground(style.WarningPrimary).Render("startable")
// 	} else if m.pipelineExecution.Data.Status == openapi.EMPTY {
// 		status = statusStyle.Background(style.ErrorPrimary).Bold(true).Render("pipeline empty")
// 	}
// } else if m.pipelineExecution.IsError() {
// 	status = statusStyle.Background(style.ErrorPrimary).Bold(true).Render("could not execute")
// } else if m.pipelineExecution.IsLoading() {
// 	status = statusStyle.Background(style.GrayPrimary).Bold(true).Render(m.spinner.View() + " loading...")
// }

// subStatus := ""
// if m.pipelineExecution.IsError() {
// 	subStatus += style.Error.Render("! "+m.pipelineExecution.Error.Error()) + "\n\n"
// }
// if m.dirty {
// 	subStatus += style.Warning.Render("! Unsaved local changes") + "\n\n"
// }

// loader := " "
// if m.pipeline.IsLoading() || m.availableServices.IsLoading() || m.pipelineExecution.IsLoading() {
// 	loader += m.spinner.View()
// }

// // s := style.Title.Render("Execution pipeline") + " " + loader + "\n"
// s := ""

// leftStyle := lipgloss.NewStyle().
// 	Width(30).
// 	Padding(1, 0)
// statusStyle := lipgloss.NewStyle().Padding(1, 2).Width(28).Align(lipgloss.Center).Border(lipgloss.NormalBorder())

// // Style for the right column
// rightStyle := lipgloss.NewStyle().
// 	Padding(0, 2)

// // Always try to render the local pipeline first
// if m.pipeline.IsLoading() {
// 	subStatus += "\n" + m.spinner.View() + " Loading pipeline..." + "\n\n"
// } else if m.pipeline.IsError() {
// 	subStatus += "\n" + style.Error.Render("✗ Could not load pipeline") + style.Gray.Render(" ("+m.pipeline.Error.Error()+")") + "\n\n"
// }

func (m PipelineManagerPage) updateDialogView() string {
	return "not implemented dialog"
	// dialog := style.Primary.Bold(true).Render("Updates available") + "\n\n" + "The ASE autonomous driving pipeline has updates available. \nInstall them now?\n\n"

	// // Align to the left
	// leftAligned := ""
	// if len(m.missingServicesInstallation) > 0 {
	// 	for _, install := range m.missingServicesInstallation {
	// 		if install.IsLoading() {
	// 			leftAligned += m.spinner.View() + " Installing '" + install.Request().Name + "' v" + install.Request().LatestVersion
	// 		} else if install.IsError() {
	// 			leftAligned += style.Error.Render("✗ Could not install '"+install.Request().Name+"'") + ":\n"
	// 			for _, e := range install.Errors() {
	// 				leftAligned += "  > " + e.Error() + "\n"
	// 			}
	// 		} else {
	// 			leftAligned += style.Success.Render("✓ Installed '" + install.Request().Name + "' v" + install.Request().LatestVersion)
	// 		}
	// 		leftAligned += "\n"
	// 	}
	// } else {
	// 	for _, service := range m.defaultPipelineServices.Result() {
	// 		// Find the installed counterpart
	// 		var installed *utils.ServiceFqn
	// 		for _, i := range *m.availableServices.Data {
	// 			if i.Name == service.Name && i.Author == service.Author {
	// 				installed = &i
	// 			}
	// 		}

	// 		desc := style.Gray.Render(service.Author+"/") + lipgloss.NewStyle().Bold(true).Render(service.Name) + style.Gray.Render(" v"+service.LatestVersion)

	// 		if installed == nil {
	// 			leftAligned += style.Primary.Render("Install ") + desc
	// 		} else if installed.Version != service.LatestVersion {
	// 			leftAligned += style.Primary.Render("Update ") + desc
	// 		}
	// 		leftAligned += "\n"
	// 	}
	// }
	// dialog += lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).Render(leftAligned)

	// //
	// // The code below (and the structure of checking if actions are done)
	// // needs some majore DRY and cleanup. Just so you know.
	// //

	// // Are all actions done?
	// allDone := true
	// hasError := false
	// for _, action := range m.missingServicesInstallation {
	// 	if !action.IsDone() {
	// 		allDone = false
	// 	}
	// 	hasError = hasError || action.IsError()
	// }
	// if allDone && hasError {
	// 	dialog += "\n[r]etry failed [i]gnore"
	// }
	// if len(m.missingServicesInstallation) == 0 {
	// 	dialog += "\n[y]es [n]o"
	// }

	// return style.RenderDialog(dialog, style.AsePrimary)
}

func (m PipelineManagerPage) serviceListView(w int) string {
	selectorWidth := 3
	selectorStyle := lipgloss.NewStyle().Width(selectorWidth)
	width := w - selectorWidth
	nameStyle := lipgloss.NewStyle().Width(width / 10 * 5).Bold(true)
	versionStyle := lipgloss.NewStyle().Width(width / 10 * 1).Foreground(style.GrayPrimary).Align(lipgloss.Right)
	authorStyle := lipgloss.NewStyle().Width(width / 10 * 3).Foreground(style.GrayPrimary).PaddingLeft(1)
	faultsStyle := lipgloss.NewStyle().Width(width / 10 * 1)

	// Limit the height to 5 items
	services := m.services.Result()
	endIndex := m.selectedIndex + 1
	size := 10
	if endIndex < size {
		endIndex = size
	}
	if endIndex > len(services) {
		endIndex = len(services)
	}
	startIndex := endIndex - size
	if startIndex < 0 {
		startIndex = 0
	}
	services = services[startIndex:endIndex]

	serviceList := ""
	if m.services.HasResult() {
		for i, service := range services {
			isEnabled := m.isServiceEnabled(service)
			enabled := "\u00A0•\u00A0"
			if isEnabled {
				enabled = style.Success.Render(" ✔ ")
			}

			// Get the pipeline entry so that we can show faults
			var pipelineEntry *openapi.PipelineGet200ResponseEnabledInner = nil
			faults := "-"
			if m.pipeline.HasResult() {
				for _, e := range m.pipeline.Result().Enabled {
					if e.Service.Fq.Name == service.Name && e.Service.Fq.Author == service.Author && e.Service.Fq.Version == service.Version {
						pipelineEntry = &e
						break
					}
				}
			}
			if pipelineEntry != nil {
				faults = fmt.Sprintf("%d", pipelineEntry.Service.Faults)
			}

			name := service.Name
			if service.As != nil {
				name += " (" + *service.As + ")"
			}

			entry := selectorStyle.Render(enabled) + nameStyle.Render(name) + versionStyle.Render(service.Version) + authorStyle.Render(service.Author) + faultsStyle.Render(faults)
			if i+startIndex == m.selectedIndex {
				bg := style.AsePrimary
				if isEnabled {
					bg = style.SuccessPrimary
				}
				entry = selectorStyle.Background(bg).Render(enabled) + nameStyle.Background(bg).Render(name) + versionStyle.Background(bg).Render(service.Version) + authorStyle.Background(bg).Render(service.Author) + faultsStyle.Background(bg).Render(faults)
			}
			serviceList += entry + "\n"
		}

		if len(services) <= 0 {
			serviceList += style.Gray.Render("No services available")
		}
	} else if m.services.IsLoading() {
		serviceList += m.spinner.View() + " Loading available services..." + "\n\n"
	} else if m.services.IsError() {
		serviceList += style.Error.Render("✗ Could not load available services") + "\n\n"
	}

	suffix := ""
	if endIndex < len(services) {
		suffix = style.Gray.Render("  ...") + "\n"
	}

	return serviceList + suffix
}

func (m PipelineManagerPage) selectedServiceView(w int, h int) string {

	// width := w

	selectedService := m.selectedService()
	if selectedService == nil {
		return "None"
	}

	if m.selectedServiceDetails.HasResult() && m.selectedServiceLogs.HasResult() && utils.FqnsEqual(*selectedService, m.selectedServiceDetails.Request()) && utils.FqnsEqual(*selectedService, m.selectedServiceLogs.Request()) {
		// details := m.selectedServiceDetails.Result()
		logs := m.selectedServiceLogs.Result()

		logViewportStyle := lipgloss.NewStyle().Width(w-10).Height(h-10).Padding(1, 0)

		// Create the view
		view := ""

		view += logViewportStyle.Render(strings.Join(logs, "\n")) + "\n"

		// view += strings.Join(logs, "\n") + "\n"

		return view

	}

	return "ah"

	// Get the service

	// // See if this service is in the current pipeline
	// var pipelineEntry *openapi.PipelineGet200ResponseEnabledInner = nil
	// if m.pipeline.HasResult() {
	// 	for _, e := range m.pipeline.Result().Enabled {
	// 		if e.Service.Fq.Name == selectedService.Name && e.Service.Fq.Author == selectedService.Author && e.Service.Fq.Version == selectedService.Version {
	// 			pipelineEntry = &e
	// 			break
	// 		}
	// 	}
	// }
	// if pipelineEntry == nil {
	// 	return "No pipeline info"
	// }

	// process := pipelineEntry.Process
	// if process != nil {
	// 	// Compute how many = to print based on the CPU percentage
	// 	width -= 2
	// 	cpuBarWidth := int(float64(width) * float64(pipelineEntry.Process.Cpu) / 100)
	// 	cpuBar := strings.Repeat("=", cpuBarWidth)

	// 	return "CPU " + cpuBar + fmt.Sprintf("%d", process.Cpu)
	// } else {
	// 	return "No process info"
	// }
}

func (m PipelineManagerPage) View() string {
	// Ask to install/update basic pipeline
	if m.isDefaultPipelineUpdateAvailable() {
		return m.updateDialogView()
	}

	fullWidth := state.Get().WindowWidth - 4
	fullHeight := state.Get().WindowHeight - 4

	// Top banner that spans the entire width, to show pipeline status
	headerBg := style.GrayPrimary
	headerText := "unknown"
	if len(m.locallyEnabledServices) > 0 {
		headerBg = style.WarningPrimary
		headerText = "unsaved"
	} else if m.pipeline.HasResult() {
		switch m.pipeline.Result().Status {
		case openapi.STARTED:
			headerBg = style.SuccessPrimary
			headerText = "running"
		case openapi.STARTABLE:
			headerBg = style.AsePrimary
			headerText = "startable"
		}
	}

	headerStyle := lipgloss.NewStyle().Width(fullWidth).Align(lipgloss.Center).Padding(1, 0).Background(headerBg)
	header := ""
	if m.pipeline.IsLoading() {
		header = headerStyle.Render("Loading "+m.spinner.View()) + "\n"
	} else {
		header = headerStyle.Render("Your pipeline is "+style.Bold.Render(headerText)) + "\n"
	}

	// Create two columns: a left column (35%) and a right column (65%)
	// they should be separated by a vertical line
	columnHeight := fullHeight - 1 - len(strings.Split(header, "\n"))
	leftColumnWidth := fullWidth / 3
	leftColumn := lipgloss.NewStyle().Height(columnHeight).Width(leftColumnWidth).Padding(0, 2).Border(lipgloss.NormalBorder()).BorderLeft(false).BorderBottom(false).BorderTop(false).BorderForeground(style.GrayPrimary)
	rightColumnWidth := fullWidth - leftColumnWidth
	rightColumn := lipgloss.NewStyle().Width(rightColumnWidth).Padding(0, 4)

	// Render the left column
	left := leftColumn.Render(m.serviceListView(leftColumnWidth))
	right := rightColumn.Render(m.selectedServiceView(rightColumnWidth, columnHeight))

	return header + "\n" + lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func (m PipelineManagerPage) isQuitable() bool {
	return true
}

func (m PipelineManagerPage) keys() utils.GeneralKeyMap {
	kb := utils.NewGeneralKeyMap()
	kb.Back = key.NewBinding(
		key.WithKeys("ctrl+b"),
		key.WithHelp("ctrl+b", "back"),
	)

	// Disable keys if a dialog is shown
	if m.isDefaultPipelineUpdateAvailable() {
		return kb
	}

	kb.Up = key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	)
	kb.Down = key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	)
	kb.Save = key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save/start/stop"),
	)

	lenServices := len(m.services.Result())
	if lenServices > 0 {
		if m.selectedIndex >= 0 && m.selectedIndex < lenServices {
			selected := m.services.Result()[m.selectedIndex]
			if m.isServiceEnabled(selected) {
				kb.Toggle = key.NewBinding(
					key.WithKeys(" "),
					key.WithHelp("space", "disable"),
				)
			} else {
				kb.Toggle = key.NewBinding(
					key.WithKeys(" "),
					key.WithHelp("space", "enable"),
				)
			}
		}
	}

	kb.Details = key.NewBinding(
		key.WithKeys("ctrl+l"),
		key.WithHelp("ctrl+l", "service logs"),
	)

	return kb
}

func (m PipelineManagerPage) previousPage() *pageModel {
	var pageModel pageModel = NewStartPage()
	return &pageModel
}

//
// Shared conditionals
//

func (m PipelineManagerPage) isServiceEnabled(service openapi.FullyQualifiedService) bool {
	if len(m.locallyEnabledServices) > 0 {
		// Is locally enabled?
		for _, s := range m.locallyEnabledServices {
			if s.Name == service.Name && s.Author == service.Author && s.Version == service.Version {
				return true
			}
		}
	} else {
		// Is enabled in the pipeline
		if m.pipeline.HasResult() {
			for _, e := range m.pipeline.Result().Enabled {
				if service.Name == e.Service.Fq.Name && service.Author == e.Service.Fq.Author && service.Version == e.Service.Fq.Version {
					return true
				}
			}
		}
	}
	return false
}

// Compares the installed services to the available services for the default pipeline, and reports missing services
func (m PipelineManagerPage) isDefaultPipelineUpdateAvailable() bool {
	// if !m.availableServices.HasData() || !m.defaultPipelineServices.HasResult() || m.ignoreUpdates {
	// 	return false
	// }

	// for _, official := range m.defaultPipelineServices.Result() {
	// 	found := false
	// 	for _, installed := range *m.availableServices.Data {
	// 		if installed.Name == official.Name && installed.Author != official.Author && installed.Version == official.LatestVersion {
	// 			found = true
	// 		}
	// 	}

	// 	if !found {
	// 		return true
	// 	}
	// }

	return false
}

//
// (Remote) actions
//

func (m PipelineManagerPage) fetchPipeline() tea.Cmd {
	return tui.PerformActionV2(&m.pipeline, nil, func() (*openapi.PipelineGet200Response, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		api := remote.ToApiClient()
		status, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}
		return status, nil
	})
}

// Save the pipeline and start/stop it, based on the current status
func (m PipelineManagerPage) togglePipeline() tea.Cmd {
	return tui.PerformActionV2(&m.pipeline, nil, func() (*openapi.PipelineGet200Response, []error) {
		enabledServices := []openapi.PipelinePostRequestInner{}
		for _, service := range enabledServices {
			enabledServices = append(enabledServices, openapi.PipelinePostRequestInner{
				Fq: service.Fq,
			})
		}
		if len(enabledServices) == 0 {
			return nil, []error{fmt.Errorf("No services enabled")}
		}

		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}
		api := remote.ToApiClient()

		// First, get the status of the pipeline
		status, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		if status.Status == openapi.STARTED {
			// If the pipeline is running, stop it and do nothing else. Updating the pipeline
			// only happens when the rover is not running
			htt, err := api.PipelineAPI.PipelineStopPost(
				context.Background(),
			).Execute()

			if err != nil {
				return nil, []error{utils.ParseHTTPError(err, htt)}
			}
		} else if status.Status == openapi.STARTABLE || status.Status == openapi.EMPTY {
			// If the pipeline is startable, update the pipeline, build and start

			// Set the new pipeline
			req := api.PipelineAPI.PipelinePost(
				context.Background(),
			)
			req = req.PipelinePostRequestInner(enabledServices)
			htt, err = req.Execute()

			if err != nil {
				return nil, []error{utils.ParseHTTPError(err, htt)}
			}

			// Pipeline has been updated successfully, so we can build it
			// First, build all services
			// this is currently done very simple: it is not checked when the last build time was or if the services changed
			// in theory, if the services did not change, we should not need to build them again
			for _, service := range enabledServices {
				htt, err := api.ServicesAPI.ServicesAuthorServiceVersionPost(
					context.Background(),
					service.Fq.Author,
					service.Fq.Name,
					service.Fq.Version,
				).Execute()
				if err != nil {
					return nil, []error{fmt.Errorf("Failed to build service %s: %s", service.Fq.Name, utils.ParseHTTPError(err, htt))}
				}
			}

			htt, err := api.PipelineAPI.PipelineStartPost(
				context.Background(),
			).Execute()

			if err != nil {
				return nil, []error{utils.ParseHTTPError(err, htt)}
			}
		}

		// Finally, fetch the status again, to return
		status, htt, err = api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		return status, nil
	})
}

// Best effort stop pipeline, does not report errors if the pipeline is already stopped
func (m PipelineManagerPage) stopPipeline() tea.Cmd {
	return tui.PerformActionV2(&m.pipeline, nil, func() (*openapi.PipelineGet200Response, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		// Stop the pipeline
		api := remote.ToApiClient()
		_, _ = api.PipelineAPI.PipelineStopPost(
			context.Background(),
		).Execute()
		// nb: errors are ignored

		// Finally, fetch the status again, to return
		status, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		return status, nil
	})
}

func (m PipelineManagerPage) fetchServices() tea.Cmd {
	return tui.PerformActionV2(&m.services, nil, func() (*[]openapi.FullyQualifiedService, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		// Fetch all authors
		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.FqnsGet(
			context.Background(),
		).Execute()
		if err != nil && htt != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		return &res, nil
	})
}

func (m PipelineManagerPage) fetchDefaultServiceReleases() tea.Cmd {
	return tui.PerformActionV2(&m.defaultPipelineServices, nil, func() (*[]utils.UpdateAvailable, []error) {
		releases := make([]utils.UpdateAvailable, 0)

		for _, official := range defaultPipeline {
			service, err := utils.CheckForGithubUpdate(official, "VU-ASE", "none")
			if err != nil {
				return nil, []error{err}
			} else if service == nil {
				return nil, []error{fmt.Errorf("Service %s not found", official)}
			}
			releases = append(releases, *service)
		}

		return &releases, nil
	})
}

func (m PipelineManagerPage) fetchServiceLogs(service openapi.FullyQualifiedService) tea.Cmd {
	return tui.PerformActionV2(&m.selectedServiceLogs, &service, func() (*[]string, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		// Fetch logs for provided service
		api := remote.ToApiClient()
		res, htt, err := api.PipelineAPI.LogsAuthorNameVersionGet(
			context.Background(),
			service.Author,
			service.Name,
			service.Version,
		).Execute()
		if err != nil && htt != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		return &res, nil
	})
}

func (m PipelineManagerPage) fetchServiceDetails(service openapi.FullyQualifiedService) tea.Cmd {
	return tui.PerformActionV2(&m.selectedServiceDetails, &service, func() (*openapi.ServicesAuthorServiceVersionGet200Response, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		// Fetch logs for provided service
		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorServiceVersionGet(
			context.Background(),
			service.Author,
			service.Name,
			service.Version,
		).Execute()
		if err != nil && htt != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		return res, nil
	})
}

func (m PipelineManagerPage) installService(action *tui.ActionV2[utils.UpdateAvailable, openapi.FetchPost200Response], service utils.UpdateAvailable) tea.Cmd {
	return tui.PerformActionV2(action, &service, func() (*openapi.FetchPost200Response, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		api := remote.ToApiClient()

		// Find the URL to install
		url := ""
		if len(service.Assets) > 1 {
			// Pick the first zip file asset
			for _, asset := range service.Assets {
				if asset.ContentType == "application/zip" {
					url = asset.Url
					break
				}
			}
		} else if len(service.Assets) == 1 {
			url = service.Assets[0].Url
		}
		if url == "" {
			return nil, []error{fmt.Errorf("No install candidate (zip file asset) found for service %s", service.Name)}
		}

		// Fetch the pipeline
		req := api.ServicesAPI.FetchPost(
			context.Background(),
		)
		body := openapi.FetchPostRequest{
			Url: service.Assets[0].Url,
		}
		req = req.FetchPostRequest(body)
		res, htt, err := req.Execute()
		if err != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		fqn := openapi.FullyQualifiedService{
			Name:    res.Fq.Name,
			Author:  res.Fq.Author,
			Version: res.Fq.Version,
		}
		if fqn.Name != service.Name {
			return nil, []error{fmt.Errorf("Roverd failed to install service %s. It installed %s instead", fqn.Name, service.Name)}
		} else if fqn.Version != service.LatestVersion {
			return nil, []error{fmt.Errorf("Roverd failed to install service %s@%s. It installed %s@%s instead", fqn.Name, service.LatestVersion, service.Name, fqn.Version)}
		} else if !strings.EqualFold(fqn.Author, service.Author) {
			return nil, []error{fmt.Errorf("Roverd failed to install service %s@%s by %s. It installed %s@%s by %s instead", fqn.Name, service.LatestVersion, service.Author, service.Name, fqn.Version, fqn.Author)}
		}

		return res, nil
	})
}

func (m PipelineManagerPage) selectedService() *openapi.FullyQualifiedService {
	if m.selectedIndex >= 0 && m.selectedIndex < len(m.services.Result()) {
		return &m.services.Result()[m.selectedIndex]
	}
	return nil
}
