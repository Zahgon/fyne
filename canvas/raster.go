package canvas

import (
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Raster)(nil)

type Raster struct {
	baseObject

	Generator func(w, h int) image.Image

	Translucency float64

	ScaleMode ImageScale
}

func (r *Raster) Alpha() float64 { _ = "STUB: not implemented"; return 0 }

func (r *Raster) Hide() { _ = "STUB: not implemented"; return }

func (r *Raster) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (r *Raster) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *Raster) Refresh() { _ = "STUB: not implemented"; return }

func NewRaster(generate func(w, h int) image.Image) *Raster { _ = "STUB: not implemented"; return nil }

type pixelRaster struct {
	r *Raster

	img draw.Image
}

func NewRasterWithPixels(pixelColor func(x, y, w, h int) color.Color) *Raster {
	_ = "STUB: not implemented"
	return nil
}

func NewRasterFromImage(img image.Image) *Raster { _ = "STUB: not implemented"; return nil }
