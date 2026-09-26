package chat

import (
	"unicode"

	"github.com/carafie/chatea/api/expected"
	"github.com/carafie/chatea/api/text"
)

var (
	ErrRoomNameTooShort = expected.New("ROOM_NAME_TOO_SHORT", "room name is too short")
	ErrRoomNameTooLong  = expected.New("ROOM_NAME_TOO_LONG", "room name is too long")
	ErrRoomNameInvalid  = expected.New("ROOM_NAME_INVALID", "room name is invalid")
)

var roomNameParser = text.Parser{
	MinChars: 2,
	MaxChars: 20,
	CharValid: func(char rune) bool {
		return unicode.IsPrint(char) && !unicode.IsSpace(char)
	},
}

type Room struct {
	Name string
}

func NewRoom(name string) (Room, error) {
	name, err := roomNameParser.Parse(name)
	switch err {
	case nil:
	case text.ErrTooShort:
		return Room{}, ErrRoomNameTooShort
	case text.ErrTooLong:
		return Room{}, ErrRoomNameTooLong
	default:
		return Room{}, ErrRoomNameInvalid
	}
	return Room{Name: name}, nil
}
