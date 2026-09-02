package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type scaleItems struct {
	scale   float32
	name    string
	preview *canvas.Text
	button  *widget.Button
}

var scales = [...]*scaleItems{
	{scale: 0.5, name: "Tiny"},
	{scale: 0.8, name: "Small"},
	{scale: 1, name: "Normal"},
	{scale: 1.3, name: "Large"},
	{scale: 1.8, name: "Huge"},
}

func (*Settings) appliedScale(value float32) { _ = "STUB: not implemented"; return }

func (s *Settings) chooseScale(value float32) { _ = "STUB: not implemented"; return }

func (s *Settings) makeScaleButtons() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (s *Settings) makeScaleGroup(scale float32) *widget.Card {
	_ = "STUB: not implemented"
	return nil
}

func (*Settings) makeScalePreviews(value float32) []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func (*Settings) refreshScalePreviews() { _ = "STUB: not implemented"; return }

type refreshMonitor struct {
	widget.Label
	settings *Settings
}

func (r *refreshMonitor) Refresh() { _ = "STUB: not implemented"; return }

func newRefreshMonitor(s *Settings) *refreshMonitor { _ = "STUB: not implemented"; return nil }
