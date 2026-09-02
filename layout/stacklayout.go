package layout

import (
	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*stackLayout)(nil)

type stackLayout struct{}

func NewStackLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func NewMaxLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func (stackLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (stackLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
