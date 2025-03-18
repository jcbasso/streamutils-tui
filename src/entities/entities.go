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
	Action   string
	Username string
	Channel  string
	Message  string
}
