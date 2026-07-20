package tutorials

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func confirmCallback(response bool) { _ = "STUB: not implemented"; return }

func colorPicked(c color.Color, w fyne.Window) { _ = "STUB: not implemented"; return }

func dialogScreen(win fyne.Window) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func imageOpened(f fyne.URIReadCloser) { _ = "STUB: not implemented"; return }

func fileSaved(f fyne.URIWriteCloser, w fyne.Window) { _ = "STUB: not implemented"; return }

func loadImage(f fyne.URIReadCloser) *canvas.Image { _ = "STUB: not implemented"; return nil }

func showImage(f fyne.URIReadCloser) { _ = "STUB: not implemented"; return }
