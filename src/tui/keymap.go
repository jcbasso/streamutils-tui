package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Quit  key.Binding
	Debug key.Binding
	Enter key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit"),
		),
		Debug: key.NewBinding(
			key.WithKeys("D"),
			key.WithHelp("D", "toggle debug"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "send message"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Debug}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Debug},
	}
}
