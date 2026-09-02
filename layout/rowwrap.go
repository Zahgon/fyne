package layout

import (
	"fyne.io/fyne/v2"
)

type rowWrapLayout struct {
	horizontalPadding float32
	minSize           fyne.Size
	verticalPadding   float32
}

func NewRowWrapLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func NewRowWrapLayoutWithCustomPadding(horizontal, vertical float32) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

var _ fyne.Layout = (*rowWrapLayout)(nil)

func (l *rowWrapLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (l *rowWrapLayout) minHeight(rowHeight float32, rowCount int) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (l *rowWrapLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	_ = "STUB: not implemented"
	return
}
