package layout

import (
	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*paddedLayout)(nil)

type paddedLayout struct{}

func (l paddedLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (l paddedLayout) MinSize(objects []fyne.CanvasObject) (min fyne.Size) {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func NewPaddedLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }
