package layout

import (
	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*borderLayout)(nil)

type borderLayout struct {
	top, bottom, left, right fyne.CanvasObject
}

func NewBorderLayout(top, bottom, left, right fyne.CanvasObject) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func (b *borderLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (b *borderLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
