package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func makeAnimationScreen(_ fyne.Window) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func makeAnimationCanvas() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func makeAnimationCurves() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func makeAnimationCurveItem(label string, curve fyne.AnimationCurve, yOff float32) (
	text *widget.Label, box fyne.CanvasObject, anim *fyne.Animation,
) {
	_ = "STUB: not implemented"
	return nil, *new(fyne.CanvasObject), nil
}

type themedBox struct {
	widget.BaseWidget
}

func newThemedBox() *themedBox { _ = "STUB: not implemented"; return nil }

func (b *themedBox) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

type themedBoxRenderer struct {
	bg      *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (r *themedBoxRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *themedBoxRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *themedBoxRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *themedBoxRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *themedBoxRenderer) Refresh() { _ = "STUB: not implemented"; return }
