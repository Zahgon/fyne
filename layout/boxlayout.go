package layout

import (
	"fyne.io/fyne/v2"
)

func NewVBoxLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func NewHBoxLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func NewCustomPaddedHBoxLayout(padding float32) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func NewCustomPaddedVBoxLayout(padding float32) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

var _ fyne.Layout = (*vBoxLayout)(nil)

type vBoxLayout struct {
	paddingFunc func() float32
}

func (v vBoxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (v vBoxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

var _ fyne.Layout = (*hBoxLayout)(nil)

type hBoxLayout struct {
	paddingFunc func() float32
}

func (g hBoxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (g hBoxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func isVerticalSpacer(obj fyne.CanvasObject) bool { _ = "STUB: not implemented"; return false }

func isHorizontalSpacer(obj fyne.CanvasObject) bool { _ = "STUB: not implemented"; return false }
