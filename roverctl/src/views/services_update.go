package views

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/tui"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (i UpdatableItem) FilterValue() string { return i.RoverdSource.Name }

type UpdateListItemDelegate struct{}

func (d UpdateListItemDelegate) Height() int                             { return 1 }
func (d UpdateListItemDelegate) Spacing() int                            { return 0 }
func (d UpdateListItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d UpdateListItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(UpdatableItem)
	if !ok {
		return
	}

	str := i.RoverdSource.Name + style.Primary.Render(" v"+i.RoverdSource.Version) + style.Gray.Render(" -> ") + style.Primary.Render("v"+i.Release.NewVersion) + style.Gray.Render(" (from "+i.RoverdSource.Url+")")

	prefix := "[ ]"
	if i.Queued {
		prefix = "[" + style.Success.Render("✓") + "]"
	}

	str = prefix + " " + str

	fn := lipgloss.NewStyle().Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().Bold(true).Render("> " + strings.Join(s, " "))
		}
	} else {
		str = "  " + str
	}

	fmt.Fprint(w, fn(str))
}

// ServicesUpdateKeyMap defines a set of keybindings. To work for help it must satisfy key.Map
type ServicesUpdateKeyMap struct {
	Retry     key.Binding
	Confirm   key.Binding
	Quit      key.Binding
	Queue     key.Binding
	QueueAll  key.Binding
	QueueNone key.Binding
}

// replace with download manager
type OfficialRelease struct {
	NewVersion string // The new version of the source
}

type UpdatableItem struct {
	RoverdSource openapi.SourcesGet200ResponseInner
	Release      OfficialRelease
	Queued       bool // whether the user wants to update this source
}

type ServicesUpdatePage struct {
	help           help.Model
	spinner        spinner.Model
	sourceList     tui.Action[[]UpdatableItem]
	serviceUpdates map[string]tui.Action[openapi.SourcesGet200ResponseInner]
	list           list.Model
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k ServicesUpdateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Retry, k.Queue, k.QueueAll, k.QueueNone, k.Confirm, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k ServicesUpdateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

var errorFetchSourcesKeys = ServicesUpdateKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "retry"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

var successFetchSourcesKeys = ServicesUpdateKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refetch"),
	),
	Queue: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle"),
	),
	QueueAll: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "select all"),
	),
	QueueNone: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "select none"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "update"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

// Shown when the services are being updated
var updateServicesKeys = ServicesUpdateKeyMap{
	Retry: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "retry"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

func NewServicesUpdatePage() ServicesUpdatePage {
	s := spinner.New()
	s.Spinner = spinner.Line

	sourcesList := tui.NewAction[[]UpdatableItem]("sourcesList")
	servicesList := map[string]tui.Action[openapi.SourcesGet200ResponseInner]{}

	// List
	l := list.New([]list.Item{}, UpdateListItemDelegate{}, 0, 14)
	l.Title = "Select services to update"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = style.TitleStyle
	l.Styles.PaginationStyle = style.PaginationStyle
	l.Styles.HelpStyle = style.HelpStyle
	l.AdditionalShortHelpKeys = successFetchSourcesKeys.ShortHelp

	return ServicesUpdatePage{
		spinner:        s,
		help:           help.New(),
		sourceList:     sourcesList,
		serviceUpdates: servicesList,
		list:           l,
	}
}

func (m ServicesUpdatePage) Update(msg tea.Msg) (pageModel, tea.Cmd) {

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	if cmd != nil {
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
		m.list.SetSize(msg.Width, msg.Height-1)
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, successFetchSourcesKeys.Queue):
			// Queue the selected item
			if m.list.Index() >= 0 && m.list.Index() < len(m.list.Items()) {
				item := m.list.Items()[m.list.Index()].(UpdatableItem)
				item.Queued = !item.Queued
				m.list.SetItem(m.list.Index(), item)
			}
			return m, nil
		case key.Matches(msg, successFetchSourcesKeys.QueueAll):
			// Queue all items
			items := m.list.Items()
			for i, item := range items {
				updatableItem := item.(UpdatableItem)
				updatableItem.Queued = true
				m.list.SetItem(i, updatableItem)
			}
			return m, nil
		case key.Matches(msg, successFetchSourcesKeys.QueueNone):
			// Queue none
			items := m.list.Items()
			for i, item := range items {
				updatableItem := item.(UpdatableItem)
				updatableItem.Queued = false
				m.list.SetItem(i, updatableItem)
			}
			return m, nil
		case key.Matches(msg, successFetchSourcesKeys.Confirm):
			if m.sourceList.IsSuccess() && len(m.serviceUpdates) <= 0 {
				cmds := []tea.Cmd{}
				for _, i := range m.list.Items() {
					source := i.(UpdatableItem)
					if source.Queued {
						m.serviceUpdates[source.RoverdSource.Name] = tui.NewAction[openapi.SourcesGet200ResponseInner](source.RoverdSource.Name)
						cmds = append(cmds, updateService(m, source.RoverdSource.Name))
					}
				}
				return m, tea.Batch(cmds...)
			}
			return m, nil
		case key.Matches(msg, errorFetchSourcesKeys.Retry):
			// Are there any sources being updated currently?
			updateOngoing := false
			for _, action := range m.serviceUpdates {
				if action.IsLoading() {
					updateOngoing = true
					break
				}
			}

			if m.sourceList.IsDone() && !updateOngoing {
				// Clear the service updates
				m.serviceUpdates = make(map[string]tui.Action[openapi.SourcesGet200ResponseInner])
				return m, fetchSources(m)
			}

			return m, nil
		// Restore to the initial form, but recover the form values
		// m = InitialModel(m.formValues)
		// return m, tea.Batch(m.form.Init(), m.spinner.Tick)
		// case key.Matches(msg, keys.Retry):
		// 	// Retry the connection checks
		// 	// m.isChecking = true
		// 	// return m, tea.Batch(checkRoute(m), checkAuth(m), checkRoverdVersion(m), checkRoverNumber(m))

		default:
			return m, nil
		}
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tui.ActionInit[[]UpdatableItem]:
		m.sourceList.ProcessInit(msg)
		return m, nil
	case tui.ActionResult[[]UpdatableItem]:
		m.sourceList.ProcessResult(msg)
		if m.sourceList.IsSuccess() {
			// Populate the list
			items := []list.Item{}
			for _, source := range *m.sourceList.Data {
				items = append(items, UpdatableItem{
					RoverdSource: source.RoverdSource,
					Release:      source.Release,
					Queued:       true,
				})
			}
			m.list.SetItems(items)
		}
		return m, nil
	case tui.ActionInit[openapi.SourcesGet200ResponseInner]:
		newServiceUpdates := make(map[string]tui.Action[openapi.SourcesGet200ResponseInner])
		for k, v := range m.serviceUpdates {
			newServiceUpdates[k] = v
		}

		if action, ok := newServiceUpdates[msg.Name]; ok {
			action.ProcessInit(msg)
			newServiceUpdates[msg.Name] = action
		}

		m.serviceUpdates = newServiceUpdates
		return m, nil
	case tui.ActionResult[openapi.SourcesGet200ResponseInner]:
		newServiceUpdates := make(map[string]tui.Action[openapi.SourcesGet200ResponseInner])
		for k, v := range m.serviceUpdates {
			newServiceUpdates[k] = v
		}

		if action, ok := newServiceUpdates[msg.Name]; ok {
			action.ProcessResult(msg)
			newServiceUpdates[msg.Name] = action // Reassign the updated action back
		}

		m.serviceUpdates = newServiceUpdates
		return m, nil
	default:
		return m, nil
	}
}

func (m ServicesUpdatePage) fetchSourcesView() string {
	s := lipgloss.NewStyle().Foreground(style.AsePrimary).Render("Update services from source")

	if m.sourceList.IsLoading() {
		s += "\n\n " + m.spinner.View() + " Checking for updates..."
		return s
	}

	if m.sourceList.IsSuccess() {
		if len(*m.sourceList.Data) == 0 {
			s += "\n\n This rover has no enabled sources, nothing to update"
		} else {
			return m.list.View()
		}
	} else {
		s += "\n\n " + lipgloss.NewStyle().Foreground(style.ErrorPrimary).Render("Failed to fetch sources")
		s += "\n\n" + m.help.View(errorFetchSourcesKeys)
	}

	return s
}

func (m ServicesUpdatePage) testConnectionView() string {
	s := lipgloss.NewStyle().Foreground(style.AsePrimary).Render("Updating sources") + "\n"

	errorUpdates := 0
	successUpdates := 0

	// Convert the map to a slice, alphabetically sorted by key name
	keys := make([]string, 0, len(m.serviceUpdates))
	for k := range m.serviceUpdates {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		s += "\n - " + lipgloss.NewStyle().Bold(true).Render(name) + " \n  "
		action := m.serviceUpdates[name]

		if action.IsLoading() {
			s += " " + m.spinner.View() + lipgloss.NewStyle().Foreground(style.GrayPrimary).Render(" Updating...")
		} else if action.IsSuccess() {
			successUpdates++

			// Find the old version
			oldVersion := "unknown"
			if source := m.sourceList.Data; source != nil {
				for _, s := range *source {
					if s.RoverdSource.Name == name {
						oldVersion = s.RoverdSource.Version
						break
					}
				}
			}

			if oldVersion != action.Data.Version {
				s += lipgloss.NewStyle().Foreground(style.SuccessPrimary).Render(" ✓ Updated v" + oldVersion + " -> v" + action.Data.Version)
			} else {
				s += lipgloss.NewStyle().Foreground(style.AsePrimary).Render(" ✓ Unchanged at v" + action.Data.Version)
			}

		} else if action.IsError() {
			errorUpdates++
			s += lipgloss.NewStyle().Foreground(style.ErrorPrimary).Render(" ✗ Failed to update")
			if action.Error != nil {
				s += style.Gray.Render(" (" + action.Error.Error() + ")")
			}
		}
	}

	if errorUpdates+successUpdates == len(m.serviceUpdates) {
		s += "\n\n" + m.help.View(updateServicesKeys)
	}

	return s
}

func (m ServicesUpdatePage) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchSources(m))
}

func (m ServicesUpdatePage) View() string {
	if len(m.serviceUpdates) > 0 {
		return m.testConnectionView()
	}

	return m.fetchSourcesView()
}

func fetchSources(m ServicesUpdatePage) tea.Cmd {
	return tui.PerformAction(&m.sourceList, func() (*[]UpdatableItem, error) {
		// Wait 10 seconds
		time.Sleep(1 * time.Second)
		// return nil, errors.New("Failed to connect")

		// Mock sources fetching from roverd
		//! remove
		sources := []openapi.SourcesGet200ResponseInner{
			{
				Name:    ("source1"),
				Version: ("1.0.0"),
				Url:     ("https://example.com/source1"),
				Sha:     openapi.PtrString("1234567890"),
			},
			{
				Name:    ("source2a"),
				Version: ("1.0.0"),
				Url:     ("https://example.com/source2"),
				Sha:     openapi.PtrString("1234567890"),
			},
			{
				Name:    ("source3"),
				Version: ("1.0.1"),
				Url:     ("https://example.com/source1"),
				Sha:     openapi.PtrString("1234567890"),
			},
		}

		// Mock update fetching from download manager
		updates := []UpdatableItem{}
		for _, source := range sources {
			// Mock fetching the latest version from the download manager
			//! remove

			updates = append(updates, UpdatableItem{
				RoverdSource: source,
				Release: OfficialRelease{
					NewVersion: "1.0.1",
				},
			})
		}

		return &updates, nil
	})
}

func updateService(m ServicesUpdatePage, name string) tea.Cmd {
	update := m.serviceUpdates[name]
	return tui.PerformAction(&update, func() (*openapi.SourcesGet200ResponseInner, error) {
		// Wait for a random duration between 1 and 10 seconds
		time.Sleep(time.Duration(1+rand.Intn(10)) * time.Second)

		// Mock sources
		//! remove
		if len(name)%2 == 0 {
			res := openapi.SourcesGet200ResponseInner{
				Version: ("1.0.1"),
			}
			return &res, fmt.Errorf("Failed to update %s", name)
		}

		res := openapi.SourcesGet200ResponseInner{
			Version: ("1.0.1"),
		}
		return &res, nil
	})
}

func (m ServicesUpdatePage) isQuitable() bool {
	return true
}

func (m ServicesUpdatePage) keys() utils.GeneralKeyMap {
	return utils.NewGeneralKeyMap()
}

func (m ServicesUpdatePage) previousPage() *pageModel {
	return nil
}
