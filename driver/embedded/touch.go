package embedded

import "fyne.io/fyne/v2"

type TouchDownEvent struct {
	Position fyne.Position
	ID       int
}

func (*TouchDownEvent) isEvent() { _ = "STUB: not implemented"; return }

type TouchMoveEvent struct {
	Position fyne.Position
	ID       int
}

func (*TouchMoveEvent) isEvent() { _ = "STUB: not implemented"; return }

type TouchUpEvent struct {
	Position fyne.Position
	ID       int
}

func (*TouchUpEvent) isEvent() { _ = "STUB: not implemented"; return }
