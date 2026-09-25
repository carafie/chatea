package chat

import (
	"errors"
	"unicode"

	"github.com/carafie/chatea/api/text"
)

var ErrRoomNameInvalid = errors.New("room name is invalid")

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
	if err != nil {
		return Room{}, ErrRoomNameInvalid
	}
	return Room{Name: name}, nil
}
