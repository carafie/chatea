package chat

import (
	"errors"
	"unicode"

	"github.com/carafie/chatea/api/text"
)

var ErrMessageInvalid = errors.New("message is invalid")

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
	if err != nil {
		return "", ErrMessageInvalid
	}
	return Message(message), nil
}
