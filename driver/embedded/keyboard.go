package embedded

import "fyne.io/fyne/v2"

type KeyDirection uint8

const (
	KeyPressed KeyDirection = iota

	KeyReleased
)

type KeyEvent struct {
	Name      fyne.KeyName
	Direction KeyDirection
}

func (*KeyEvent) isEvent() { _ = "STUB: not implemented"; return }

type CharacterEvent struct {
	Rune rune
}

func (*CharacterEvent) isEvent() { _ = "STUB: not implemented"; return }
