package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Arc)(nil)

type Arc struct {
	baseObject

	FillColor    color.Color
	StartAngle   float32
	EndAngle     float32
	CornerRadius float32
	StrokeColor  color.Color
	StrokeWidth  float32
	CutoutRatio  float32
}

func (a *Arc) Hide() { _ = "STUB: not implemented"; return }

func (a *Arc) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (a *Arc) Refresh() { _ = "STUB: not implemented"; return }

func (a *Arc) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewArc(startAngle, endAngle, cutoutRatio float32, color color.Color) *Arc {
	_ = "STUB: not implemented"
	return nil
}

func NewPieArc(startAngle, endAngle float32, color color.Color) *Arc {
	_ = "STUB: not implemented"
	return nil
}

func NewDoughnutArc(startAngle, endAngle float32, color color.Color) *Arc {
	_ = "STUB: not implemented"
	return nil
}
