package outbound

import (
	"encoding/json/v2"
	"time"
)

type RawEvent = []byte

type Event struct {
	Type Type `json:"type"`
	Data any  `json:"data"`
}

func NewEvent(typ Type, data any) Event {
	return Event{
		Type: typ,
		Data: data,
	}
}

func (e Event) Format() (RawEvent, error) {
	return json.Marshal(e)
}

type Type string

const TypeMessage Type = "message"

type Message struct {
	GuestName string    `json:"guest_name"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(guestName, message string) Message {
	return Message{
		GuestName: guestName,
		Message:   message,
		CreatedAt: time.Now().UTC(),
	}
}
