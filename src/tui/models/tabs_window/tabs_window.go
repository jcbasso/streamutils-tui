package tabs_window

import (
	teaHelp "github.com/charmbracelet/bubbles/help"
	"streamutils-tui/src/tui/models/sizeable"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = (*Model)(nil)
var _ teaHelp.KeyMap = (*Model)(nil)

type Tab struct {
	Title string
	Model sizeable.Model
}

type Model struct {
	Tabs             []Tab
	ActiveTab        int
	Width            int
	Height           int
	Style            lipgloss.Style // Style for the overall TabsWindow
	ActiveTabStyle   lipgloss.Style // Style for the active tab
	InactiveTabStyle lipgloss.Style // Style for the inactive tabs
	Separator        string
	SeparatorStyle   lipgloss.Style // Style for the separator
	KeyMap           KeyMap
	recursiveKeyMaps []teaHelp.KeyMap
}

func New(tabs []Tab) Model {
	separatorStyle := lipgloss.NewStyle().Foreground(DefaultTabsWindowStyle().GetBorderTopForeground()) // Use tabsWindowStyle
	defaultKeyMap := DefaultKeyMap()
	keyMaps := []teaHelp.KeyMap{defaultKeyMap}
	for _, tab := range tabs {
		keyMaps = append(keyMaps, tab.Model)
	}

	return Model{
		Tabs:             tabs,
		KeyMap:           defaultKeyMap,
		Separator:        DefaultTabsWindowStyle().GetBorderStyle().Top,
		SeparatorStyle:   separatorStyle,
		Style:            DefaultTabsWindowStyle(),
		ActiveTabStyle:   DefaultActiveTabStyle(),
		InactiveTabStyle: DefaultInactiveTabStyle(),
		recursiveKeyMaps: keyMaps,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		for i := range m.Tabs {
			m.Tabs[i].Model.SetWidth(m.Width - m.Style.GetHorizontalFrameSize())
			m.Tabs[i].Model.SetHeight(m.Height - m.Style.GetVerticalFrameSize() - m.ActiveTabStyle.GetHeight())
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Tab):
			m.ActiveTab = (m.ActiveTab + 1) % len(m.Tabs)
		case key.Matches(msg, m.KeyMap.ShiftTab):
			m.ActiveTab = (((m.ActiveTab - 1) % len(m.Tabs)) + len(m.Tabs)) % len(m.Tabs)
		}
	}

	// Update only the active tab's model.
	activeModel, cmd := m.Tabs[m.ActiveTab].Model.Update(msg)
	m.Tabs[m.ActiveTab].Model = activeModel
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) UpdateSpecific(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	_, cmd = m.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var (
		tabs            []string
		border          = m.Style.GetBorderStyle()
		borderColor     = lipgloss.NewStyle().Foreground(m.Style.GetBorderTopForeground()).Background(m.Style.GetBorderTopBackground())
		topBorderStr    string
		tabsBorderSpace = 1 // Space that separates the start and end of the tabs from the left and right borders
	)

	for i, tab := range m.Tabs {
		var tabStyle lipgloss.Style
		if i == m.ActiveTab {
			tabStyle = m.ActiveTabStyle
		} else {
			tabStyle = m.InactiveTabStyle
		}
		tabs = append(tabs, tabStyle.Render(tab.Title))
		if i < len(m.Tabs)-1 {
			tabs = append(tabs, m.SeparatorStyle.Render(m.Separator))
		}
	}
	tabsRow := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

	availableWidth := m.Width - m.Style.GetHorizontalBorderSize()
	activeTitleRow := m.ActiveTabStyle.Render(m.Tabs[m.ActiveTab].Title)

	if availableWidth >= lipgloss.Width(tabsRow) {
		topBorderStr = buildTopRow(tabsRow, availableWidth, m.Style, tabsBorderSpace)
	} else if availableWidth >= lipgloss.Width(activeTitleRow) {
		topBorderStr = buildTopRow(activeTitleRow, availableWidth, m.Style, tabsBorderSpace)
	} else {
		more := "…"
		row := lipgloss.NewStyle().MaxWidth(availableWidth-lipgloss.Width(more)).Render(activeTitleRow) + more
		topBorderStr = buildTopRow(row, availableWidth, m.Style, tabsBorderSpace)
	}
	topView := lipgloss.JoinHorizontal(lipgloss.Left, borderColor.Render(border.TopLeft), topBorderStr, borderColor.Render(border.TopRight))
	view := removeFirstRow(m.Style.Render(m.Tabs[m.ActiveTab].Model.View()))

	return lipgloss.JoinVertical(lipgloss.Top, topView, view)
}

func (m Model) ShortHelp() []key.Binding {
	var bindings []key.Binding
	for _, elem := range m.recursiveKeyMaps {
		bindings = append(bindings, elem.ShortHelp()...)
	}
	return bindings
}

func (m Model) FullHelp() [][]key.Binding {
	var bindings [][]key.Binding
	for _, elem := range m.recursiveKeyMaps {
		bindings = append(bindings, elem.FullHelp()...)
	}
	return bindings
}
