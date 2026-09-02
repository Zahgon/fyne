package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Navigation struct {
	widget.BaseWidget

	Root      fyne.CanvasObject
	Title     string
	OnBack    func() `json:"-"`
	OnForward func() `json:"-"`

	level  int
	stack  fyne.Container
	titles []string
}

func NewNavigation(root fyne.CanvasObject) *Navigation { _ = "STUB: not implemented"; return nil }

func NewNavigationWithTitle(root fyne.CanvasObject, s string) *Navigation {
	_ = "STUB: not implemented"
	return nil
}

func (nav *Navigation) Push(obj fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (nav *Navigation) PushWithTitle(obj fyne.CanvasObject, s string) {
	_ = "STUB: not implemented"
	return
}

func (nav *Navigation) Back() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (nav *Navigation) Forward() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (nav *Navigation) SetTitle(s string) { _ = "STUB: not implemented"; return }

func (nav *Navigation) SetCurrentTitle(s string) { _ = "STUB: not implemented"; return }

func (nav *Navigation) setup() { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*navigatorRenderer)(nil)

type navigatorRenderer struct {
	nav     *Navigation
	back    widget.Button
	forward widget.Button
	title   widget.Label
	objects []fyne.CanvasObject
}

func (nav *Navigation) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (*navigatorRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *navigatorRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *navigatorRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *navigatorRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *navigatorRenderer) Refresh() { _ = "STUB: not implemented"; return }
