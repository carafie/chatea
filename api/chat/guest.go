package chat

import (
	"errors"
	"unicode"

	"github.com/carafie/chatea/api/text"
)

var ErrGuestNameInvalid = errors.New("guest name is invalid")

var guestNameParser = text.Parser{
	MinChars: 2,
	MaxChars: 20,
	CharValid: func(char rune) bool {
		return unicode.IsPrint(char) && !unicode.IsSpace(char)
	},
}

type Guest struct {
	Name string
}

func NewGuest(name string) (Guest, error) {
	name, err := guestNameParser.Parse(name)
	if err != nil {
		return Guest{}, ErrGuestNameInvalid
	}
	return Guest{Name: name}, nil
}
