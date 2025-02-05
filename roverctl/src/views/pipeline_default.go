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

type PipelineDownloadDefaultPage struct {
	spinner spinner.Model

	// action
	installed tui.Action[[]openapi.FetchPost200Response]
}

func NewPipelineDownloadDefaultPage() PipelineDownloadDefaultPage {
	i := tui.NewAction[[]openapi.FetchPost200Response]("installed")
	return PipelineDownloadDefaultPage{
		spinner:   spinner.New(),
		installed: i,
	}
}

//
// Page model methods
//

func (m PipelineDownloadDefaultPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
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
	case tui.ActionInit[[]openapi.FetchPost200Response]:
		m.installed.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[[]openapi.FetchPost200Response]:
		m.installed.ProcessResult(msg)
		return m, nil
	}

	return m, nil
}

func (m PipelineDownloadDefaultPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.installBasicPipeline())
}

func (m PipelineDownloadDefaultPage) View() string {
	s := style.Title.Render("Download default pipeline page") + "\n\n"

	if m.installed.IsError() {
		s += style.Error.Render("✗ Could not install pipeline") + style.Gray.Render(" ("+m.installed.Error.Error()+")") + "\n\n"
	} else if m.installed.IsSuccess() {
		s += style.Success.Render("✓ Pipeline installed") + "\n\n"
		s += "  The following services were installed:\n"
		s += "    - " + style.Gray.Render("imaging") + ": https://github.com/VU-ASE/imaging\n"
		s += "    - " + style.Gray.Render("controller") + ": https://github.com/VU-ASE/controller\n"
		s += "    - " + style.Gray.Render("actuator") + ": https://github.com/VU-ASE/actuator\n\n"
		s += "  Take a look at their respective repositories for more information.\n\n"

	} else {
		s += m.spinner.View() + " Installing latest official ASE pipeline..." + "\n\n"
	}

	return s
}

func (m PipelineDownloadDefaultPage) isQuitable() bool {
	return true
}

func (m PipelineDownloadDefaultPage) keys() utils.GeneralKeyMap {
	kb := utils.NewGeneralKeyMap()
	if m.installed.IsLoading() {
		kb.Back.SetEnabled(false)
	}
	return kb
}

func (m PipelineDownloadDefaultPage) previousPage() *pageModel {
	var pageModel pageModel = NewStartPage()
	return &pageModel
}

//
// Actions
//

func (m PipelineDownloadDefaultPage) installBasicPipeline() tea.Cmd {
	return tui.PerformAction(&m.installed, func() (*[]openapi.FetchPost200Response, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		// First, save the pipeline
		api := remote.ToApiClient()
		req := api.ServicesAPI.FetchPost(
			context.Background(),
		)

		services := []string{
			"imaging", "controller", "actuator",
		}

		ress := make([]openapi.FetchPost200Response, 0)
		for _, service := range services {

			baseUrl := "https://github.com/VU-ASE/" + service + "/releases/latest"

			// Visit the URL and follow the redirect
			releaseUrl, err := utils.FollowRedirects(baseUrl)
			if err != nil {
				return nil, err
			}

			// Download url is in the form .../<service>.zip
			url := releaseUrl + "/" + service + ".zip"
			url = strings.Replace(url, "tag", "download", 1)
			pipelineReq := openapi.FetchPostRequest{
				Url: url,
			}
			req = req.FetchPostRequest(pipelineReq)
			res, htt, err := req.Execute()
			if err != nil {
				return nil, utils.ParseHTTPError(err, htt)
			}
			ress = append(ress, *res)
		}
		return &ress, nil
	})
}
