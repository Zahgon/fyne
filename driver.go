package fyne

import "time"

type Driver interface {
	CreateWindow(string) Window

	AllWindows() []Window

	RenderedTextSize(text string, fontSize float32, style TextStyle, source Resource) (size Size, baseline float32)

	CanvasForObject(CanvasObject) Canvas

	AbsolutePositionForObject(CanvasObject) Position

	Device() Device

	Run()

	Quit()

	StartAnimation(*Animation)

	StopAnimation(*Animation)

	DoubleTapDelay() time.Duration

	SetDisableScreenBlanking(bool)

	DoFromGoroutine(fn func(), wait bool)
}
