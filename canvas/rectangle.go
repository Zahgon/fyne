package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Rectangle)(nil)

type Rectangle struct {
	baseObject

	FillColor   color.Color
	StrokeColor color.Color
	StrokeWidth float32

	CornerRadius float32

	Aspect float32

	TopRightCornerRadius float32

	TopLeftCornerRadius float32

	BottomRightCornerRadius float32

	BottomLeftCornerRadius float32

	Shadow Shadow
}

func (r *Rectangle) Hide() { _ = "STUB: not implemented"; return }

func (r *Rectangle) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (r *Rectangle) Refresh() { _ = "STUB: not implemented"; return }

func (r *Rectangle) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewRectangle(color color.Color) *Rectangle { _ = "STUB: not implemented"; return nil }
