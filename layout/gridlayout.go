package layout

import (
	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*gridLayout)(nil)

type gridLayout struct {
	Cols            int
	vertical, adapt bool
}

func NewAdaptiveGridLayout(rowcols int) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func NewGridLayout(cols int) fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func NewGridLayoutWithColumns(cols int) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func NewGridLayoutWithRows(rows int) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func (g *gridLayout) horizontal() bool { _ = "STUB: not implemented"; return false }

func (g *gridLayout) countRows(objects []fyne.CanvasObject) int {
	_ = "STUB: not implemented"
	return 0
}

func getLeading(size float64, offset int) float32 { _ = "STUB: not implemented"; return 0 }

func getTrailing(size float64, offset int) float32 { _ = "STUB: not implemented"; return 0 }

func (g *gridLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (g *gridLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
