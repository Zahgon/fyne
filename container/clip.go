package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*Clip)(nil)

type Clip struct {
	widget.BaseWidget
	Content fyne.CanvasObject
}

func NewClip(content fyne.CanvasObject) *Clip { _ = "STUB: not implemented"; return nil }

func (c *Clip) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *Clip) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

type clipRenderer struct {
	c       *Clip
	objects []fyne.CanvasObject
}

func newClipRenderer(c *Clip) *clipRenderer { _ = "STUB: not implemented"; return nil }

func (*clipRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *clipRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *clipRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *clipRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *clipRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (*clipRenderer) IsClip() { _ = "STUB: not implemented"; return }
