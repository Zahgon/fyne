package layout

import (
	"fyne.io/fyne/v2"
)

const formLayoutCols = 2

var _ fyne.Layout = (*formLayout)(nil)

type formLayout struct{}

func (f *formLayout) calculateTableSizes(objects []fyne.CanvasObject, containerWidth float32) (labelWidth float32, contentWidth float32, height float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func (f *formLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (f *formLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func NewFormLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }
