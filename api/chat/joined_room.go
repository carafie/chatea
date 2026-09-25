package chat

import (
	"errors"

	"github.com/carafie/chatea/api/pubsub"
	"github.com/carafie/chatea/api/wire/inbound"
	"github.com/carafie/chatea/api/wire/outbound"
)

var ErrEventInvalid = errors.New("event is invalid")

type JoinedRoom struct {
	conn *pubsub.Connection
}

func JoinRoom(conn *pubsub.Connection) *JoinedRoom {
	return &JoinedRoom{conn: conn}
}

func (r *JoinedRoom) Send(rawEvent inbound.RawEvent) error {
	inEvent, err := inbound.ParseEvent(rawEvent)
	if err != nil {
		return ErrEventInvalid
	}

	switch inEvent.Type {
	case inbound.TypeMessage:
		if err := r.sendMessage(inEvent); err != nil {
			return err
		}
	default:
		return ErrEventInvalid
	}

	return nil
}

func (r *JoinedRoom) sendMessage(inEvent inbound.Event) error {
	inMessage, err := inbound.ParseMessage(inEvent.Data)
	if err != nil {
		return ErrEventInvalid
	}
	guest, err := NewGuest(inMessage.GuestName)
	if err != nil {
		return err
	}
	message, err := NewMessage(inMessage.Message)
	if err != nil {
		return err
	}

	outEvent := outbound.NewEvent(
		outbound.TypeMessage,
		outbound.NewMessage(guest.Name, message),
	)
	rawOutEvent, err := outEvent.Format()
	if err != nil {
		return ErrEventInvalid
	}
	r.conn.Publish(rawOutEvent)

	return nil
}

func (r *JoinedRoom) Receive() <-chan outbound.RawEvent {
	return r.conn.Subscribe()
}

func (r *JoinedRoom) Leave() {
	r.conn.Close()
}
