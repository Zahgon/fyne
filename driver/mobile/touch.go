package mobile

import "fyne.io/fyne/v2"

type TouchEvent struct {
	fyne.PointEvent

	ID int
}

type Touchable interface {
	TouchDown(*TouchEvent)
	TouchUp(*TouchEvent)
	TouchCancel(*TouchEvent)
}

type Movable interface {
	TouchMoved(*TouchEvent)
}
