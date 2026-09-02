package test

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
	fynedriver "fyne.io/fyne/v2/driver"
)

type SoftwarePainter = fynedriver.Painter

type driver struct {
	device       device
	painter      fynedriver.Painter
	windows      []fyne.Window
	windowsMutex sync.RWMutex
}

var _ fyne.Driver = (*driver)(nil)

func NewDriver() fyne.Driver { _ = "STUB: not implemented"; return *new(fyne.Driver) }

func NewDriverWithPainter(p fynedriver.Painter) fyne.Driver {
	_ = "STUB: not implemented"
	return *new(fyne.Driver)
}

func (*driver) DoFromGoroutine(f func(), _ bool) { _ = "STUB: not implemented"; return }

func (d *driver) AbsolutePositionForObject(co fyne.CanvasObject) fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func (d *driver) AllWindows() []fyne.Window { _ = "STUB: not implemented"; return nil }

func (d *driver) CanvasForObject(fyne.CanvasObject) fyne.Canvas {
	_ = "STUB: not implemented"
	return *new(fyne.Canvas)
}

func (d *driver) CreateWindow(title string) fyne.Window {
	_ = "STUB: not implemented"
	return *new(fyne.Window)
}

func (d *driver) Device() fyne.Device { _ = "STUB: not implemented"; return *new(fyne.Device) }

func (*driver) RenderedTextSize(text string, size float32, style fyne.TextStyle, source fyne.Resource) (fyne.Size, float32) {
	_ = "STUB: not implemented"
	return *new(fyne.Size), 0
}

func (*driver) Run() { _ = "STUB: not implemented"; return }

func (*driver) StartAnimation(a *fyne.Animation) { _ = "STUB: not implemented"; return }

func (*driver) StopAnimation(*fyne.Animation) { _ = "STUB: not implemented"; return }

func (*driver) Quit() { _ = "STUB: not implemented"; return }

func (*driver) Clipboard() fyne.Clipboard { _ = "STUB: not implemented"; return *new(fyne.Clipboard) }

func (d *driver) removeWindow(w *window) { _ = "STUB: not implemented"; return }

func (*driver) DoubleTapDelay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (*driver) SetDisableScreenBlanking(_ bool) { _ = "STUB: not implemented"; return }
