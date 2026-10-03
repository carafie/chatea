package chat

import (
	"context"

	"github.com/carafie/chatea/api/errorsx"
	"github.com/carafie/chatea/api/pubsub"
	"github.com/carafie/chatea/api/wire/inbound"
	"github.com/carafie/chatea/api/wire/outbound"
)

var ErrConnInternal = errorsx.New("INTERNAL_ERROR", "connection internal error")

type ConnReader func(ctx context.Context) ([]byte, error)

type ConnWriter func(ctx context.Context, data []byte) error

type Conn struct {
	pubsub *pubsub.Connection
}

func NewConn(pubsub *pubsub.Connection) *Conn {
	return &Conn{pubsub: pubsub}
}

func (c *Conn) ListenAndServe(ctx context.Context, r ConnReader, w ConnWriter) {
	go func() {
		rawEvents := c.pubsub.Subscribe()
		for {
			select {
			case <-ctx.Done():
				return
			case rawEvent, ok := <-rawEvents:
				if !ok {
					return
				}
				if err := w(ctx, rawEvent); err != nil {
					return
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			rawEvent, err := r(ctx)
			if err != nil {
				return
			}
			c.handle(ctx, w, rawEvent)
		}
	}
}

func (c *Conn) handle(ctx context.Context, w ConnWriter, rawEvent []byte) {
	event, err := inbound.NewEvent(rawEvent)
	if err != nil {
		c.handleError(ctx, err, w)
		return
	}

	switch event.Type {
	case inbound.TypeMessage:
		if err := c.handleMessage(event); err != nil {
			c.handleError(ctx, err, w)
		}
	}
}

func (c *Conn) handleMessage(inEvent inbound.Event) error {
	inMessage, err := inbound.NewMessage(inEvent.Data)
	if err != nil {
		return err
	}
	guest, err := NewGuest(inMessage.GuestName, inMessage.GuestColor)
	if err != nil {
		return err
	}
	message, err := NewMessage(inMessage.Message)
	if err != nil {
		return err
	}

	outMessage := outbound.NewMessage(guest.Name, guest.Color, message)
	outEvent := outbound.NewEvent(outbound.TypeMessage, outMessage)
	rawOutEvent, err := outEvent.Raw()
	if err != nil {
		return err
	}
	c.pubsub.Publish(rawOutEvent)

	return nil
}

func (c *Conn) handleError(ctx context.Context, err error, w ConnWriter) {
	code := ErrConnInternal.Code
	if err, ok := errorsx.From(err); ok {
		code = err.Code
	}

	outEvent := outbound.NewEvent(
		outbound.TypeError,
		outbound.NewError(code),
	)
	rawOutEvent, err := outEvent.Raw()
	if err != nil {
		return
	}

	_ = w(ctx, rawOutEvent)
}

func (c *Conn) Close() {
	c.pubsub.Close()
}
