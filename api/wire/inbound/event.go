package inbound

import (
	"encoding/json/jsontext"
	"encoding/json/v2"

	"github.com/carafie/chatea/api/errorsx"
)

var (
	ErrEventInvalid   = errorsx.New("INBOUND_EVENT_INVALID", "inbound event is invalid")
	ErrMessageInvalid = errorsx.New("INBOUND_MESSAGE_INVALID", "inbound message is invalid")
)

type Event struct {
	Type Type           `json:"type"`
	Data jsontext.Value `json:"data"`
}

func NewEvent(rawEvent []byte) (Event, error) {
	var event Event
	if err := json.Unmarshal(rawEvent, &event); err != nil {
		return Event{}, ErrEventInvalid
	}
	return event, nil
}

type Type string

const TypeMessage Type = "message"

type Message struct {
	GuestName  string `json:"guest_name"`
	GuestColor string `json:"guest_color"`
	Message    string `json:"message"`
}

func NewMessage(rawMessage []byte) (Message, error) {
	var message Message
	if err := json.Unmarshal(rawMessage, &message); err != nil {
		return Message{}, ErrMessageInvalid
	}
	return message, nil
}
