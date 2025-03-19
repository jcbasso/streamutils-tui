package help

import (
	teaHelp "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"streamutils-tui/src/tui/colors"
)

var (
	defaultHelpModelStyle = teaHelp.Styles{
		Ellipsis:       lipgloss.NewStyle().Foreground(colors.White),
		ShortKey:       lipgloss.NewStyle().Foreground(colors.BrightWhite),
		ShortDesc:      lipgloss.NewStyle().Foreground(colors.White),
		ShortSeparator: lipgloss.NewStyle().Foreground(colors.White),
		FullKey:        lipgloss.NewStyle().Foreground(colors.BrightWhite),
		FullDesc:       lipgloss.NewStyle().Foreground(colors.White),
		FullSeparator:  lipgloss.NewStyle().Foreground(colors.White),
	}

	defaultHelpStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Height(1)
)

type Model struct {
	teaHelp.Model
	Width int
	keys  teaHelp.KeyMap
	Style lipgloss.Style
}

func New(keys teaHelp.KeyMap) Model {
	helpModel := teaHelp.New()
	helpModel.Styles = defaultHelpModelStyle
	return Model{
		Model: helpModel,
		keys:  keys,
		Style: defaultHelpStyle,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Only set the width; height is fixed.
		m.Width = msg.Width
		m.Model.Width = msg.Width
	}
	return m, nil
}

func (m Model) View() string {
	return m.Style.Render(m.Model.ShortHelpView(m.keys.ShortHelp()))
}
