package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var _ fyne.Widget = (*Activity)(nil)

type Activity struct {
	BaseWidget

	started bool
}

func NewActivity() *Activity { _ = "STUB: not implemented"; return nil }

func (a *Activity) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (a *Activity) Start() { _ = "STUB: not implemented"; return }

func (a *Activity) Stop() { _ = "STUB: not implemented"; return }

func (a *Activity) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

var _ fyne.WidgetRenderer = (*activityRenderer)(nil)

type activityRenderer struct {
	anim   *fyne.Animation
	dots   []fyne.CanvasObject
	parent *Activity

	bound      fyne.Size
	maxCol     color.NRGBA
	maxRad     float32
	wasStarted bool
}

func (a *activityRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (a *activityRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (a *activityRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (a *activityRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (a *activityRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (a *activityRenderer) animate(done float32) { _ = "STUB: not implemented"; return }

func (a *activityRenderer) scaleDot(dot *canvas.Circle, off float32) {
	_ = "STUB: not implemented"
	return
}

//revive:disable-line:add-constant

func (a *activityRenderer) start() { _ = "STUB: not implemented"; return }

func (a *activityRenderer) stop() { _ = "STUB: not implemented"; return }

func (a *activityRenderer) drawStaticEllipsis() { _ = "STUB: not implemented"; return }

func (a *activityRenderer) hideDots() { _ = "STUB: not implemented"; return }

func (a *activityRenderer) updateColor() { _ = "STUB: not implemented"; return }
