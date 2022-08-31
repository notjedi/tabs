package tabs

import (
	"github.com/charmbracelet/lipgloss"
)

var(
	docStyle = lipgloss.NewStyle().Margin(1, 2, 1, 2)

    inactiveTab = lipgloss.NewStyle().
		Faint(true).
		Padding(0, 2)
	activeTab = inactiveTab.
			Copy().
			Faint(false).
			Bold(true).
			Background(lipgloss.AdaptiveColor{Light: "006", Dark: "008"}).
			Foreground(lipgloss.AdaptiveColor{Light: "000", Dark: "015"})
)
