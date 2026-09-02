package mobile

import (
	"fyne.io/fyne/v2"
)

type KeyboardType int32

const (
	DefaultKeyboard KeyboardType = iota

	SingleLineKeyboard

	NumberKeyboard

	PasswordKeyboard
)

type Keyboardable interface {
	fyne.Focusable

	Keyboard() KeyboardType
}
