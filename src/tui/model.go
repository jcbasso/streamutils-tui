package tui

import (
	teaHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"streamutils-tui/src/entities"
	"streamutils-tui/src/tui/models/debug"
	"streamutils-tui/src/tui/models/help"
	"streamutils-tui/src/tui/models/overlay"
	"streamutils-tui/src/tui/models/tabs_window"
)

type Model struct {
	messages   chan entities.Response
	tabsWindow tabs_window.Model
	help       help.Model
	debug      *debug.Model
	KeyMap     KeyMap
	ready      bool
	width      int
	height     int
}

func New(messages chan entities.Response) Model {
	tabs := []tabs_window.Tab{
		{Title: "Chat", Model: generateSizeableChat()},
	}

	tabsWindow := tabs_window.New(tabs)

	keyMap := DefaultKeyMap()
	keyMaps := []teaHelp.KeyMap{keyMap, tabsWindow}

	return Model{
		tabsWindow: tabsWindow,
		help:       help.New(helpKeyMaps{keyMaps: keyMaps}),
		debug:      debug.GetInstance(),
		KeyMap:     keyMap,
		messages:   messages,
	}
}

func (m Model) Run() {
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m Model) Init() tea.Cmd {
	return listenToMessages(m.messages)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.KeyMap.Debug):
			m.debug.Active = !m.debug.Active
		}
	case tea.WindowSizeMsg:
		if !m.ready {
			m.ready = true
		}

		m.width = msg.Width
		m.height = msg.Height

		helpHeight := m.help.Style.GetVerticalFrameSize() + m.help.Style.GetHeight()
		tabWindowHeight := m.height - helpHeight

		m.tabsWindow.Width = m.width
		m.tabsWindow.Height = max(0, tabWindowHeight)

		m.help.Width = m.width
	case entities.Response:
		cmds = append(cmds, listenToMessages(m.messages))
	}

	m.tabsWindow, cmd = m.tabsWindow.UpdateSpecific(msg)
	cmds = append(cmds, cmd)

	m.help, cmd = m.help.Update(msg)
	cmds = append(cmds, cmd)

	m.debug, cmd = m.debug.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	mainView := lipgloss.JoinVertical(lipgloss.Top,
		m.tabsWindow.View(),
		m.help.View(),
	)

	if m.debug.Active {
		debugView := m.debug.View()
		debugWidth := lipgloss.Width(debugView)
		debugHeight := lipgloss.Height(debugView)

		x := (m.width - debugWidth) / 2
		y := (m.height - debugHeight) / 2

		return overlay.PlaceOverlay(x, y, debugView, mainView, false)
	}

	return mainView
}
