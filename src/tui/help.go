package tui

import (
	teaHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type helpKeyMaps struct {
	keyMaps []teaHelp.KeyMap
}

func (hm helpKeyMaps) ShortHelp() []key.Binding {
	var bindings []key.Binding
	for _, elem := range hm.keyMaps {
		bindings = append(bindings, elem.ShortHelp()...)
	}
	return bindings
}

func (hm helpKeyMaps) FullHelp() [][]key.Binding {
	var bindings [][]key.Binding
	for _, elem := range hm.keyMaps {
		bindings = append(bindings, elem.FullHelp()...)
	}
	return bindings
}

func (hm helpKeyMaps) OtherFunc() [][]key.Binding {
	var bindings [][]key.Binding
	for _, elem := range hm.keyMaps {
		bindings = append(bindings, elem.FullHelp()...)
	}
	return bindings
}
