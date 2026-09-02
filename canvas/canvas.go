package canvas

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
)

const (
	RadiusMaximum float32 = math.MaxFloat32
)

func Refresh(obj fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func RecolorSVG(svgContent []byte, c color.Color) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func repaint(obj fyne.CanvasObject) { _ = "STUB: not implemented"; return }
