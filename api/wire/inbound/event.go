package inbound

import (
	"encoding/json"
	"encoding/json/jsontext"
)

type RawEvent = []byte

type Event struct {
	Type Type           `json:"type"`
	Data jsontext.Value `json:"data"`
}

func ParseEvent(rawEvent RawEvent) (Event, error) {
	var event Event
	err := json.Unmarshal(rawEvent, &event)
	return event, err
}

type Type string

const TypeMessage Type = "message"

type Message struct {
	GuestName string `json:"guest_name"`
	Message   string `json:"message"`
}

func ParseMessage(rawMessage []byte) (Message, error) {
	var message Message
	err := json.Unmarshal(rawMessage, &message)
	return message, err
}
