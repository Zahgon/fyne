package test

import (
	"fyne.io/fyne/v2"
)

type device struct{}

var _ fyne.Device = (*device)(nil)

func (*device) Orientation() fyne.DeviceOrientation {
	_ = "STUB: not implemented"
	return *new(fyne.DeviceOrientation)
}

func (*device) HasKeyboard() bool { _ = "STUB: not implemented"; return false }

func (d *device) SystemScale() float32 { _ = "STUB: not implemented"; return 0 }

func (*device) SystemScaleForWindow(fyne.Window) float32 { _ = "STUB: not implemented"; return 0 }

func (*device) Locale() fyne.Locale { _ = "STUB: not implemented"; return *new(fyne.Locale) }

func (*device) IsBrowser() bool { _ = "STUB: not implemented"; return false }
