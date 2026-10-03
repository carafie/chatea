package web

import (
	"context"
	"net/http"
	"time"

	"github.com/carafie/chatea/api/errorsx"
	"github.com/carafie/chatea/api/service"
	"github.com/coder/websocket"
)

type Handler struct {
	serviceHandler      *service.Handler
	connectWriteTimeout time.Duration
}

func NewHandler(serviceHandler *service.Handler, connectWriteTimeout time.Duration) *Handler {
	if serviceHandler == nil {
		panic("web.NewHandler: *service.Handler cannot be nil")
	}
	return &Handler{
		serviceHandler:      serviceHandler,
		connectWriteTimeout: connectWriteTimeout,
	}
}

type ConnectParams struct {
	RoomName string
}

func (h *Handler) Connect(w http.ResponseWriter, r *http.Request) {
	params := ConnectParams{RoomName: r.PathValue(roomNamePathKey)}

	chatConn, err := h.serviceHandler.Connect(service.ConnectParams{
		RoomName: params.RoomName,
	})
	if err != nil {
		resp := response{statusCode: http.StatusInternalServerError}
		if err, ok := errorsx.From(err); ok {
			resp.statusCode = http.StatusBadRequest
			resp.body = errorBody{Code: err.Code}
		}
		resp.respond(w)
		return
	}
	defer chatConn.Close()

	wsConn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		return
	}
	defer wsConn.CloseNow()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reader := func(ctx context.Context) ([]byte, error) {
		_, data, err := wsConn.Read(ctx)
		return data, err
	}
	writer := func(ctx context.Context, data []byte) error {
		ctx, cancel := context.WithTimeout(ctx, h.connectWriteTimeout)
		defer cancel()
		return wsConn.Write(ctx, websocket.MessageText, data)
	}

	chatConn.ListenAndServe(ctx, reader, writer)
}
