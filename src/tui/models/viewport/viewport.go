package viewport

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = (*Model)(nil)
var _ help.KeyMap = (*Model)(nil)

type Model struct {
	viewport.Model
}

func New(width int, height int, content string) Model {
	vp := viewport.New(width, height)
	vp.SetContent(content)

	return Model{
		Model: vp,
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m Model) UpdateSpecific(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	_, cmd = m.Update(msg)
	return m, cmd
}

func (m Model) ShortHelp() []key.Binding {
	return []key.Binding{m.KeyMap.Up, m.KeyMap.Down}
}

func (m Model) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{m.KeyMap.Up, m.KeyMap.Down, m.KeyMap.PageUp, m.KeyMap.PageDown, m.KeyMap.HalfPageUp, m.KeyMap.HalfPageDown},
	}
}
