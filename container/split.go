package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.CanvasObject = (*Split)(nil)

var splitDefaultThickness = float32(5)

type Split struct {
	widget.BaseWidget
	Offset     float64
	Horizontal bool
	Leading    fyne.CanvasObject
	Trailing   fyne.CanvasObject

	offsetUpdated bool
}

func NewHSplit(leading, trailing fyne.CanvasObject) *Split { _ = "STUB: not implemented"; return nil }

func NewVSplit(top, bottom fyne.CanvasObject) *Split { _ = "STUB: not implemented"; return nil }

func newSplitContainer(horizontal bool, leading, trailing fyne.CanvasObject) *Split {
	_ = "STUB: not implemented"
	return nil
}

func (s *Split) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (s *Split) ExtendBaseWidget(wid fyne.Widget) { _ = "STUB: not implemented"; return }

func (s *Split) SetOffset(offset float64) { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*splitContainerRenderer)(nil)

type splitContainerRenderer struct {
	split   *Split
	divider *divider
	objects []fyne.CanvasObject
}

func (*splitContainerRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *splitContainerRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *splitContainerRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *splitContainerRenderer) Objects() []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func (r *splitContainerRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *splitContainerRenderer) computeSplitLengths(total, lMin, tMin float32) (leading, trailing float32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (r *splitContainerRenderer) minLeadingWidth() float32 { _ = "STUB: not implemented"; return 0 }

func (r *splitContainerRenderer) minLeadingHeight() float32 { _ = "STUB: not implemented"; return 0 }

func (r *splitContainerRenderer) minTrailingWidth() float32 { _ = "STUB: not implemented"; return 0 }

func (r *splitContainerRenderer) minTrailingHeight() float32 { _ = "STUB: not implemented"; return 0 }

var (
	_ fyne.CanvasObject  = (*divider)(nil)
	_ fyne.Draggable     = (*divider)(nil)
	_ desktop.Cursorable = (*divider)(nil)
	_ desktop.Hoverable  = (*divider)(nil)
)

type divider struct {
	widget.BaseWidget
	split          *Split
	hovered        bool
	startDragOff   *fyne.Position
	currentDragPos fyne.Position
}

func newDivider(split *Split) *divider { _ = "STUB: not implemented"; return nil }

func (d *divider) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (d *divider) Cursor() desktop.Cursor { _ = "STUB: not implemented"; return *new(desktop.Cursor) }

func (d *divider) DragEnd() { _ = "STUB: not implemented"; return }

func (d *divider) Dragged(e *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (d *divider) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (*divider) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (d *divider) MouseOut() { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*dividerRenderer)(nil)

type dividerRenderer struct {
	divider    *divider
	background *canvas.Rectangle
	foreground *canvas.Rectangle
	objects    []fyne.CanvasObject
}

func (*dividerRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *dividerRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *dividerRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *dividerRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *dividerRenderer) Refresh() { _ = "STUB: not implemented"; return }

func dividerTheme(d *divider) fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func dividerThickness(d *divider) float32 { _ = "STUB: not implemented"; return 0 }

func dividerLength(d *divider) float32 { _ = "STUB: not implemented"; return 0 }

//revive:disable-line:add-constant

func handleThickness(d *divider) float32 { _ = "STUB: not implemented"; return 0 }

func handleLength(d *divider) float32 { _ = "STUB: not implemented"; return 0 }
