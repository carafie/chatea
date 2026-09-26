package service

import (
	"github.com/carafie/chatea/api/chat"
	"github.com/carafie/chatea/api/pubsub"
)

type Handler struct {
	pubsub *pubsub.Client
}

func NewHandler(pubsub *pubsub.Client) *Handler {
	if pubsub == nil {
		panic("service.NewHandler: *pubsub.Client cannot be nil")
	}
	return &Handler{pubsub: pubsub}
}

type ConnectParams struct {
	RoomName string
}

func (h *Handler) Connect(params ConnectParams) (*chat.Conn, error) {
	room, err := chat.NewRoom(params.RoomName)
	if err != nil {
		return nil, err
	}
	conn := h.pubsub.Connect(room.Name)
	return chat.NewConn(conn), nil
}
