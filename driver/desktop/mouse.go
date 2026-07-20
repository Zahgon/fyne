package desktop

import "fyne.io/fyne/v2"

type MouseButton int

const (
	MouseButtonPrimary MouseButton = 1 << iota

	MouseButtonSecondary

	MouseButtonTertiary

	LeftMouseButton = MouseButtonPrimary

	RightMouseButton = MouseButtonSecondary
)

type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier fyne.KeyModifier
}

type Mouseable interface {
	MouseDown(*MouseEvent)
	MouseUp(*MouseEvent)
}

type Hoverable interface {
	MouseIn(*MouseEvent)

	MouseMoved(*MouseEvent)

	MouseOut()
}
