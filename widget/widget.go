package widget

import (
	"fyne.io/fyne/v2"
)

type BaseWidget struct {
	noCopy noCopy

	size     fyne.Size
	position fyne.Position
	Hidden   bool

	impl       fyne.Widget
	themeCache fyne.Theme
}

func (w *BaseWidget) ExtendBaseWidget(wid fyne.Widget) { _ = "STUB: not implemented"; return }

func (w *BaseWidget) Size() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (w *BaseWidget) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (w *BaseWidget) Position() fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func (w *BaseWidget) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (w *BaseWidget) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (w *BaseWidget) Visible() bool { _ = "STUB: not implemented"; return false }

func (w *BaseWidget) Show() { _ = "STUB: not implemented"; return }

func (w *BaseWidget) Hide() { _ = "STUB: not implemented"; return }

func (w *BaseWidget) Refresh() { _ = "STUB: not implemented"; return }

func (w *BaseWidget) Theme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func (w *BaseWidget) super() fyne.Widget { _ = "STUB: not implemented"; return *new(fyne.Widget) }

type DisableableWidget struct {
	BaseWidget

	disabled bool
}

func (w *DisableableWidget) Enable() { _ = "STUB: not implemented"; return }

func (w *DisableableWidget) Disable() { _ = "STUB: not implemented"; return }

func (w *DisableableWidget) Disabled() bool { _ = "STUB: not implemented"; return false }

func NewSimpleRenderer(object fyne.CanvasObject) fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

type Orientation int

const (
	Horizontal Orientation = 0
	Vertical   Orientation = 1

	Adaptive Orientation = 2
)

type noCopy struct{}

func (*noCopy) Lock() { _ = "STUB: not implemented"; return }

func (*noCopy) Unlock() { _ = "STUB: not implemented"; return }
