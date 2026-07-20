//go:build !tamago && !noos

package test

import (
	"image"
	"testing"

	"fyne.io/fyne/v2"
)

func AssertCanvasTappableAt(t *testing.T, c fyne.Canvas, pos fyne.Position) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertObjectRendersToImage(t *testing.T, masterFilename string, o fyne.CanvasObject, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertObjectRendersToMarkup(t *testing.T, masterFilename string, o fyne.CanvasObject, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertImageMatches(t *testing.T, masterFilename string, img image.Image, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertRendersToImage(t *testing.T, masterFilename string, c fyne.Canvas, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertRendersToMarkup(t *testing.T, masterFilename string, c fyne.Canvas, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func ApplyTheme(t *testing.T, theme fyne.Theme) { _ = "STUB: not implemented"; return }

func TempWidgetRenderer(t *testing.T, wid fyne.Widget) fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func WithTestTheme(t *testing.T, f func()) { _ = "STUB: not implemented"; return }
