package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*RegularPolygon)(nil)

type RegularPolygon struct {
	baseObject

	FillColor    color.Color
	StrokeColor  color.Color
	StrokeWidth  float32
	CornerRadius float32
	Angle        float32
	Sides        uint
}

func (r *RegularPolygon) Hide() { _ = "STUB: not implemented"; return }

func (r *RegularPolygon) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (r *RegularPolygon) Refresh() { _ = "STUB: not implemented"; return }

func (r *RegularPolygon) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewRegularPolygon(sides uint, color color.Color) *RegularPolygon {
	_ = "STUB: not implemented"
	return nil
}
