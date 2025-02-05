package views

import (
	"context"
	"fmt"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//
// All keys
//

// Keys to navigate
type PipelineDetailsKeyMap struct {
	Retry   key.Binding
	Confirm key.Binding
	Quit    key.Binding
}

// Shown when the services are being updated
var pipelineDetailsKeysRegular = PipelineDetailsKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

func (k PipelineDetailsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Confirm, k.Quit}
}

func (k PipelineDetailsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

//
// The page model
//

type PipelineDetailsPage struct {
	help    help.Model
	spinner spinner.Model
	service PipelineOverviewServiceInfo
	details tui.Action[openapi.ServicesAuthorServiceVersionGet200Response]
}

func NewPipelineDetailsPage(s PipelineOverviewServiceInfo) PipelineDetailsPage {
	return PipelineDetailsPage{
		spinner: spinner.New(),
		help:    help.New(),
		service: s,
		details: tui.NewAction[openapi.ServicesAuthorServiceVersionGet200Response]("details"),
	}
}

//
// Page model methods
//

func (m PipelineDetailsPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tui.ActionInit[openapi.ServicesAuthorServiceVersionGet200Response]:
		m.details.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[openapi.ServicesAuthorServiceVersionGet200Response]:
		m.details.ProcessResult(msg)
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, pipelineDetailsKeysRegular.Quit):
			return m, tea.Quit
		case key.Matches(msg, pipelineDetailsKeysRegular.Retry):
			return m, m.fetchVersionDetails()
		case key.Matches(msg, pipelineDetailsKeysRegular.Confirm):
			// todo:
			return m, nil
		}
	}

	return m, nil
}

func (m PipelineDetailsPage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchVersionDetails())
}

func (m PipelineDetailsPage) View() string {
	s := style.Title.Render("Details") + "\n\n"

	if m.details.IsError() {
		s += style.Error.Render("Error fetching service details: ") + m.details.Error.Error()
	} else if m.details.HasData() {
		service := *m.details.Data

		// Calculate column width (subtract padding and borders)
		columnWidth := (state.Get().WindowWidth - 4 - 6) / 3 // Adjust for padding and borders

		// Define styles for each column
		columnStyle := lipgloss.NewStyle().
			Width(columnWidth)

		// About
		about := style.Gray.Render("Service: ") + m.service.Name + "\n"
		about += style.Gray.Render("Author: ") + m.service.Author + "\n"
		about += style.Gray.Render("Version: ") + m.service.Version + "\n"
		about += style.Gray.Render("Last build: ")
		if service.BuiltAt != nil {
			about += time.Unix(*service.BuiltAt/1000, 0).Format("2006-01-02 15:04:05")
		} else {
			about += "not built yet"
		}
		about += "\n"

		inputs := ""
		if len(service.Inputs) <= 0 {
			inputs += style.Gray.Render("This service does not depend on streams from other services")
		} else {
			inputs += style.Gray.Render("Inputs: ") + "\n"

			for _, i := range m.service.Configuration.Inputs {
				for _, s := range i.Streams {
					inputs += style.Gray.Render("- ") + s + style.Gray.Render(" from ") + i.Service + "\n"
				}
			}
		}

		outputs := ""
		if len(m.service.Configuration.Outputs) <= 0 {
			outputs += style.Gray.Render("This service does not provide streams to other services")
		} else {
			outputs += style.Gray.Render("Outputs: ") + "\n"

			for _, o := range m.service.Configuration.Outputs {
				outputs += style.Gray.Render("- ") + o + "\n"
			}
		}

		row := lipgloss.JoinHorizontal(lipgloss.Top,
			columnStyle.Render(about),
			columnStyle.Render(inputs),
			columnStyle.Render(outputs),
		)

		s += row
	} else {
		s += m.spinner.View() + " Fetching service details..."
	}

	return s + "\n" + m.help.View(pipelineDetailsKeysRegular)
}

func (m PipelineDetailsPage) fetchVersionDetails() tea.Cmd {
	return tui.PerformAction(&m.details, func() (*openapi.ServicesAuthorServiceVersionGet200Response, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, htt, err := api.ServicesAPI.ServicesAuthorServiceVersionGet(
			context.Background(),
			m.service.Author,
			m.service.Name,
			m.service.Version,
		).Execute()

		if err != nil && htt != nil {
			return nil, utils.ParseHTTPError(err, htt)
		}

		return res, err
	})
}

func (m PipelineDetailsPage) isQuitable() bool {
	return true
}

func (m PipelineDetailsPage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m PipelineDetailsPage) previousPage() *pageModel {
	return nil
}
