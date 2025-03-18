package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"streamutils-tui/src"
	"streamutils-tui/src/twitch"
)

func main() {
	username, accessToken := src.LoadEnv()

	twitchClient := twitch.New(accessToken, username, "irc.chat.twitch.tv:6667")

	err := twitchClient.Join("jcwithc")
	if err != nil {
		panic(err.Error())
		return
	}

	responses := twitchClient.StreamChat()
	if err != nil {
		panic(err.Error())
		return
	}

	for response := range responses {
		if response.Error != nil {
			panic(response.Error.Error())
		}
		if response.Payload.Action != "PRIVMSG" {
			continue
		}

		name := response.Payload.Username
		if response.Tags.DisplayName != "" {
			name = response.Tags.DisplayName
		}
		colorizedName := colorizeString(name, response.Tags.Color)
		//fmt.Printf("%s#%s [%s]: %s\n", colorizedName, response.Payload.Channel, response.Payload.Action, response.Payload.Message) // TODO: Remove this
		fmt.Printf("%s %s: %s\n", response.Tags.Time.Format("15:04"), colorizedName, response.Payload.Message) // TODO: Remove this
	}
}

func colorizeString(text string, hexColor string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(hexColor))
	return style.Render(text)
}
