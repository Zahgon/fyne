package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*BezierCurve)(nil)

type BezierCurve struct {
	baseObject

	StartPoint    fyne.Position
	EndPoint      fyne.Position
	ControlPoints []fyne.Position
	StrokeColor   color.Color
	StrokeWidth   float32
}

func (r *BezierCurve) Hide() { _ = "STUB: not implemented"; return }

func (r *BezierCurve) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (r *BezierCurve) Refresh() { _ = "STUB: not implemented"; return }

func (r *BezierCurve) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewLinearBezierCurve(startPoint, endPoint fyne.Position, c color.Color) *BezierCurve {
	_ = "STUB: not implemented"
	return nil
}

func NewQuadraticBezierCurve(startPoint, controlPoint, endPoint fyne.Position, c color.Color) *BezierCurve {
	_ = "STUB: not implemented"
	return nil
}

func NewCubicBezierCurve(startPoint, controlPoint1, controlPoint2, endPoint fyne.Position, c color.Color) *BezierCurve {
	_ = "STUB: not implemented"
	return nil
}
