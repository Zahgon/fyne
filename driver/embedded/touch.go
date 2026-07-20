package embedded

import "fyne.io/fyne/v2"

type TouchDownEvent struct {
	Position fyne.Position
	ID       int
}

func (t *TouchDownEvent) isEvent() { _ = "STUB: not implemented"; return }

type TouchMoveEvent struct {
	Position fyne.Position
	ID       int
}

func (t *TouchMoveEvent) isEvent() { _ = "STUB: not implemented"; return }

type TouchUpEvent struct {
	Position fyne.Position
	ID       int
}

func (t *TouchUpEvent) isEvent() { _ = "STUB: not implemented"; return }
