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
	Name  string
	Color string
}

func NewGuest(name, color string) (Guest, error) {
	name, err := guestNameParser.Parse(name)
	if err != nil {
		return Guest{}, ErrGuestNameInvalid
	}

	switch color {
	case ColorRed, ColorOrange, ColorGreen, ColorTeal, ColorBlue, ColorPurple, ColorPink:
	default:
		color = ColorRed
	}

	return Guest{Name: name, Color: color}, nil
}

const (
	ColorRed    = "red"
	ColorOrange = "orange"
	ColorGreen  = "green"
	ColorTeal   = "teal"
	ColorBlue   = "blue"
	ColorPurple = "purple"
	ColorPink   = "pink"
)
