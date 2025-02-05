package style

import (
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/charmbracelet/lipgloss"
)

func RenderDialog(dialog string, theme lipgloss.Color) string {
	s := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Padding(1, 4).Render(dialog)
	s = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(theme).
		// Background(theme).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).Render(s)

	return lipgloss.Place(state.Get().WindowWidth-4, state.Get().WindowHeight-2,
		lipgloss.Center, lipgloss.Center,
		s,
		lipgloss.WithWhitespaceChars(""),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#874BFD")),
	)
}
