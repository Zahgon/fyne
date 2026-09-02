package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/widget"
)

var (
	_ fyne.Widget       = (*radioItem)(nil)
	_ desktop.Hoverable = (*radioItem)(nil)
	_ fyne.Tappable     = (*radioItem)(nil)
	_ fyne.Focusable    = (*radioItem)(nil)
)

func newRadioItem(label string, onTap func(*radioItem)) *radioItem {
	_ = "STUB: not implemented"
	return nil
}

type radioItem struct {
	DisableableWidget

	Label    string
	Selected bool

	focused bool
	hovered bool
	onTap   func(item *radioItem)
}

func (i *radioItem) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (i *radioItem) FocusGained() { _ = "STUB: not implemented"; return }

func (i *radioItem) FocusLost() { _ = "STUB: not implemented"; return }

func (i *radioItem) MouseIn(_ *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (*radioItem) MouseMoved(_ *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (i *radioItem) MouseOut() { _ = "STUB: not implemented"; return }

func (i *radioItem) SetSelected(selected bool) { _ = "STUB: not implemented"; return }

func (i *radioItem) Tapped(_ *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (*radioItem) TypedKey(_ *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (i *radioItem) TypedRune(r rune) { _ = "STUB: not implemented"; return }

func (i *radioItem) toggle() { _ = "STUB: not implemented"; return }

type radioItemRenderer struct {
	widget.BaseRenderer
	item *radioItem

	focusIndicator canvas.Circle
	icon, over     canvas.Image
	label          *canvas.Text
}

func (r *radioItemRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *radioItemRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *radioItemRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *radioItemRenderer) update() { _ = "STUB: not implemented"; return }
