package chat

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"streamutils-tui/src/entities"
	"streamutils-tui/src/tui/colors"
	"streamutils-tui/src/tui/models/viewport"

	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = (*Model)(nil)
var _ help.KeyMap = (*Model)(nil)

type Model struct {
	viewport.Model
	KeyMap
	content string
}

func New() Model {
	vp := viewport.New(0, 0, "")
	return Model{
		Model:  vp,
		KeyMap: DefaultKeyMap(),
	}
}

func (m *Model) SetKeyMap(keyMap KeyMap) {
	m.KeyMap = keyMap
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case entities.Response:
		switch {
		case msg.Payload.Action == "SUCCESS":
			if m.content != "" {
				m.content += "\n"
			}
			m.content += colorizeString("Welcome to the chat room!", colors.White)
			m.SetContent(m.content)
			m.GotoBottom()
		case msg.Payload.Action == "PRIVMSG":
			if m.content != "" {
				m.content += "\n"
			}

			name := msg.Payload.Username
			if msg.Tags.DisplayName != "" {
				name = msg.Tags.DisplayName
			}
			var nameColor lipgloss.TerminalColor = lipgloss.Color(msg.Tags.Color)
			if msg.Tags.Color == "" {
				// Setting default color to twitch magenta
				nameColor = colors.Magenta
			}

			colorizedName := colorizeString(name, nameColor)
			colorizedTime := colorizeString(msg.Tags.Time.Format("15:04"), colors.White)
			m.content += fmt.Sprintf("%s %s: %s", colorizedTime, colorizedName, msg.Payload.Message)
			m.SetContent(m.content)
			m.GotoBottom()
		}
	}

	m.Model, cmd = m.Model.UpdateSpecific(msg) // Only update the viewport model
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func colorizeString(text string, color lipgloss.TerminalColor) string {
	style := lipgloss.NewStyle().Foreground(color)
	return style.Render(text)
}

func (m Model) View() string {
	return m.Model.View()
}

func (m Model) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (m Model) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
