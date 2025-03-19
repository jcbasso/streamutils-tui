package entities

import "time"

type Response struct {
	Payload Payload
	Tags    Tags
	Error   error
}

type Tags struct {
	Color       string
	DisplayName string
	Time        time.Time
}

type Payload struct {
	Action   string // TODO: Move to enum
	Username string
	Channel  string
	Message  string
}
