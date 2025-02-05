package style

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// Lipgloss colors that can be easily reused
var AsePrimary = lipgloss.Color("#0077B3")
var WarningPrimary = lipgloss.Color("#FFA500")
var SuccessPrimary = lipgloss.Color("#008000")
var ErrorPrimary = lipgloss.Color("#FF0000")
var GrayPrimary = lipgloss.Color("#808080")
var SuccessLight = lipgloss.Color("#7bc990")

// Lipgloss classes that can be easily reused
var Title = lipgloss.NewStyle().Foreground(AsePrimary).Bold(true)
var Subtitle = lipgloss.NewStyle().Foreground(GrayPrimary)
var Primary = lipgloss.NewStyle().Foreground(AsePrimary)
var Error = lipgloss.NewStyle().Foreground(ErrorPrimary)
var Warning = lipgloss.NewStyle().Foreground(WarningPrimary)
var Success = lipgloss.NewStyle().Foreground(SuccessPrimary)
var Gray = lipgloss.NewStyle().Foreground(GrayPrimary)

func Color(c lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(c))
}

// Form theme
var FormTheme = huh.ThemeBase()

func RenderColor(view string, color lipgloss.Color) string {
	return lipgloss.NewStyle().Foreground(color).Render(view)
}
