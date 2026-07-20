package layout

import (
	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*CustomPaddedLayout)(nil)

type CustomPaddedLayout struct {
	TopPadding    float32
	BottomPadding float32
	LeftPadding   float32
	RightPadding  float32
}

func NewCustomPaddedLayout(padTop, padBottom, padLeft, padRight float32) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func (c CustomPaddedLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (c CustomPaddedLayout) MinSize(objects []fyne.CanvasObject) (min fyne.Size) {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
