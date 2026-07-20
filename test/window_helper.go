//go:build !tamago && !noos

package test

import (
	"testing"

	"fyne.io/fyne/v2"
)

func NewTempWindow(t testing.TB, content fyne.CanvasObject) fyne.Window {
	_ = "STUB: not implemented"
	return *new(fyne.Window)
}
