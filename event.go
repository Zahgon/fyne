package fyne

type HardwareKey struct {
	ScanCode int
}

type KeyEvent struct {
	Name KeyName

	Physical HardwareKey
}

type PointEvent struct {
	AbsolutePosition Position
	Position         Position
}

type ScrollEvent struct {
	PointEvent
	Scrolled Delta
}

type DragEvent struct {
	PointEvent
	Dragged Delta
}
