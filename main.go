package main

import (
	"streamutils-tui/src"
	"streamutils-tui/src/tui"
	"streamutils-tui/src/twitch"
)

func main() {
	username, accessToken, channel := src.LoadEnv()

	twitchClient := twitch.New(accessToken, username, "irc.chat.twitch.tv:6667")

	err := twitchClient.Join(channel)
	if err != nil {
		panic(err.Error())
		return
	}

	messages := twitchClient.StreamChat()
	if err != nil {
		panic(err.Error())
		return
	}
	//c := make(chan entities.Response)
	program := tui.New(messages)
	program.Run()
}
