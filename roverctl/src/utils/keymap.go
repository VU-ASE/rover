package utils

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

// This is a reusable key map that can be used for many different views
// (you can select which key.Binding to implement)
type GeneralKeyMap struct {
	help.KeyMap
	// For non-interactive mode
	Quit key.Binding
	// General navigation
	Back   key.Binding
	Retry  key.Binding
	Cancel key.Binding
	// Directions
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	Next     key.Binding
	Previous key.Binding
	// Operations on items
	Delete  key.Binding
	New     key.Binding
	Confirm key.Binding
	Save    key.Binding
	// Misc
	Toggle    key.Binding
	Logs      key.Binding
	Details   key.Binding
	Configure key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k GeneralKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit,
		k.Back,
		k.Retry,
		k.Cancel,
		k.Previous,
		k.Next,
		k.Up,
		k.Down,
		k.Left,
		k.Right,
		k.Delete,
		k.New,
		k.Confirm,
		k.Save,
		k.Toggle,
		k.Logs,
		k.Details,
		k.Configure,
	}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k GeneralKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

// For at least implementing the quit key
func NewGeneralKeyMap() GeneralKeyMap {
	return GeneralKeyMap{
		Back: key.NewBinding(
			key.WithKeys("ctrl+b"),
			key.WithHelp("ctrl+b", "back"),
		),
	}
}
