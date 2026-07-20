package layout

import "fyne.io/fyne/v2"

var _ fyne.Layout = (*centerLayout)(nil)

type centerLayout struct{}

func NewCenterLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func (c *centerLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (c *centerLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
