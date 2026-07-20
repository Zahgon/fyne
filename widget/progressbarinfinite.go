package widget

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
)

const (
	infiniteRefreshRate              = 50 * time.Millisecond
	maxProgressBarInfiniteWidthRatio = 1.0 / 5
	minProgressBarInfiniteWidthRatio = 1.0 / 20
	progressBarInfiniteStepSizeRatio = 1.0 / 50
)

type infProgressRenderer struct {
	widget.BaseRenderer
	background, bar canvas.Rectangle
	animation       fyne.Animation
	wasRunning      bool
	progress        *ProgressBarInfinite
}

func (p *infProgressRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (p *infProgressRenderer) updateBar(done float32) { _ = "STUB: not implemented"; return }

func (p *infProgressRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (p *infProgressRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (p *infProgressRenderer) start() { _ = "STUB: not implemented"; return }

func (p *infProgressRenderer) stop() { _ = "STUB: not implemented"; return }

func (p *infProgressRenderer) Destroy() { _ = "STUB: not implemented"; return }

type ProgressBarInfinite struct {
	BaseWidget
	running bool
}

func (p *ProgressBarInfinite) Show() { _ = "STUB: not implemented"; return }

func (p *ProgressBarInfinite) Hide() { _ = "STUB: not implemented"; return }

func (p *ProgressBarInfinite) Start() { _ = "STUB: not implemented"; return }

func (p *ProgressBarInfinite) Stop() { _ = "STUB: not implemented"; return }

func (p *ProgressBarInfinite) Running() bool { _ = "STUB: not implemented"; return false }

func (p *ProgressBarInfinite) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (p *ProgressBarInfinite) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func NewProgressBarInfinite() *ProgressBarInfinite { _ = "STUB: not implemented"; return nil }
