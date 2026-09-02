package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Ellipse)(nil)

type Ellipse struct {
	baseObject

	FillColor   color.Color
	StrokeColor color.Color
	StrokeWidth float32
	Shadow      Shadow
}

func (e *Ellipse) Hide() { _ = "STUB: not implemented"; return }

func (e *Ellipse) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (e *Ellipse) Refresh() { _ = "STUB: not implemented"; return }

func (e *Ellipse) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewEllipse(c color.Color) *Ellipse { _ = "STUB: not implemented"; return nil }
