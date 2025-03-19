package debug

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"streamutils-tui/src/tui/colors"
	"strings"
)

var (
	defaultDebugStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colors.White).
		Padding(1, 2).
		Align(lipgloss.Center)
)

type Model struct {
	Active   bool
	Messages []string
	Style    lipgloss.Style
}

var instance = (*Model)(nil)

func GetInstance() *Model {
	if instance == nil {
		new()
	}
	return instance
}

func new() *Model {
	instance = &Model{
		Active:   false,
		Messages: []string{},
		Style:    defaultDebugStyle,
	}
	return instance
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) LogDebug(s string) {
	m.Messages = append(m.Messages, s)
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	if !m.Active {
		return ""
	}
	return m.Style.Render(strings.Join(m.Messages, "\n"))
}
