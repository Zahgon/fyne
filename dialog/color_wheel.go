package dialog

import (
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	internalwidget "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

var (
	_ fyne.Widget    = (*colorWheel)(nil)
	_ fyne.Tappable  = (*colorWheel)(nil)
	_ fyne.Draggable = (*colorWheel)(nil)
)

type colorWheel struct {
	widget.BaseWidget
	generator func(w, h int) image.Image
	cache     draw.Image
	onChange  func(int, int, int, uint8)

	Hue                   int
	Saturation, Lightness int
	Alpha                 uint8
}

func newColorWheel(onChange func(int, int, int, uint8)) *colorWheel {
	_ = "STUB: not implemented"
	return nil
}

func (a *colorWheel) Cursor() desktop.Cursor {
	_ = "STUB: not implemented"
	return *new(desktop.Cursor)
}

func (a *colorWheel) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (a *colorWheel) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (a *colorWheel) SetHSLA(hue, saturation, lightness int, alpha uint8) {
	_ = "STUB: not implemented"
	return
}

func (a *colorWheel) Tapped(event *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (a *colorWheel) Dragged(event *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (a *colorWheel) DragEnd() { _ = "STUB: not implemented"; return }

func (a *colorWheel) colorAt(x, y, w, h int) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (a *colorWheel) locationForPosition(pos fyne.Position) (x, y int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (a *colorWheel) selection(width, height float32) (float32, float32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (a *colorWheel) trigger(pos fyne.Position) { _ = "STUB: not implemented"; return }

type colorWheelRenderer struct {
	internalwidget.BaseRenderer
	area       *colorWheel
	background *canvas.Raster
	raster     *canvas.Raster
	x, y       *canvas.Line
}

func (r *colorWheelRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *colorWheelRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

//revive:disable-line:add-constant

func (r *colorWheelRenderer) Refresh() { _ = "STUB: not implemented"; return }
