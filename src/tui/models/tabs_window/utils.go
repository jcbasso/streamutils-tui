package tabs_window

import (
	"github.com/charmbracelet/lipgloss"
	"streamutils-tui/src/tui/colors"
	"strings"
)

func buildTopRow(
	row string,
	width int,
	style lipgloss.Style,
	tabsBorderSpace int,
) string {
	borderStyle := style.GetBorderStyle()
	topBorderColor := lipgloss.NewStyle().Foreground(style.GetBorderTopForeground())

	leftBorder := strings.Repeat(borderStyle.Top, tabsBorderSpace)
	rightBorder := strings.Repeat(borderStyle.Top, width-lipgloss.Width(leftBorder)-lipgloss.Width(row))

	return lipgloss.JoinHorizontal(lipgloss.Top, topBorderColor.Render(leftBorder), row, topBorderColor.Render(rightBorder))
}

// TODO: Fix it to make this whole model to work with margins
// Using this function, so we don't have to change the style and the height of the viewport
func removeFirstRow(s string) string {
	index := strings.IndexByte(s, '\n')
	if index == -1 {
		return ""
	}
	return s[index+1:]
}

func DefaultActiveTabStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(colors.BrightWhite).
		Padding(0, 1)
}

func DefaultInactiveTabStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(colors.White).
		Padding(0, 1)
}

func DefaultTabsWindowStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colors.BrightGreen)
}
