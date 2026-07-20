package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*ArbitraryPolygon)(nil)

type ArbitraryPolygon struct {
	baseObject

	Points           []fyne.Position
	NormalizedPoints bool
	CornerRadii      []float32
	FillColor        color.Color
	StrokeColor      color.Color
	StrokeWidth      float32
}

func (p *ArbitraryPolygon) Hide() { _ = "STUB: not implemented"; return }

func (p *ArbitraryPolygon) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (p *ArbitraryPolygon) Refresh() { _ = "STUB: not implemented"; return }

func (p *ArbitraryPolygon) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewArbitraryPolygon(points []fyne.Position, fill color.Color) *ArbitraryPolygon {
	_ = "STUB: not implemented"
	return nil
}
