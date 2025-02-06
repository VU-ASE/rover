package views

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	probing "github.com/prometheus-community/pro-bing"
)

type ConnectionInitFormValues struct {
	index      string
	username   string
	password   string
	customHost string
}

type ConnectionsInitPage struct {
	form          *huh.Form
	spinner       spinner.Model
	routeExists   tui.Action[bool]
	authValid     tui.Action[bool]
	roverdVersion tui.Action[string]
	roverStatus   tui.Action[openapi.StatusGet200Response]
	isChecking    bool
	formValues    *ConnectionInitFormValues
	host          string // the ip or hostname of the rover to connect to
	error         error  // any errors that occurred
}

func createForm(formValues *ConnectionInitFormValues, advanced bool) *huh.Form {
	if advanced {
		return huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Enter the Rover index (1-20, inclusive)").
					CharLimit(3).
					Prompt("> ").
					Value(&formValues.index).
					Validate(func(s string) error {
						index, err := strconv.Atoi(s)
						if err != nil || index < 1 || index > 20 {
							return fmt.Errorf("Please enter a valid Rover index between 1 and 20 (inclusive)")
						}
						return nil
					}),
				huh.NewInput().
					Title("Enter the Roverd username").
					CharLimit(255).
					Prompt("> ").
					Value((&formValues.username)),
				huh.NewInput().
					Title("Enter the Roverd password").
					CharLimit(255).
					Prompt("> ").
					EchoMode(huh.EchoModePassword).
					Value((&formValues.password)),
				huh.NewInput().
					Title("(optional) Enter a custom hostname or IP address to connect to").
					CharLimit(255).
					Prompt("> ").
					Value(&formValues.customHost),
			),
		).WithTheme(style.FormTheme).WithShowHelp(false)
	} else {
		return huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Enter the Rover index (1-20, inclusive)").
					CharLimit(3).
					Prompt("> ").
					Value(&formValues.index).
					Validate(func(s string) error {
						index, err := strconv.Atoi(s)
						if err != nil || index < 1 || index > 20 {
							return fmt.Errorf("Please enter a valid Rover index between 1 and 20 (inclusive)")
						}
						return nil
					}),
			),
		).WithTheme(style.FormTheme).WithShowHelp(false)
	}
}

func NewConnectionsInitPage(val *ConnectionInitFormValues) ConnectionsInitPage {
	s := spinner.New()
	s.Spinner = spinner.Line

	formValues := &ConnectionInitFormValues{
		index:      "",
		username:   "debix",
		password:   "debix",
		customHost: "",
	}
	if val != nil {
		formValues = val
	}

	routeExistsAction := tui.NewAction[bool]("routeExists")
	authValidAction := tui.NewAction[bool]("authValid")
	roverdVersionAction := tui.NewAction[string]("roverdVersion")
	roverStatusAction := tui.NewAction[openapi.StatusGet200Response]("roverStatus")

	return ConnectionsInitPage{
		spinner:       s,
		formValues:    formValues,
		host:          "",
		routeExists:   routeExistsAction,
		authValid:     authValidAction,
		roverdVersion: roverdVersionAction,
		roverStatus:   roverStatusAction,
		isChecking:    false,
		error:         nil,
		form:          createForm(formValues, false),
	}
}

func (m ConnectionsInitPage) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+a":
			// Enter advanced mode
			m.form = createForm(m.formValues, true)
			return m, m.form.Init()
		}
	}

	if m.form.State == huh.StateCompleted {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.keys().Back):
				// Restore to the initial form, but recover the form values
				m = NewConnectionsInitPage(m.formValues)
				return m, tea.Batch(m.form.Init(), m.spinner.Tick)
			case key.Matches(msg, m.keys().Retry):
				// Retry the connection checks
				m.isChecking = true
				return m, tea.Batch(m.checkRoute(), m.checkAuth(), m.checkRoverdVersion(), m.checkRoverStatus())
			case key.Matches(msg, m.keys().Save) && !m.allChecksSuccessful():
				// Force save the connection
				m.saveConnection()
				return RootScreen(state.Get()).SwitchScreen(NewStartPage())
			case key.Matches(msg, m.keys().Confirm) && m.allChecksSuccessful():
				// Continue to main screen again (already saved)
				return RootScreen(state.Get()).SwitchScreen(NewStartPage())
			}
		}
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tui.ActionInit[bool]:
		m.routeExists.ProcessInit(msg)
		m.authValid.ProcessInit(msg)
		return m, nil
	case tui.ActionInit[string]:
		m.roverdVersion.ProcessInit(msg)
		return m, nil
	case tui.ActionInit[openapi.StatusGet200Response]:
		m.roverStatus.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[bool]:
		m.authValid.ProcessResult(msg)
		m.routeExists.ProcessResult(msg)
		if m.allChecksSuccessful() {
			m.saveConnection()
		}
		return m, nil
	case tui.ActionResult[string]:
		m.roverdVersion.ProcessResult(msg)
		if m.allChecksSuccessful() {
			m.saveConnection()
		}
		return m, nil
	case tui.ActionResult[openapi.StatusGet200Response]:
		m.roverStatus.ProcessResult(msg)
		if m.allChecksSuccessful() {
			m.saveConnection()
		}
		return m, nil
	default:
		cmds := []tea.Cmd{}
		form, cmd := m.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.form = f
			if f.State == huh.StateCompleted && !m.isChecking {
				m.isChecking = true

				index, err := strconv.Atoi(m.formValues.index)
				if err != nil || index < 1 || index > 20 {
					m.routeExists = tui.NewAction[bool]("routeExists")
					return m, cmd
				}
				m.host = fmt.Sprintf("192.168.0.%d", index+100)
				if len(m.formValues.customHost) > 0 {
					m.host = m.formValues.customHost
				}

				// We are optimistic, start all checks in parallel
				cmds = append(cmds, m.checkRoute(), m.checkAuth(), m.checkRoverdVersion(), m.checkRoverStatus())
			}
		}
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}
}

// the update view with the view method
func (m ConnectionsInitPage) enterDetailsView() string {
	// Introduction
	s := lipgloss.NewStyle().Foreground(style.AsePrimary).Render("Connect to a Rover")

	s += "\n\n" + m.form.View()

	return s
}

func (m ConnectionsInitPage) testConnectionView() string {
	s := lipgloss.NewStyle().Foreground(style.AsePrimary).Render("Connecting to Rover")

	if m.routeExists.IsLoading() || m.authValid.IsLoading() || m.roverdVersion.IsLoading() || m.roverStatus.IsLoading() {
		s += "\n\n " + m.spinner.View() + " Performing connection checks..."

		return s
	}

	if !m.routeExists.IsSuccess() {
		s += "\n\n ✗ " + lipgloss.NewStyle().Foreground(style.ErrorPrimary).Render("No route could be established to the Rover. Are you sure it is powered on? (Tried "+m.host+")\n   Read how to connect at: https://docs.ase.vu.nl/docs/tutorials/setting-up-your-workspace/accessing-the-network")
		if m.routeExists.Error != nil {
			s += "\n   (" + m.routeExists.Error.Error() + ")"
		}
	} else {
		s += "\n\n ✓ " + lipgloss.NewStyle().Foreground(style.SuccessPrimary).Render("Established route to "+m.host)
	}

	if m.routeExists.IsSuccess() {
		index, _ := strconv.Atoi(m.formValues.index)
		if !m.roverStatus.IsSuccess() {
			s += "\n ✗ " + lipgloss.NewStyle().Foreground(style.ErrorPrimary).Render("Could not determine rover status")
			if m.roverStatus.Error != nil {
				s += " (" + m.roverStatus.Error.Error() + ")"
			}
		} else if m.roverStatus.Data.GetRoverId() != int32(index) {
			s += "\n ! " + lipgloss.NewStyle().Foreground(style.WarningPrimary).Render("This Rover presented itself as Rover "+strconv.Itoa(int(m.roverStatus.Data.GetRoverId()))+" ("+*m.roverStatus.Data.RoverName+") but you wanted to connect to Rover "+m.formValues.index)
		} else {
			s += "\n ✓ " + lipgloss.NewStyle().Foreground(style.SuccessPrimary).Render("Discovered Rover "+m.formValues.index+" ("+*m.roverStatus.Data.RoverName+")")
		}

		if !m.roverdVersion.IsSuccess() {
			s += "\n ✗ " + lipgloss.NewStyle().Foreground(style.ErrorPrimary).Render("Could not determine roverd version")
			if m.roverdVersion.Error != nil {
				s += " (" + m.roverdVersion.Error.Error() + ")"
			}
		} else if *m.roverdVersion.Data == version {
			s += "\n ✓ " + lipgloss.NewStyle().Foreground(style.SuccessPrimary).Render("Roverd is running at version "+*m.roverdVersion.Data)
		} else {
			s += "\n ! " + style.Warning.Render("Roverd is running at version "+*m.roverdVersion.Data+" but you are using roverctl version "+version+".")
			s += style.Gray.Render(" This might cause issues in the future")
		}

		if !m.authValid.IsSuccess() {
			s += "\n ✗ " + lipgloss.NewStyle().Foreground(style.ErrorPrimary).Render("Authentication to the roverd endpoint failed. Please check your credentials")
			if m.authValid.Error != nil {
				s += " (" + m.authValid.Error.Error() + ")"
			}
		} else {
			s += "\n ✓ " + lipgloss.NewStyle().Foreground(style.SuccessPrimary).Render("Authentication successful")
		}
	}

	if !m.routeExists.IsSuccess() || !m.authValid.IsSuccess() || !m.roverdVersion.IsSuccess() {
		s += "\n\n" + lipgloss.NewStyle().Foreground(style.GrayPrimary).Render("This connection configuration is not valid and should not be saved.")
	} else {
		s += "\n\n" + "You are all set! This connection is saved."
	}

	return s
}

func (m ConnectionsInitPage) Init() tea.Cmd {
	return tea.Batch(m.form.Init(), m.spinner.Tick)
}

func (m ConnectionsInitPage) View() string {
	s := ""
	if m.form.State == huh.StateCompleted {
		s = m.testConnectionView()
	} else {
		s = m.enterDetailsView()
	}

	return s
}

// Save to state (which will later be saved to disk)
func (m ConnectionsInitPage) saveConnection() {
	name := "unnamed"
	if m.roverStatus.IsSuccess() {
		name = *m.roverStatus.Data.RoverName
	}

	state.Get().RoverConnections = state.Get().RoverConnections.Add(configuration.RoverConnection{
		Name:     "Rover " + m.formValues.index + " (" + name + ")",
		Host:     m.host,
		Username: m.formValues.username,
		Password: m.formValues.password,
	})
}

func (m ConnectionsInitPage) allChecksSuccessful() bool {
	return m.routeExists.IsSuccess() && m.authValid.IsSuccess() && m.roverdVersion.IsSuccess() && m.roverStatus.IsSuccess()
}

func (m ConnectionsInitPage) checkRoute() tea.Cmd {
	return tui.PerformAction(&m.routeExists, func() (*bool, error) {
		// Separate the host from the port
		parts := strings.Split(m.host, ":")
		if len(parts) <= 0 {
			return nil, fmt.Errorf("Invalid host format")
		}

		host := parts[0]
		ping, _ := probing.NewPinger(host)
		ping.Count = 3
		ping.Timeout = 10 * time.Second
		err := ping.Run()

		valid := ping.Statistics().PacketsRecv > 0
		if !valid {
			err = fmt.Errorf("No route to host")
		}
		return &valid, err
	})
}

func (m ConnectionsInitPage) checkAuth() tea.Cmd {
	return tui.PerformAction(&m.authValid, func() (*bool, error) {
		// Send a protected request to the roverd endpoint
		c := configuration.RoverConnection{
			Host:     m.host,
			Username: m.formValues.username,
			Password: m.formValues.password,
		}
		a := c.ToApiClient()

		_, _, err := a.PipelineAPI.PipelineGet(context.Background()).Execute()
		res := err == nil
		return &res, err
	})
}

func (m ConnectionsInitPage) checkRoverdVersion() tea.Cmd {
	return tui.PerformAction(&m.roverdVersion, func() (*string, error) {
		c := configuration.RoverConnection{
			Host:     m.host,
			Username: m.formValues.username,
			Password: m.formValues.password,
		}
		a := c.ToApiClient()

		res, _, err := a.HealthAPI.StatusGet(context.Background()).Execute()
		if err != nil {
			return nil, err
		}

		version := res.Version
		return &version, nil
	})
}

func (m ConnectionsInitPage) checkRoverStatus() tea.Cmd {
	return tui.PerformAction(&m.roverStatus, func() (*openapi.StatusGet200Response, error) {
		c := configuration.RoverConnection{
			Host:     m.host,
			Username: m.formValues.username,
			Password: m.formValues.password,
		}
		a := c.ToApiClient()

		res, _, err := a.HealthAPI.StatusGet(context.Background()).Execute()
		if err != nil {
			return nil, err
		}

		return res, err
	})
}

func (m ConnectionsInitPage) isQuitable() bool {
	return m.form.State != huh.StateCompleted || ((!(m.routeExists.IsError() || m.authValid.IsError() || m.roverdVersion.IsError() || m.roverStatus.IsError())) && (m.routeExists.IsDone() && m.authValid.IsDone() && m.roverdVersion.IsDone() && m.roverStatus.IsDone()))
}

func (m ConnectionsInitPage) keys() utils.GeneralKeyMap {
	kb := utils.NewGeneralKeyMap()
	if m.form.State != huh.StateCompleted {
		kb.Confirm = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "connect"),
		)
	} else if m.allChecksSuccessful() {
		kb.Confirm = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "close"),
		)
	}
	if (m.routeExists.IsError() || m.authValid.IsError() || m.roverdVersion.IsError() || m.roverStatus.IsError()) && (m.routeExists.IsDone() && m.authValid.IsDone() && m.roverdVersion.IsDone() && m.roverStatus.IsDone()) {
		kb.Retry = key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "retry"),
		)
		kb.Save = key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "save anyway"),
		)
	}

	return kb
}

func (m ConnectionsInitPage) previousPage() *pageModel {
	var pageModel pageModel = NewStartPage()
	return &pageModel
}
