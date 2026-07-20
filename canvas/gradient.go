package canvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
)

type LinearGradient struct {
	baseObject

	StartColor color.Color
	EndColor   color.Color
	Angle      float64
}

func (g *LinearGradient) Generate(iw, ih int) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func (g *LinearGradient) Hide() { _ = "STUB: not implemented"; return }

func (g *LinearGradient) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (g *LinearGradient) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (g *LinearGradient) Refresh() { _ = "STUB: not implemented"; return }

type RadialGradient struct {
	baseObject

	StartColor color.Color
	EndColor   color.Color

	CenterOffsetX, CenterOffsetY float64
}

func (g *RadialGradient) Generate(iw, ih int) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func (g *RadialGradient) Hide() { _ = "STUB: not implemented"; return }

func (g *RadialGradient) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (g *RadialGradient) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (g *RadialGradient) Refresh() { _ = "STUB: not implemented"; return }

func calculatePixel(d float64, startColor, endColor color.Color) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func computeGradient(generator func(x, y float64) float64, w, h int, startColor, endColor color.Color) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func NewHorizontalGradient(start, end color.Color) *LinearGradient {
	_ = "STUB: not implemented"
	return nil
}

func NewLinearGradient(start, end color.Color, angle float64) *LinearGradient {
	_ = "STUB: not implemented"
	return nil
}

func NewRadialGradient(start, end color.Color) *RadialGradient {
	_ = "STUB: not implemented"
	return nil
}

func NewVerticalGradient(start color.Color, end color.Color) *LinearGradient {
	_ = "STUB: not implemented"
	return nil
}
