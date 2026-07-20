package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
)

var _ fyne.Widget = (*PopUp)(nil)

type PopUp struct {
	BaseWidget

	Content fyne.CanvasObject
	Canvas  fyne.Canvas

	overlay       *widget.OverlayContainer
	modal, manual bool
}

func (p *PopUp) Hide() { _ = "STUB: not implemented"; return }

func (p *PopUp) Refresh() { _ = "STUB: not implemented"; return }

func (p *PopUp) Show() { _ = "STUB: not implemented"; return }

func (p *PopUp) ShowAtPosition(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (p *PopUp) ShowAtRelativePosition(rel fyne.Position, to fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return
}

func (p *PopUp) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (p *PopUp) TappedSecondary(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (p *PopUp) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (p *PopUp) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func ShowPopUpAtPosition(content fyne.CanvasObject, canvas fyne.Canvas, pos fyne.Position) {
	_ = "STUB: not implemented"
	return
}

func ShowPopUpAtRelativePosition(content fyne.CanvasObject, canvas fyne.Canvas, rel fyne.Position, to fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return
}

func newPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp {
	_ = "STUB: not implemented"
	return nil
}

func NewPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp {
	_ = "STUB: not implemented"
	return nil
}

func ShowPopUp(content fyne.CanvasObject, canvas fyne.Canvas) { _ = "STUB: not implemented"; return }

func newModalPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp {
	_ = "STUB: not implemented"
	return nil
}

func NewModalPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp {
	_ = "STUB: not implemented"
	return nil
}

func ShowModalPopUp(content fyne.CanvasObject, canvas fyne.Canvas) {
	_ = "STUB: not implemented"
	return
}

type popUpBaseRenderer struct {
	popUp      *PopUp
	background *canvas.Rectangle
}

func (r *popUpBaseRenderer) padding() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

type popUpRenderer struct {
	widget.BaseRenderer
	popUpBaseRenderer
}

func (r *popUpRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *popUpRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *popUpRenderer) Refresh() { _ = "STUB: not implemented"; return }

func withRelativePosition(rel fyne.Position, to fyne.CanvasObject, f func(position fyne.Position)) {
	_ = "STUB: not implemented"
	return
}
