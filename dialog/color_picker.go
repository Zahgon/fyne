package dialog

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func newColorBasicPicker(callback func(color.Color)) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func newColorGreyscalePicker(callback func(color.Color)) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func newColorRecentPicker(callback func(color.Color)) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

var _ fyne.Widget = (*colorAdvancedPicker)(nil)

type colorAdvancedPicker struct {
	widget.BaseWidget
	Red, Green, Blue, Alpha uint8
	Hue                     int
	Saturation, Lightness   int
	ColorModel              string
	previousColor           color.Color

	onChange func(color.Color)
}

func newColorAdvancedPicker(c color.Color, onChange func(color.Color)) *colorAdvancedPicker {
	_ = "STUB: not implemented"
	return nil
}

func (p *colorAdvancedPicker) Color() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (p *colorAdvancedPicker) SetColor(c color.Color) { _ = "STUB: not implemented"; return }

func (p *colorAdvancedPicker) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (p *colorAdvancedPicker) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

//gosec:disable G115 -- r’s value is limited by newColorChannel

//gosec:disable G115 -- g’s value is limited by newColorChannel

//gosec:disable G115 -- b’s value is limited by newColorChannel

//gosec:disable G115 -- a’s value is limited by newColorChannel

func (p *colorAdvancedPicker) setHSLA(h, s, l int, a uint8) { _ = "STUB: not implemented"; return }

func (p *colorAdvancedPicker) setRGBA(r, g, b, a uint8) { _ = "STUB: not implemented"; return }

func (p *colorAdvancedPicker) updateColor(c color.Color) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *colorAdvancedPicker) updateHSLA(h, s, l int, a uint8) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *colorAdvancedPicker) updateRGBA(r, g, b, a uint8) bool {
	_ = "STUB: not implemented"
	return false
}

var _ fyne.WidgetRenderer = (*colorPickerRenderer)(nil)

type colorPickerRenderer struct {
	fyne.WidgetRenderer
	picker            *colorAdvancedPicker
	redChannel        *colorChannel
	greenChannel      *colorChannel
	blueChannel       *colorChannel
	hueChannel        *colorChannel
	saturationChannel *colorChannel
	lightnessChannel  *colorChannel
	wheel             *colorWheel
	preview           *colorPreview
	alphaChannel      *colorChannel
	hex               *userChangeEntry
	contents          fyne.CanvasObject
}

func (r *colorPickerRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *colorPickerRenderer) updateObjects() { _ = "STUB: not implemented"; return }
