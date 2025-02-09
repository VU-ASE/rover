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
	"github.com/lempiy/dgraph"
	"github.com/lempiy/dgraph/core"
)

var defaultPipeline = []string{
	"imaging", "controller", "actuator",
}

//
// The page model
//

type PipelineManagerPage struct {
	spinner spinner.Model

	// The pipeline as first loaded from the Rover, and later maybe edited
	pipeline                   tui.Action[[]utils.ServiceFqn]             // list of enabled services
	availableServices          tui.Action[[]utils.ServiceFqn]             // list of *all* services that can be enabled
	pipelineGraph              tui.Action[string]                         // preserved in the model to avoid re-rendering in .View(), based on the latest service information
	pipelineExecution          tui.Action[openapi.PipelineGet200Response] // Will save and toggle the execution of the current pipeline, returning any errors that might occur
	processedAvailableServices []PipelineService                          // faster method of computing the available services

	// Filtering through the list of available services
	filterValue   textinput.Model
	filterEnabled bool
	selectedIndex int
	dirty         bool // whether the pipeline has been edited
	ignoreUpdates bool

	// Actions to install the basic pipeline
	defaultPipelineServices     tui.ActionV2[any, []utils.UpdateAvailable]                           // all latest services in the default autonomous driving pipeline
	missingServicesInstallation []*tui.ActionV2[utils.UpdateAvailable, openapi.FetchPost200Response] // installation if missing services (if started)
}

type PipelineService struct {
	service utils.ServiceFqn
	enabled bool
}

type ServiceDetails struct {
	service utils.ServiceFqn
	details openapi.ServicesAuthorServiceVersionGet200Response
}

func NewPipelineManagerPage() PipelineManagerPage {
	// Create actions
	pl := tui.NewAction[[]utils.ServiceFqn]("pipeline")
	as := tui.NewAction[[]utils.ServiceFqn]("availableServices")
	pg := tui.NewAction[string]("pipelineGraph")
	te := tui.NewAction[openapi.PipelineGet200Response]("toggleExecution")

	ti := textinput.New()
	ti.Placeholder = "Type to filter services..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 100

	return PipelineManagerPage{
		spinner:           spinner.New(),
		pipeline:          pl,
		availableServices: as,
		pipelineGraph:     pg,
		pipelineExecution: te,
		filterValue:       ti,
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
		case msg.String() == "y" && m.defaultPipelineServicesMissing() && len(m.missingServicesInstallation) == 0:
			for _, service := range m.defaultPipelineServices.Result() {
				// Find the installed counterpart
				var installed *utils.ServiceFqn
				for _, i := range *m.availableServices.Data {
					if i.Name == service.Name && i.Author == service.Author {
						installed = &i
					}
				}

				if installed == nil || installed.Version != service.LatestVersion {
					// Create a new action and start it
					action := tui.NewActionV2[utils.UpdateAvailable, openapi.FetchPost200Response]()
					m.missingServicesInstallation = append(m.missingServicesInstallation, &action)
					cmd = tea.Batch(cmd, m.installService(&action, service))
				}
			}
			return m, cmd
		case msg.String() == "r" && m.defaultPipelineServicesMissing() && len(m.missingServicesInstallation) > 0:
			// All actions need to be done
			for _, action := range m.missingServicesInstallation {
				if !action.IsDone() {
					return m, nil
				}
			}

			// Reset the actions
			m.missingServicesInstallation = make([]*tui.ActionV2[utils.UpdateAvailable, openapi.FetchPost200Response], 0)
			for _, service := range m.defaultPipelineServices.Result() {
				// Find the installed counterpart
				var installed *utils.ServiceFqn
				for _, i := range *m.availableServices.Data {
					if i.Name == service.Name && i.Author == service.Author {
						installed = &i
					}
				}

				if installed == nil || installed.Version != service.LatestVersion {
					// Create a new action and start it
					action := tui.NewActionV2[utils.UpdateAvailable, openapi.FetchPost200Response]()
					m.missingServicesInstallation = append(m.missingServicesInstallation, &action)
					cmd = tea.Batch(cmd, m.installService(&action, service))
				}
			}
			return m, cmd
		case msg.String() == "n", msg.String() == "i":
			// All actions need to be done
			for _, action := range m.missingServicesInstallation {
				if !action.IsDone() {
					return m, nil
				}
			}

			m.ignoreUpdates = true
			return m, nil

		case key.Matches(msg, m.keys().Back):
			return m, tea.Quit
		case key.Matches(msg, m.keys().Up):
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case key.Matches(msg, m.keys().Down):
			if m.selectedIndex < len(m.processedAvailableServices)-1 {
				m.selectedIndex++
			}
		case key.Matches(msg, m.keys().Toggle):
			if m.selectedIndex >= 0 && m.selectedIndex < len(m.processedAvailableServices) && m.pipeline.HasData() {
				newData := *m.pipeline.Data
				// Toggle the service
				s := m.processedAvailableServices[m.selectedIndex]
				if s.enabled {
					// Remove the service
					newData = make([]utils.ServiceFqn, 0)
					for _, e := range *m.pipeline.Data {
						if e.Name != s.service.Name || e.Author != s.service.Author || e.Version != s.service.Version {
							newData = append(newData, e)
						}
					}
				} else {
					// If there is already a service with this name (only name), remove it first
					newData = make([]utils.ServiceFqn, 0)
					for _, e := range *m.pipeline.Data {
						if e.Name != s.service.Name {
							newData = append(newData, e)
						}
					}
					newData = append(newData, s.service)
				}

				// Modify original pipeline data
				m.pipeline.Data = &newData
				m.dirty = true
				m.processAvailableServices()
				return m, m.stopPipelineExecution()
			}
		case key.Matches(msg, m.keys().Save):
			if !m.pipelineExecution.IsLoading() {
				return m, m.togglePipelineExecution()
			} else {
				return m, nil
			}
		case key.Matches(msg, m.keys().Details):
			if m.selectedIndex >= 0 && m.selectedIndex < len(m.processedAvailableServices) {
				s := m.processedAvailableServices[m.selectedIndex]
				return RootScreen(state.Get()).SwitchScreen(NewPipelineLogsPage(s.service.Name, s.service.Author, s.service.Version))
			}
		case key.Matches(msg, m.keys().Configure):
			m.filterEnabled = !m.filterEnabled
			m.processAvailableServices()
			return m, nil
		}

	// Action catchers
	case tui.ActionUpdate[any, any]:
		m.defaultPipelineServices.ProcessUpdate(msg)
		for _, action := range m.missingServicesInstallation {
			action.ProcessUpdate(msg)
		}
		return m, nil
	case tui.ActionInit[[]utils.ServiceFqn]:
		m.pipeline.ProcessInit(msg)
		m.availableServices.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[[]utils.ServiceFqn]:
		m.pipeline.ProcessResult(msg)
		m.availableServices.ProcessResult(msg)
		m.processAvailableServices()

		// We don't want to keep asking to install the basic pipeline the entire time
		if m.availableServices.HasData() && len(m.missingServicesInstallation) == 0 {
			missingServices := make([]openapi.FullyQualifiedService, 0)
			for _, official := range defaultPipeline {
				found := false
				for _, installed := range *m.availableServices.Data {
					if installed.Name == official && installed.Author == "vu-ase" {
						found = true
					}
				}

				if !found {
					missingServices = append(missingServices, openapi.FullyQualifiedService{
						Name:   official,
						Author: "vu-ase",
					})
				}
			}
		}
		return m, m.renderPipelineGraph()
	case tui.ActionInit[string]:
		m.pipelineGraph.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[string]:
		m.pipelineGraph.ProcessResult(msg)
		return m, nil
	case tui.ActionInit[openapi.PipelineGet200Response]:
		m.pipelineExecution.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[openapi.PipelineGet200Response]:
		m.pipelineExecution.ProcessResult(msg)
		if m.pipelineExecution.HasData() && m.pipelineExecution.Data.Status == openapi.STARTED {
			m.dirty = false
		}
		return m, nil
	case tui.ActionInit[[]openapi.FetchPost200Response]:
		return m, nil
	case tui.ActionResult[[]openapi.FetchPost200Response]:
		return m, tea.Batch(m.fetchRemoteServices(), m.fetchRemotePipeline())
	}

	if m.selectedIndex < 0 {
		m.selectedIndex = 0
	}

	// m.filterValue, cmd = m.filterValue.Update(msg)
	m.processAvailableServices()
	return m, cmd
}

func (m PipelineManagerPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchRemotePipeline(), m.fetchRemoteServices(), m.fetchExecutionStatus(), m.fetchDefaultServiceReleases(), textinput.Blink)
}

func (m PipelineManagerPage) View() string {
	// Ask to install basic pipeline
	if m.defaultPipelineServicesMissing() {
		dialog := style.Primary.Bold(true).Render("Updates available") + "\n\n" + "The ASE autonomous driving pipeline has updates available. \nInstall them now?\n\n"

		// Align to the left
		leftAligned := ""
		if len(m.missingServicesInstallation) > 0 {
			for _, install := range m.missingServicesInstallation {
				if install.IsLoading() {
					leftAligned += m.spinner.View() + " Installing '" + install.Request().Name + "' v" + install.Request().LatestVersion
				} else if install.IsError() {
					leftAligned += style.Error.Render("✗ Could not install '"+install.Request().Name+"'") + ":\n"
					for _, e := range install.Errors() {
						leftAligned += "  > " + e.Error() + "\n"
					}
				} else {
					leftAligned += style.Success.Render("✓ Installed '" + install.Request().Name + "' v" + install.Request().LatestVersion)
				}
				leftAligned += "\n"
			}
		} else {
			for _, service := range m.defaultPipelineServices.Result() {
				// Find the installed counterpart
				var installed *utils.ServiceFqn
				for _, i := range *m.availableServices.Data {
					if i.Name == service.Name && i.Author == service.Author {
						installed = &i
					}
				}

				desc := style.Gray.Render(service.Author+"/") + lipgloss.NewStyle().Bold(true).Render(service.Name) + style.Gray.Render(" v"+service.LatestVersion)

				if installed == nil {
					leftAligned += style.Primary.Render("Install ") + desc
				} else if installed.Version != service.LatestVersion {
					leftAligned += style.Primary.Render("Update ") + desc
				}
				leftAligned += "\n"
			}
		}
		dialog += lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).Render(leftAligned)

		//
		// The code below (and the structure of checking if actions are done)
		// needs some majore DRY and cleanup. Just so you know.
		//

		// Are all actions done?
		allDone := true
		hasError := false
		for _, action := range m.missingServicesInstallation {
			if !action.IsDone() {
				allDone = false
			}
			hasError = hasError || action.IsError()
		}
		if allDone && hasError {
			dialog += "\n[r]etry failed [i]gnore"
		}
		if len(m.missingServicesInstallation) == 0 {
			dialog += "\n[y]es [n]o"
		}

		return style.RenderDialog(dialog, style.AsePrimary)
	}

	loader := " "
	if m.pipeline.IsLoading() || m.availableServices.IsLoading() || m.pipelineExecution.IsLoading() {
		loader += m.spinner.View()
	}

	s := style.Title.Render("Execution pipeline") + " " + loader + "\n"

	serviceList := "\n"

	if m.availableServices.HasData() {
		// s += m.filterValue.View()[2:] + "\n\n"
		for i, service := range m.processedAvailableServices {
			selector := style.Gray.Render("[ ] ")
			if service.enabled {
				selector = "[" + style.Success.Render("✔") + "] "
			}
			prefix := "   " + selector
			if i == m.selectedIndex {
				prefix = " > " + selector
			}

			serviceList += prefix + style.Gray.Render(service.service.Author+"/") + lipgloss.NewStyle().Bold(true).Render(service.service.Name) + style.Gray.Render(" v"+service.service.Version) + "\n"
		}

		if len(m.processedAvailableServices) <= 0 {
			serviceList += style.Gray.Render("No services available")
		}
	} else if m.availableServices.IsLoading() {
		serviceList += m.spinner.View() + " Loading available services..." + "\n\n"
	} else if m.availableServices.IsError() {
		serviceList += style.Error.Render("✗ Could not load available services") + style.Gray.Render(" ("+m.availableServices.Error.Error()+")") + "\n\n"
	}

	leftStyle := lipgloss.NewStyle().
		Width(30).
		Padding(1, 0)
	statusStyle := lipgloss.NewStyle().Padding(1, 2).Width(28).Align(lipgloss.Center)

	// Style for the right column
	rightStyle := lipgloss.NewStyle().
		Padding(0, 2)

	status := statusStyle.Background(style.GrayPrimary).Bold(true).Render("unknown")
	if m.pipelineExecution.HasData() {
		if m.pipelineExecution.Data.Status == openapi.STARTED {
			status = statusStyle.Background(style.SuccessPrimary).Bold(true).Render("running")
		} else if m.pipelineExecution.Data.Status == openapi.STARTABLE {
			status = statusStyle.Background(style.WarningPrimary).Bold(true).Render("startable")
		} else if m.pipelineExecution.Data.Status == openapi.EMPTY {
			status = statusStyle.Background(style.ErrorPrimary).Bold(true).Render("pipeline empty")
		}
	} else if m.pipelineExecution.IsError() {
		status = statusStyle.Background(style.ErrorPrimary).Bold(true).Render("could not execute")
	} else if m.pipelineExecution.IsLoading() {
		status = statusStyle.Background(style.GrayPrimary).Bold(true).Render(m.spinner.View() + " loading...")
	}

	subStatus := ""
	if m.pipelineExecution.IsError() {
		subStatus += style.Error.Render("! "+m.pipelineExecution.Error.Error()) + "\n\n"
		// subStatus += "\n" + style.Error.Render("✗ Could not toggle execution") + style.Gray.Render(" ("++")") + "\n\n"
	}
	if m.dirty {
		subStatus += style.Warning.Render("! Unsaved local changes") + "\n\n"
	}

	// Always try to render the local pipeline first
	if m.pipeline.IsLoading() {
		subStatus += "\n" + m.spinner.View() + " Loading pipeline..." + "\n\n"
	} else if m.pipeline.IsError() {
		subStatus += "\n" + style.Error.Render("✗ Could not load pipeline") + style.Gray.Render(" ("+m.pipeline.Error.Error()+")") + "\n\n"
	}

	return s + lipgloss.JoinHorizontal(lipgloss.Top, leftStyle.Render(status+"\n\n"+subStatus), rightStyle.Render(serviceList))

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
	if m.defaultPipelineServicesMissing() {
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
	if m.filterEnabled {
		kb.Configure = key.NewBinding(
			key.WithKeys("ctrl+e"),
			key.WithHelp("ctrl+e", "show all"),
		)
	} else {
		kb.Configure = key.NewBinding(
			key.WithKeys("ctrl+e"),
			key.WithHelp("ctrl+e", "show only enabled"),
		)
	}

	if len(m.processedAvailableServices) > 0 {
		if m.selectedIndex >= 0 && m.selectedIndex < len(m.processedAvailableServices) {
			selected := m.processedAvailableServices[m.selectedIndex]
			if selected.enabled {
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
// (Remote) actions
//

func (m PipelineManagerPage) fetchRemotePipeline() tea.Cmd {
	return tui.PerformAction(&m.pipeline, func() (*[]utils.ServiceFqn, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()

		// First, fetch all services and the status of the current pipeline
		pipeline, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		enabled := make([]utils.ServiceFqn, 0)
		for _, e := range pipeline.Enabled {
			enabled = append(enabled, utils.ServiceFqn{
				Name:    e.Service.Fq.Name,
				Author:  e.Service.Fq.Author,
				Version: e.Service.Fq.Version,
			})
		}
		return &enabled, nil
	})
}

func (m PipelineManagerPage) fetchExecutionStatus() tea.Cmd {
	return tui.PerformAction(&m.pipelineExecution, func() (*openapi.PipelineGet200Response, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()

		// First, fetch all services and the status of the current pipeline
		status, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return status, nil
	})
}

// Save the pipeline and start/stop it, based on the current status
func (m PipelineManagerPage) togglePipelineExecution() tea.Cmd {
	return tui.PerformAction(&m.pipelineExecution, func() (*openapi.PipelineGet200Response, error) {
		if len(*m.pipeline.Data) <= 0 {
			return nil, fmt.Errorf("No services enabled")
		}

		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()

		// First, get the status of the pipeline
		status, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		if status.Status == openapi.STARTED {
			// If the pipeline is running, stop it and do nothing else. Updating the pipeline
			// only happens when the rover is not running
			htt, err := api.PipelineAPI.PipelineStopPost(
				context.Background(),
			).Execute()

			if err != nil {
				return nil, utils.ParseHTTPError(err, htt)
			}
		} else if status.Status == openapi.STARTABLE || status.Status == openapi.EMPTY {
			// If the pipeline is startable, update the pipeline, build and start

			// Set the new pipeline
			req := api.PipelineAPI.PipelinePost(
				context.Background(),
			)

			pipelineReq := []openapi.PipelinePostRequestInner{}
			if m.pipeline.HasData() {
				for _, service := range *m.pipeline.Data {
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
			htt, err = req.Execute()

			if err != nil {
				if htt != nil {
					// Try to parse as a unmet stream error:
					httpRes := make([]byte, htt.ContentLength)
					_, err = htt.Body.Read(httpRes)
					deps, err := utils.TransformValidationError(string(httpRes))
					if err == nil {
						return nil, fmt.Errorf("Some services have unmet inputs:\n - %s", strings.Join(deps, "\n - "))
					}
				}
				return nil, utils.ParseHTTPError(err, htt)
			}

			// Pipeline has been updated successfully, so we can build it
			if len(*m.pipeline.Data) > 0 {
				// First, build all services
				// this is currently done very simple: it is not checked when the last build time was or if the services changed
				// in theory, if the services did not change, we should not need to build them again
				for _, service := range *m.pipeline.Data {
					htt, err := api.ServicesAPI.ServicesAuthorServiceVersionPost(
						context.Background(),
						service.Author,
						service.Name,
						service.Version,
					).Execute()
					if err != nil {
						return nil, fmt.Errorf("Failed to build service %s: %s", service.Name, utils.ParseHTTPError(err, htt))
					}
				}

			}

			htt, err := api.PipelineAPI.PipelineStartPost(
				context.Background(),
			).Execute()

			if err != nil {
				return nil, utils.ParseHTTPError(err, htt)
			}
		}

		// Finally, fetch the status again, to return
		status, htt, err = api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return status, nil
	})
}

// Best effort stop pipeline, does not report errors if the pipeline is already stopped
func (m PipelineManagerPage) stopPipelineExecution() tea.Cmd {
	return tui.PerformAction(&m.pipelineExecution, func() (*openapi.PipelineGet200Response, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
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
			return nil, utils.ParseHTTPError(err, htt)
		}

		return status, nil
	})
}

func (m PipelineManagerPage) fetchRemoteServices() tea.Cmd {
	return tui.PerformAction(&m.availableServices, func() (*[]utils.ServiceFqn, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		// Fetch all authors
		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesGet(
			context.Background(),
		).Execute()
		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		// Fetch all services for each author
		installed := make([]utils.ServiceFqn, 0)
		for _, author := range res {
			services, htt, err := api.ServicesAPI.ServicesAuthorGet(
				context.Background(),
				author,
			).Execute()
			if err != nil && htt != nil {
				return nil, utils.ParseHTTPError(err, htt)
			}

			// Fetch all versions for each service
			for _, service := range services {
				versions, htt, err := api.ServicesAPI.ServicesAuthorServiceGet(
					context.Background(),
					author,
					service,
				).Execute()
				if err != nil && htt != nil {
					return nil, utils.ParseHTTPError(err, htt)
				}

				for _, version := range versions {
					installed = append(installed, utils.ServiceFqn{
						Name:    service,
						Author:  author,
						Version: version,
					})
				}
			}
		}

		return &installed, err
	})
}

// Fetches service details and then creates a nice pipeline graph
func (m PipelineManagerPage) renderPipelineGraph() tea.Cmd {
	return tui.PerformAction(&m.pipelineGraph, func() (*string, error) {
		if !m.pipeline.HasData() {
			return nil, fmt.Errorf("No pipeline data")
		}

		// First fetch all the service details of the enabled services
		enabledDetails := make([]ServiceDetails, 0)
		for _, service := range *m.pipeline.Data {
			details, htt, err := state.Get().RoverConnections.GetActive().ToApiClient().ServicesAPI.ServicesAuthorServiceVersionGet(
				context.Background(),
				service.Author,
				service.Name,
				service.Version,
			).Execute()
			if err != nil && htt != nil {
				return nil, utils.ParseHTTPError(err, htt)
			}
			enabledDetails = append(enabledDetails, ServiceDetails{
				service: service,
				details: *details,
			})
		}

		res := ""
		// Create the pipeline graph based on enabled services
		nodes := make([]core.NodeInput, 0)
		for _, service := range *m.pipeline.Data {
			nodes = append(nodes, core.NodeInput{
				Id: service.Name,
				Next: func() []string {
					// Find services that depend on an output of this service
					found := make([]string, 0)
					for _, s := range enabledDetails {
						if s.service.Name != service.Name {
							for _, input := range s.details.Inputs {
								if input.Service == service.Name {
									found = append(found, s.service.Name)
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
			res = style.Gray.Render("This pipeline is empty")
		} else if err != nil {
			res = "Failed to draw pipeline\n"
		} else {
			res = fmt.Sprintf("%s\n", canvas)
		}

		return &res, nil
	})
}

// Compares the installed services to the available services for the default pipeline, and reports missing services
func (m PipelineManagerPage) defaultPipelineServicesMissing() bool {
	if !m.availableServices.IsSuccess() || !m.defaultPipelineServices.IsSuccess() || m.ignoreUpdates {
		return false
	}

	for _, official := range m.defaultPipelineServices.Result() {
		found := false
		for _, installed := range *m.availableServices.Data {
			if installed.Name == official.Name && installed.Author != official.Author && installed.Version == official.LatestVersion {
				found = true
			}
		}

		if !found {
			return true
		}
	}

	return false

}

// Filters the available services based on the filter value and adds the enabled services
func (m *PipelineManagerPage) processAvailableServices() {
	if !m.availableServices.HasData() || !m.pipeline.HasData() {
		m.processedAvailableServices = make([]PipelineService, 0)
		return
	}

	processed := make([]PipelineService, 0)
	// First add all enabled services
	for _, service := range *m.availableServices.Data {
		// Skip illegal services
		// todo: this is a hotfix, make it proper
		if service.Name == "transceiver" {
			continue
		}

		// Does it match the filter?
		if m.filterValue.Value() != "" && !strings.Contains(strings.ToLower(service.Name), strings.ToLower(m.filterValue.Value())) && !strings.Contains(strings.ToLower(service.Author), strings.ToLower(m.filterValue.Value())) && !strings.Contains(strings.ToLower(service.Version), strings.ToLower(m.filterValue.Value())) {
			continue
		}

		enabled := false
		for _, e := range *m.pipeline.Data {
			if e.Name == service.Name && e.Author == service.Author && e.Version == service.Version {
				enabled = true
				break
			}
		}

		if !enabled && m.filterEnabled {
			continue
		}
		processed = append(processed, PipelineService{
			service: service,
			enabled: enabled,
		})
	}

	m.processedAvailableServices = processed
	if m.selectedIndex >= len(m.processedAvailableServices) {
		m.selectedIndex = len(m.processedAvailableServices) - 1
	}
}

func (m PipelineManagerPage) fetchDefaultServiceReleases() tea.Cmd {
	return tui.PerformActionV2(&m.defaultPipelineServices, nil, func() (*[]utils.UpdateAvailable, []error) {
		releases := make([]utils.UpdateAvailable, 0)

		releases = []utils.UpdateAvailable{
			{
				Name:          "imaging",
				Author:        "VU-ASE",
				LatestVersion: "0.0.1",
			},
			{
				Name:          "controller",
				Author:        "VU-ASE",
				LatestVersion: "0.0.1",
			},
			{
				Name:          "actuator",
				Author:        "VU-ASE",
				LatestVersion: "0.0.1",
			},
		}

		// for _, official := range defaultPipeline {
		// 	service, err := utils.CheckForGithubUpdate(official, "VU-ASE", "none")
		// 	if err != nil {
		// 		return nil, []error{err}
		// 	} else if service == nil {
		// 		return nil, []error{fmt.Errorf("Service %s not found", official)}
		// 	}
		// 	releases = append(releases, *service)
		// }

		return &releases, nil
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
