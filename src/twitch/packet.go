package twitch

import (
	"fmt"
	"regexp"
	"streamutils-tui/src/entities"
)

func ParsePacket(message string) (entities.Payload, error) {
	re := regexp.MustCompile("^:(?P<Username>\\w+)![^ ]*\\s*(?P<Action>\\w+)\\s*(#(?P<Channel>\\w+))?\\s*(:(?P<Message>.*))?$")
	match := re.FindStringSubmatch(message)
	if len(match) < 6 {
		return entities.Payload{}, fmt.Errorf("unable to parse packet: %s", message)
	}

	paramsMap := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if name == "" {
			continue
		}
		if i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	res := entities.Payload{
		Action:   paramsMap["Action"],
		Username: paramsMap["Username"],
		Channel:  paramsMap["Channel"],
		Message:  paramsMap["Message"],
	}
	return res, nil
}
