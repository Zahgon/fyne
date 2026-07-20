package test

import (
	"fyne.io/fyne/v2"
	fynedriver "fyne.io/fyne/v2/driver"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/driver/software"
)

type WindowlessCanvas = software.WindowlessCanvas

var dummyCanvas software.WindowlessCanvas

func Canvas() fyne.Canvas { _ = "STUB: not implemented"; return *new(fyne.Canvas) }

func NewCanvas() software.WindowlessCanvas {
	_ = "STUB: not implemented"
	return *new(software.WindowlessCanvas)
}

func NewCanvasWithPainter(painter fynedriver.Painter) software.WindowlessCanvas {
	_ = "STUB: not implemented"
	return *new(software.WindowlessCanvas)
}

func NewTransparentCanvasWithPainter(painter fynedriver.Painter) software.WindowlessCanvas {
	_ = "STUB: not implemented"
	return *new(software.WindowlessCanvas)
}

func wrapCanvas(c software.WindowlessCanvas) *canvas { _ = "STUB: not implemented"; return nil }

type canvas struct {
	software.WindowlessCanvas
	hovered desktop.Hoverable
}
