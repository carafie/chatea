package chat

import (
	"unicode"

	"github.com/carafie/chatea/api/expected"
	"github.com/carafie/chatea/api/text"
)

var (
	ErrGuestNameTooShort = expected.New("GUEST_NAME_TOO_SHORT", "guest name is too short")
	ErrGuestNameTooLong  = expected.New("GUEST_NAME_TOO_LONG", "guest name is too long")
	ErrGuestNameInvalid  = expected.New("GUEST_NAME_INVALID", "guest name is invalid")
)

var guestNameParser = text.Parser{
	MinChars: 2,
	MaxChars: 20,
	CharValid: func(char rune) bool {
		return unicode.IsPrint(char) && !unicode.IsSpace(char)
	},
}

type Guest struct {
	Name  string
	Color Color
}

func NewGuest(name, color string) (Guest, error) {
	name, err := guestNameParser.Parse(name)
	switch err {
	case nil:
	case text.ErrTooShort:
		return Guest{}, ErrGuestNameTooShort
	case text.ErrTooLong:
		return Guest{}, ErrGuestNameTooLong
	default:
		return Guest{}, ErrGuestNameInvalid
	}
	return Guest{Name: name, Color: NewColor(color)}, nil
}

type Color = string

func NewColor(color string) Color {
	switch color {
	case ColorRed, ColorOrange, ColorGreen, ColorTeal, ColorBlue, ColorPurple, ColorPink:
		return color
	default:
		return ColorRed
	}
}

const (
	ColorRed    Color = "red"
	ColorOrange Color = "orange"
	ColorGreen  Color = "green"
	ColorTeal   Color = "teal"
	ColorBlue   Color = "blue"
	ColorPurple Color = "purple"
	ColorPink   Color = "pink"
)
