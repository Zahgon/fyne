package tutorials

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type iconInfo struct {
	name string
	icon fyne.Resource
}

type browser struct {
	current int
	icons   []iconInfo

	name *widget.Select
	icon *widget.Icon
}

func (b *browser) setIcon(index int) { _ = "STUB: not implemented"; return }

func iconScreen(_ fyne.Window) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func checkerPattern(x, y, _, _ int) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func iconList(icons []iconInfo) []string { _ = "STUB: not implemented"; return nil }

func loadIcons() []iconInfo { _ = "STUB: not implemented"; return nil }
