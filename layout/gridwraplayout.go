package layout

import (
	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*gridWrapLayout)(nil)

type gridWrapLayout struct {
	CellSize fyne.Size
	colCount int
	rowCount int
}

func NewGridWrapLayout(size fyne.Size) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func (g *gridWrapLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (g *gridWrapLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
