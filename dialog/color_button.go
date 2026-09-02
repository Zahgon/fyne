package dialog

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	internalwidget "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

var (
	_ fyne.Widget       = (*colorButton)(nil)
	_ desktop.Hoverable = (*colorButton)(nil)
)

type colorButton struct {
	widget.BaseWidget
	color   color.Color
	onTap   func(color.Color)
	hovered bool
}

func newColorButton(c color.Color, onTap func(color.Color)) *colorButton {
	_ = "STUB: not implemented"
	return nil
}

func (b *colorButton) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (b *colorButton) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *colorButton) MouseOut() { _ = "STUB: not implemented"; return }

func (*colorButton) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *colorButton) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (b *colorButton) SetColor(c color.Color) { _ = "STUB: not implemented"; return }

func (b *colorButton) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

type colorButtonRenderer struct {
	internalwidget.BaseRenderer
	button     *colorButton
	background *canvas.Raster
	rectangle  *canvas.Rectangle
}

func (r *colorButtonRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *colorButtonRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

//revive:disable-line:add-constant

func (r *colorButtonRenderer) Refresh() { _ = "STUB: not implemented"; return }
