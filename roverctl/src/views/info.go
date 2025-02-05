package views

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/VU-ASE/roverctl/src/configuration"
	"github.com/VU-ASE/roverctl/src/openapi"
	"github.com/VU-ASE/roverctl/src/state"
	"github.com/VU-ASE/roverctl/src/style"
	"github.com/VU-ASE/roverctl/src/tui"
	"github.com/VU-ASE/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

var version = "UNSET"

type InfoPage struct {
	// Fetch information about roverd and the rover
	remoteInfo tui.Action[openapi.StatusGet200Response]
	spinner    spinner.Model
}

func NewInfoPage() InfoPage {
	ri := tui.NewAction[openapi.StatusGet200Response]("remoteInfo")
	sp := spinner.New()
	return InfoPage{
		remoteInfo: ri,
		spinner:    sp,
	}
}

func (m InfoPage) Init() tea.Cmd {
	return tea.Batch(m.fetchInfo(), m.spinner.Tick)
}

func (m InfoPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tui.ActionInit[openapi.StatusGet200Response]:
		m.remoteInfo.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[openapi.StatusGet200Response]:
		m.remoteInfo.ProcessResult(msg)
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m InfoPage) View() string {
	s := style.Title.Render(`░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░       ░▒▓██████▓▒░ ░▒▓███████▓▒░▒▓████████▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        
 ░▒▓█▓▒▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        
 ░▒▓█▓▒▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░      ░▒▓████████▓▒░░▒▓██████▓▒░░▒▓██████▓▒░   
  ░▒▓█▓▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
  ░▒▓█▓▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
   ░▒▓██▓▒░   ░▒▓██████▓▒░       ░▒▓█▓▒░░▒▓█▓▒░▒▓███████▓▒░░▒▓████████▓▒░ `)
	s += "\n\nBrought to you by the Vrije Universiteit Amsterdam ASE-Team.\nCheck out ase.vu.nl for more information\n"

	s += "\n" + style.Title.Render("Roverctl") + "\n"
	s += style.Gray.Render("Build version: ") + version + "\n"
	s += style.Gray.Render("Author: ") + state.Get().Config.Author + "\n"
	s += style.Gray.Render("Configuration location: ") + configuration.LocalConfigDir() + "\n"
	s += style.Gray.Render("Architecture: ") + runtime.GOOS + "/" + runtime.GOARCH + "\n"

	s += "\n" + style.Title.Render("Roverd") + "\n"
	if state.Get().RoverConnections.GetActive() != nil {
		if m.remoteInfo.IsSuccess() {
			if m.remoteInfo.Data.RoverId != nil {
				str := fmt.Sprintf("%d", *m.remoteInfo.Data.RoverId)
				s += style.Gray.Render("Rover: ") + str
				if m.remoteInfo.Data.RoverName != nil {
					s += " (" + *m.remoteInfo.Data.RoverName + ")"
				}
				s += "\n"
			}
			s += style.Gray.Render("Build version: ") + m.remoteInfo.Data.Version
			// if m.remoteInfo.Data.Version != version {
			// 	s += style.Error.Render(" (mismatch)")
			// } else {
			// 	s += style.Success.Render(" (match)")
			// }
			s += "\n"
			s += style.Gray.Render("Status: ")
			if m.remoteInfo.Data.Status == openapi.AllowedDaemonStatusEnumValues[0] {
				s += "operational"
			} else if m.remoteInfo.Data.Status == openapi.AllowedDaemonStatusEnumValues[1] {
				s += style.Warning.Render("recoverable")
			} else if m.remoteInfo.Data.Status == openapi.AllowedDaemonStatusEnumValues[2] {
				s += style.Error.Render("Unrecoverable")
			} else {
				s += style.Error.Render("Unknown")
			}
			if m.remoteInfo.Data.ErrorMessage != nil {
				s += style.Gray.Render(" (" + *m.remoteInfo.Data.ErrorMessage + ")")
			}
			s += "\n"
			s += style.Gray.Render("OS: ") + m.remoteInfo.Data.Os + "\n"
			s += style.Gray.Render("System time: ") + time.Unix(m.remoteInfo.Data.Systime/1000, 0).String() + "\n"
			upt := fmt.Sprintf("%ds", m.remoteInfo.Data.Uptime/1000)
			if m.remoteInfo.Data.Uptime > 60*1000 {
				upt = fmt.Sprintf("%dm", m.remoteInfo.Data.Uptime/(60*1000))
			}

			s += style.Gray.Render("Uptime: ") + upt + style.Gray.Render(" since ") + time.Unix((m.remoteInfo.Data.Systime-int64(m.remoteInfo.Data.Uptime))/1000, 0).String() + "\n"

		} else if m.remoteInfo.IsError() {
			s += style.Error.Render("Failed to fetch information") + style.Gray.Render(" ("+m.remoteInfo.Error.Error()+")")
		} else {
			s += m.spinner.View() + " Fetching information..."
		}
	} else {
		s += style.Gray.Render("No active rover connection configured")
	}

	return s
}

func (m InfoPage) fetchInfo() tea.Cmd {
	return tui.PerformAction(&m.remoteInfo, func() (*openapi.StatusGet200Response, error) {
		remote := state.Get().RoverConnections.GetActive()
		if remote == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := remote.ToApiClient()
		res, _, err := api.HealthAPI.StatusGet(
			context.Background(),
		).Execute()

		return res, err
	})
}

func (m InfoPage) isQuitable() bool {
	return true
}

func (m InfoPage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m InfoPage) previousPage() *pageModel {
	var pageModel pageModel = NewStartPage()
	return &pageModel
}
