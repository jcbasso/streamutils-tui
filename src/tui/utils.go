package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"streamutils-tui/src/entities"
	"streamutils-tui/src/tui/models/chat"
	"streamutils-tui/src/tui/models/sizeable"
	"streamutils-tui/src/tui/models/viewport"
)

func generateSizeableChat() sizeable.Model {
	chatModel := chat.New()
	return sizeable.NewWrapper(
		&chatModel,
		func(w int) { chatModel.Width = w },
		func(h int) { chatModel.Height = h },
		func() int { return chatModel.Width },
		func() int { return chatModel.Height },
		func() []key.Binding { return chatModel.ShortHelp() },
		func() [][]key.Binding { return chatModel.FullHelp() },
	)
}

func generateSizeableViewport(content string) sizeable.Model {
	model := viewport.New(0, 0, content)
	return sizeable.NewWrapper(
		&model,
		func(w int) { model.Model.Width = w },
		func(h int) { model.Model.Height = h },
		func() int { return model.Model.Width },
		func() int { return model.Model.Height },
		func() []key.Binding { return model.ShortHelp() },
		func() [][]key.Binding { return model.FullHelp() },
	)
}

func listenToMessages(messages chan entities.Response) tea.Cmd {
	return func() tea.Msg {
		return <-messages
	}
}
