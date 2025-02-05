package style

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// Styling
var (
	Docstyle          = lipgloss.NewStyle().Margin(1, 2)
	TitleStyle        = lipgloss.NewStyle().MarginLeft(0).Foreground(AsePrimary)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(AsePrimary)
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4).Foreground(AsePrimary)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(0).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

var (
	listDelegateColor = lipgloss.AdaptiveColor{
		Light: "#141414",
		Dark:  "#f2f2f2",
	}
)

// Delegates (i.e. list items)
func DefaultListDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(listDelegateColor).Bold(true).BorderForeground(listDelegateColor)
	d.Styles.SelectedDesc = d.Styles.SelectedDesc.Foreground(listDelegateColor).BorderForeground(listDelegateColor)
	return d
}
