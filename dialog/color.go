package dialog

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	checkeredBoxSize       = 8
	checkeredNumberOfRings = 12

	preferenceRecents    = "color_recents"
	preferenceMaxRecents = 7
)

type ColorPickerDialog struct {
	*dialog
	Advanced bool
	color    color.Color
	callback func(c color.Color)
	advanced *widget.Accordion
	picker   *colorAdvancedPicker
}

func NewColorPicker(title, message string, callback func(c color.Color), parent fyne.Window) *ColorPickerDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowColorPicker(title, message string, callback func(c color.Color), parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}

func (p *ColorPickerDialog) Refresh() { _ = "STUB: not implemented"; return }

func (p *ColorPickerDialog) SetColor(c color.Color) { _ = "STUB: not implemented"; return }

func (p *ColorPickerDialog) Show() { _ = "STUB: not implemented"; return }

func (p *ColorPickerDialog) createSimplePickers() (contents []fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return nil
}

func (p *ColorPickerDialog) selectColor(c color.Color) { _ = "STUB: not implemented"; return }

func (p *ColorPickerDialog) updateUI() { _ = "STUB: not implemented"; return }

func clamp(value, min, max int) int { _ = "STUB: not implemented"; return 0 }

func wrapHue(hue int) int { _ = "STUB: not implemented"; return 0 }

func newColorButtonBox(colors []color.Color, icon fyne.Resource, callback func(color.Color)) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

//revive:disable-line:add-constant

func newCheckeredBackground(radial bool) *canvas.Raster { _ = "STUB: not implemented"; return nil }

//revive:disable-line:add-constant

//revive:disable-line:add-constant

func readRecentColors() (recents []string) { _ = "STUB: not implemented"; return nil }

func writeRecentColor(color string) { _ = "STUB: not implemented"; return }

func colorToString(c color.Color) string { _ = "STUB: not implemented"; return "" }

func stringsToColors(ss ...string) (colors []color.Color) { _ = "STUB: not implemented"; return nil }

func rgbToHsl(r, g, b uint8) (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

//revive:disable-line:add-constant

func hslToRgb(h, s, l int) (uint8, uint8, uint8) { _ = "STUB: not implemented"; return 0, 0, 0 }

func hueToChannel(h, v1, v2 float64) float64 {
	_ = "STUB: not implemented"
	//revive:disable:add-constant
	return 0
}

//revive:enable-line:add-constant
