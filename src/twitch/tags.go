package twitch

import (
	"fmt"
	"regexp"
	"strconv"
	"streamutils-tui/src/entities"
	"strings"
	"time"
)

const DefaultColor = "#FF4500"

func ParseTags(s string) entities.Tags {
	tagStrings := strings.Split(s, ";")

	tagsMap := make(map[string]string)
	for _, tagString := range tagStrings {
		tagKeyVal := strings.Split(tagString, "=")
		if len(tagKeyVal) != 2 {
			continue
		}
		tagsMap[tagKeyVal[0]] = tagKeyVal[1]
	}

	tags := entities.Tags{Color: DefaultColor}
	tags.Color = tagsMap["color"]
	tags.DisplayName = tagsMap["display-name"]
	tags.Time, _ = convertStringToTime(tagsMap["tmi-sent-ts"])

	return tags
}

// SeparateTags Separates tags from message.
func SeparateTags(message string) (newMessage string, tags string, err error) {
	newMessage = ""
	tags = ""
	err = nil

	re := regexp.MustCompile(`^(@(?P<Tags>[^ ]+) )?(?P<Message>.*)$`)
	match := re.FindStringSubmatch(message)
	if len(match) < 4 {
		err = fmt.Errorf("unable to separate tags: %s", message)
		return
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

	newMessage = paramsMap["Message"]
	tags = paramsMap["Tags"]
	return
}

// convertStringToTime converts a string Unix timestamp (seconds or milliseconds) to a time.Time object.
func convertStringToTime(timestampStr string) (time.Time, error) {
	// Try to parse as seconds first
	timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timestamp string: %w", err)
	}

	//Check length, if 10 is seconds if 13 is milliseconds if other length throws an error
	switch len(timestampStr) {
	case 10: // Seconds
		return time.Unix(timestampInt, 0), nil
	case 13: // Milliseconds
		return time.Unix(0, timestampInt*int64(time.Millisecond)), nil //Convert ms to nanoseconds
	default:
		return time.Time{}, fmt.Errorf("invalid timestamp length.  Must be 10 (seconds) or 13 (milliseconds) digits. Length: %d", len(timestampStr))
	}
}
