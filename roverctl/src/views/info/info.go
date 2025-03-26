package view_info

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// This version variable is overridden at build time using the -ldflags flag
// See the Makefile and github build-and-release workflow.
var Version = "UNSET"

type model struct {
	// If you want to query information about roverd
	rover *configuration.RoverConnection

	// Fetch information about roverd and the rover
	remoteInfo tui.Action[openapi.Get200Response]
	spinner    spinner.Model
}

func New(rover *configuration.RoverConnection) model {
	ri := tui.NewAction[openapi.Get200Response]("remoteInfo")
	sp := spinner.New()
	return model{
		rover:      rover,
		remoteInfo: ri,
		spinner:    sp,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.fetchInfo(), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tui.ActionInit[openapi.Get200Response]:
		m.remoteInfo.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[openapi.Get200Response]:
		m.remoteInfo.ProcessResult(msg)
		return m, tea.Quit
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := style.Title.Render(`             
                  ~( @\   \
             _________]_[__/_>________
            /  ____ \ <>     |  ____  \
           =\_/ __ \_\_______|_/ __ \__D
_______________(__)_____________(__)_________________`)
	s += "\n\nPowered by the " + style.Bold.Render("Vrije Universiteit Amsterdam ASE-Team") + ".\nFind more information at " + style.Primary.Render("ase.vu.nl") + "\n"

	author := state.Get().Config.Author
	if author == "" {
		author = style.Warning.Render("not set")
	}

	s += "\n" + style.Title.Render("Roverctl") + "\n"
	s += style.Gray.Render("Build version: ") + Version + "\n"
	s += style.Gray.Render("Author: ") + author + "\n"
	s += style.Gray.Render("Configuration location: ") + configuration.LocalConfigDir() + "\n"
	s += style.Gray.Render("Architecture: ") + runtime.GOOS + "/" + runtime.GOARCH + "\n"

	if m.rover != nil {
		s += "\n" + style.Title.Render("Roverd") + "\n"
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
			if !utils.VersionsEqual(Version, m.remoteInfo.Data.Version) {
				s += style.Error.Render(" (mismatch, might be incompatible)")
			} else {
				s += style.Success.Render(" (compatible)")
			}
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

func (m model) fetchInfo() tea.Cmd {
	return tui.PerformAction(&m.remoteInfo, func() (*openapi.Get200Response, error) {
		if m.rover == nil {
			return nil, fmt.Errorf("No active rover connection")
		}

		api := m.rover.ToApiClient()
		res, _, err := api.HealthAPI.StatusGet(
			context.Background(),
		).Execute()

		return res, err
	})
}
