package tutorials

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
)

func parseURL(urlStr string) *url.URL { _ = "STUB: not implemented"; return nil }

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func withAlpha(c color.Color, alpha uint8) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

type underLayout struct {
	offset float32
}

func (u underLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (u underLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

type unpad struct {
	top, bottom bool
}

func (u unpad) Layout(objs []fyne.CanvasObject, s fyne.Size) { _ = "STUB: not implemented"; return }

func (u unpad) MinSize(_ []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func formatAuthors(lines string) string { _ = "STUB: not implemented"; return "" }
