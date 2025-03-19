package sizeable

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var _ Model = (*Wrapper)(nil)

type Wrapper struct {
	model     tea.Model
	setWidth  func(int)
	setHeight func(int)
	getWidth  func() int
	getHeight func() int
	shortHelp func() []key.Binding
	fullHelp  func() [][]key.Binding
}

func NewWrapper(
	model tea.Model,
	setWidth func(int),
	setHeight func(int),
	getWidth func() int,
	getHeight func() int,
	shortHelp func() []key.Binding,
	fullHelp func() [][]key.Binding,
) *Wrapper {
	return &Wrapper{
		model:     model,
		setWidth:  setWidth,
		setHeight: setHeight,
		getWidth:  getWidth,
		getHeight: getHeight,
		shortHelp: shortHelp,
		fullHelp:  fullHelp,
	}
}

// tea.Model Interface

func (s *Wrapper) Init() tea.Cmd {
	return s.model.Init()
}

func (s *Wrapper) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	s.model, cmd = s.model.Update(msg)
	return s, cmd
}

func (s *Wrapper) View() string {
	return s.model.View()
}

// sizeable.Model specifics interface

func (s *Wrapper) SetWidth(width int) {
	s.setWidth(width)
}

func (s *Wrapper) SetHeight(height int) {
	s.setHeight(height)
}

func (s *Wrapper) GetWidth() int {
	return s.getWidth()
}

func (s *Wrapper) GetHeight() int {
	return s.getHeight()
}

// help.KeyMap Interface

func (s *Wrapper) ShortHelp() []key.Binding {
	return s.shortHelp()
}

func (s *Wrapper) FullHelp() [][]key.Binding {
	return s.fullHelp()
}
