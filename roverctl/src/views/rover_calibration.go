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
	tea "github.com/charmbracelet/bubbletea"
)

//
// The page model
//

type RoverCalibrationPage struct {
	spinner spinner.Model

	// Actions
	// The pipeline that should be restored after calibration
	previousPipeline tui.ActionV2[any, openapi.PipelineGet200Response]
	// All installed services on the Rover
	installedServices tui.ActionV2[any, []openapi.FqnsGet200ResponseInner]
	// All services that should make the calibration pipeline as available on Github
	calibrationReleases tui.ActionV2[any, []utils.UpdateAvailable]
	// Action to install a missing service
	serviceInstallations []*tui.ActionV2[utils.UpdateAvailable, openapi.FqnsGet200ResponseInner]
}

func NewRoverCalibrationPage() RoverCalibrationPage {
	// todo

	return RoverCalibrationPage{
		spinner: spinner.New(),
		// Actions
		previousPipeline:     tui.NewActionV2[any, openapi.PipelineGet200Response](),
		installedServices:    tui.NewActionV2[any, []openapi.FqnsGet200ResponseInner](),
		calibrationReleases:  tui.NewActionV2[any, []utils.UpdateAvailable](),
		serviceInstallations: []*tui.ActionV2[utils.UpdateAvailable, openapi.FqnsGet200ResponseInner]{},
	}
}

//
// Page model methods
//

func (m RoverCalibrationPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tui.ActionUpdate[any, any]:
		m.previousPipeline.ProcessUpdate(msg)
		m.installedServices.ProcessUpdate(msg)
		m.calibrationReleases.ProcessUpdate(msg)
		for _, action := range m.serviceInstallations {
			action.ProcessUpdate(msg)
		}

		if tui.AllSuccess(m.installedServices, m.calibrationReleases) && len(m.serviceInstallations) == 0 {
			// We can now check which services are missing
			missing := make([]utils.UpdateAvailable, 0)

			for _, service := range m.calibrationReleases.Result() {
				found := false
				for _, installed := range m.installedServices.Result() {
					if installed.Name == service.Name && strings.TrimPrefix(installed.Version, "v") == strings.TrimPrefix(service.LatestVersion, "v") && strings.EqualFold(installed.Author, service.Author) {
						found = true
						break
					}
				}
				if !found {
					missing = append(missing, service)
				}
			}

			var cmds tea.Cmd
			for _, service := range missing {
				// Create a new action and start it
				action := tui.NewActionV2[utils.UpdateAvailable, openapi.FqnsGet200ResponseInner]()
				m.serviceInstallations = append(m.serviceInstallations, &action)
				cmds = tea.Batch(cmds, m.installService(&action, service))
			}

			return m, cmds
		}

		return m, nil
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch {
		case msg.String() == "r":
			return m, m.fetchPreviousPipeline()
		case key.Matches(msg, m.keys().Back):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m RoverCalibrationPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchPreviousPipeline(), m.fetchCalibrationReleases(), m.fetchInstalledServices())
}

func (m RoverCalibrationPage) View() string {
	s := style.Title.Render("Calibration") + "\n\n"

	if m.previousPipeline.IsLoading() {
		s += m.spinner.View() + " Fetching previous pipeline \n"
	} else if m.previousPipeline.IsSuccess() {
		if m.previousPipeline.Result().Status == openapi.STARTED {
			s += style.Error.Render("✗ Previous pipeline is still running. Please stop it before starting calibration") + "\n"
		} else {
			s += style.Success.Render("✓ Fetched previous pipeline") + "\n"
		}
	} else if m.previousPipeline.IsError() {
		s += style.Warning.Render("! Failed to fetch previous pipeline") + "\n"
		for _, err := range m.previousPipeline.Errors() {
			s += "  > " + err.Error() + "\n"
		}
	}

	if m.calibrationReleases.IsLoading() {
		s += m.spinner.View() + " Fetching calibration releases\n"
	} else if m.calibrationReleases.IsSuccess() {
		s += style.Success.Render("✓ Fetched calibration releases") + "\n"
	} else if m.calibrationReleases.IsError() {
		s += style.Warning.Render("! Failed to fetch calibration releases") + "\n"
		for _, err := range m.calibrationReleases.Errors() {
			s += "  > " + err.Error() + "\n"
		}
	}

	if m.installedServices.IsLoading() {
		s += m.spinner.View() + " Fetching installed services\n"
	} else if m.installedServices.IsSuccess() {
		comment := ""
		if len(m.serviceInstallations) > 0 {
			comment = ". Missing services will be installed"
		}

		s += style.Success.Render("✓ Fetched installed services"+comment) + "\n"
	} else if m.installedServices.IsError() {
		s += style.Warning.Render("! Failed to fetch installed services") + "\n"
		for _, err := range m.installedServices.Errors() {
			s += "  > " + err.Error() + "\n"
		}
	}

	for _, action := range m.serviceInstallations {
		fq := action.Request().Name + "@v" + action.Request().LatestVersion
		if action.IsLoading() {
			s += m.spinner.View() + " Installing " + fq + "\n"
		} else if action.IsSuccess() {
			fq := action.Result().Name + "@v" + action.Result().Version
			s += style.Success.Render("✓ Installed "+fq) + "\n"
		} else if action.IsError() {
			s += style.Error.Render("✗ Failed to install "+fq) + "\n"
			for _, err := range action.Errors() {
				s += "  > " + err.Error() + "\n"
			}
		} else {
			s += style.Warning.Render("! Unknown state for service installation") + "\n"
		}
	}

	return s
}

func (m RoverCalibrationPage) isQuitable() bool {
	return true
}

func (m RoverCalibrationPage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m RoverCalibrationPage) previousPage() *pageModel {
	return nil
}

//
// Actions
//

func (m RoverCalibrationPage) fetchPreviousPipeline() tea.Cmd {
	return tui.PerformActionV2(&m.previousPipeline, nil, func() (*openapi.PipelineGet200Response, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		api := remote.ToApiClient()

		// Try (best-effort, don't check errors) to stop the current pipeline
		_, _ = api.PipelineAPI.PipelineStopPost(
			context.Background(),
		).Execute()

		// Fetch the pipeline
		pipeline, htt, err := api.PipelineAPI.PipelineGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		return pipeline, nil
	})
}

func (m RoverCalibrationPage) fetchInstalledServices() tea.Cmd {
	return tui.PerformActionV2(&m.installedServices, nil, func() (*[]openapi.FqnsGet200ResponseInner, []error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, []error{fmt.Errorf("No active rover connection")}
		}

		api := remote.ToApiClient()

		// Fetch the pipeline
		fqns, htt, err := api.ServicesAPI.FqnsGet(
			context.Background(),
		).Execute()

		if err != nil {
			return nil, []error{utils.ParseHTTPError(err, htt)}
		}

		return &fqns, nil
	})
}

func (m RoverCalibrationPage) fetchCalibrationReleases() tea.Cmd {
	return tui.PerformActionV2(&m.calibrationReleases, nil, func() (*[]utils.UpdateAvailable, []error) {
		releases := make([]utils.UpdateAvailable, 0)

		imaging, err := utils.CheckForGithubUpdate("imaging", "VU-ASE", "none")
		if err != nil {
			return nil, []error{err}
		}
		releases = append(releases, *imaging)

		controller, err := utils.CheckForGithubUpdate("display", "VU-ASE", "none")
		if err != nil {
			return nil, []error{err}
		}

		releases = append(releases, *controller)

		return &releases, nil
	})
}

func (m RoverCalibrationPage) installService(action *tui.ActionV2[utils.UpdateAvailable, openapi.FqnsGet200ResponseInner], service utils.UpdateAvailable) tea.Cmd {
	return tui.PerformActionV2(action, &service, func() (*openapi.FqnsGet200ResponseInner, []error) {
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

		fqn := openapi.FqnsGet200ResponseInner{
			Name:    res.Name,
			Author:  res.Author,
			Version: res.Version,
		}
		if fqn.Name != service.Name {
			return nil, []error{fmt.Errorf("Roverd failed to install service %s. It installed %s instead", fqn.Name, service.Name)}
		} else if fqn.Version != service.LatestVersion {
			return nil, []error{fmt.Errorf("Roverd failed to install service %s@%s. It installed %s@%s instead", fqn.Name, service.LatestVersion, service.Name, fqn.Version)}
		} else if !strings.EqualFold(fqn.Author, service.Author) {
			return nil, []error{fmt.Errorf("Roverd failed to install service %s@%s by %s. It installed %s@%s by %s instead", fqn.Name, service.LatestVersion, service.Author, service.Name, fqn.Version, fqn.Author)}
		}

		return &fqn, nil
	})
}
