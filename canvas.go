package fyne

import "image"

type Canvas interface {
	Content() CanvasObject
	SetContent(CanvasObject)

	Refresh(CanvasObject)

	Focus(Focusable)

	FocusNext()

	FocusPrevious()
	Unfocus()
	Focused() Focusable

	Size() Size

	Scale() float32

	Overlays() OverlayStack

	OnTypedRune() func(rune)
	SetOnTypedRune(func(rune))
	OnTypedKey() func(*KeyEvent)
	SetOnTypedKey(func(*KeyEvent))
	AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut))
	RemoveShortcut(shortcut Shortcut)

	Capture() image.Image

	PixelCoordinateForPosition(Position) (int, int)

	InteractiveArea() (Position, Size)
}
