package chat

import (
	"unicode"

	"github.com/carafie/chatea/api/errorsx"
	"github.com/carafie/chatea/api/text"
)

var (
	ErrMessageTooShort = errorsx.New("MESSAGE_TOO_SHORT", "message is too short")
	ErrMessageTooLong  = errorsx.New("MESSAGE_TOO_LONG", "message is too long")
	ErrMessageInvalid  = errorsx.New("MESSAGE_INVALID", "message is invalid")
)

var messageParser = text.Parser{
	MinChars: 1,
	MaxChars: 500,
	CharValid: func(char rune) bool {
		return unicode.IsPrint(char)
	},
}

type Message = string

func NewMessage(message string) (Message, error) {
	message, err := messageParser.Parse(message)
	switch err {
	case nil:
	case text.ErrTooShort:
		return "", ErrMessageTooShort
	case text.ErrTooLong:
		return "", ErrMessageTooLong
	default:
		return "", ErrMessageInvalid
	}
	return Message(message), nil
}
