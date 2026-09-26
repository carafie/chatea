package outbound

import (
	"encoding/json/v2"
	"time"

	"github.com/carafie/chatea/api/expected"
)

var ErrEventInvalid = expected.New("OUTBOUND_EVENT_INVALID", "outbound event is invalid")

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

func (e Event) Raw() ([]byte, error) {
	raw, err := json.Marshal(e)
	if err != nil {
		return nil, ErrEventInvalid
	}
	return raw, nil
}

type Type string

const (
	TypeMessage Type = "message"
	TypeError   Type = "error"
)

type Message struct {
	GuestName  string    `json:"guest_name"`
	GuestColor string    `json:"guest_color"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewMessage(guestName, guestColor, message string) Message {
	return Message{
		GuestName:  guestName,
		GuestColor: guestColor,
		Message:    message,
		CreatedAt:  time.Now().UTC(),
	}
}

type Error struct {
	Code string `json:"code"`
}

func NewError(code string) Error {
	return Error{Code: code}
}
