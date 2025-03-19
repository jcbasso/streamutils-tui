package sizeable

import (
	teaHelp "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Model interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Model, tea.Cmd)
	View() string
	SetWidth(int)
	SetHeight(int)
	GetWidth() int
	GetHeight() int
	teaHelp.KeyMap
}
