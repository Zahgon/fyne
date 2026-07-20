package fyne

type CanvasObject interface {
	MinSize() Size

	Move(Position)

	Position() Position

	Resize(Size)

	Size() Size

	Hide()

	Visible() bool

	Show()

	Refresh()
}

type Disableable interface {
	Enable()
	Disable()
	Disabled() bool
}

type DoubleTappable interface {
	DoubleTapped(*PointEvent)
}

type Draggable interface {
	Dragged(*DragEvent)
	DragEnd()
}

type Focusable interface {
	FocusGained()

	FocusLost()

	TypedRune(rune)

	TypedKey(*KeyEvent)
}

type Scrollable interface {
	Scrolled(*ScrollEvent)
}

type SecondaryTappable interface {
	TappedSecondary(*PointEvent)
}

type Shortcutable interface {
	TypedShortcut(Shortcut)
}

type Tabbable interface {
	AcceptsTab() bool
}

type Tappable interface {
	Tapped(*PointEvent)
}
