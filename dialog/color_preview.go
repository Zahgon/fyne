package dialog

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	internalwidget "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

type colorPreview struct {
	widget.BaseWidget

	previous, current color.Color
}

func newColorPreview(previousColor color.Color) *colorPreview {
	_ = "STUB: not implemented"
	return nil
}

func (p *colorPreview) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (p *colorPreview) SetColor(c color.Color) { _ = "STUB: not implemented"; return }

func (p *colorPreview) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

type colorPreviewRenderer struct {
	internalwidget.BaseRenderer
	preview    *colorPreview
	background *canvas.Raster
	old, new   *canvas.Rectangle
}

func (r *colorPreviewRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *colorPreviewRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

//revive:disable-line:add-constant

func (r *colorPreviewRenderer) Refresh() { _ = "STUB: not implemented"; return }
