package messages

import "time"

type message struct {
	Username  string
	Timestamp time.Time
	Text      string
}
