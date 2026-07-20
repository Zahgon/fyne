package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/internal/widget"
)

type progressRenderer struct {
	widget.BaseRenderer
	background, bar canvas.Rectangle
	label           canvas.Text
	ratio           float32
	progress        *ProgressBar
}

func (p *progressRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (p *progressRenderer) calculateRatio() { _ = "STUB: not implemented"; return }

func (p *progressRenderer) updateBar() { _ = "STUB: not implemented"; return }

func (p *progressRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (p *progressRenderer) applyTheme() { _ = "STUB: not implemented"; return }

func (p *progressRenderer) Refresh() { _ = "STUB: not implemented"; return }

type ProgressBar struct {
	BaseWidget

	Min, Max, Value float64

	TextFormatter func() string `json:"-"`

	binder basicBinder
}

func (p *ProgressBar) Bind(data binding.Float) { _ = "STUB: not implemented"; return }

func (p *ProgressBar) SetValue(v float64) { _ = "STUB: not implemented"; return }

func (p *ProgressBar) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (p *ProgressBar) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (p *ProgressBar) Unbind() { _ = "STUB: not implemented"; return }

func NewProgressBar() *ProgressBar { _ = "STUB: not implemented"; return nil }

func NewProgressBarWithData(data binding.Float) *ProgressBar { _ = "STUB: not implemented"; return nil }

func progressBlendColor(clr color.Color) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (p *ProgressBar) updateFromData(data binding.DataItem) { _ = "STUB: not implemented"; return }
