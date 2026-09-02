package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
)

var (
	_ fyne.Widget   = (*Menu)(nil)
	_ fyne.Tappable = (*Menu)(nil)
)

type Menu struct {
	BaseWidget
	alignment     fyne.TextAlign
	Items         []fyne.CanvasObject
	OnDismiss     func() `json:"-"`
	activeItem    *menuItem
	customSized   bool
	containsCheck bool
}

func NewMenu(menu *fyne.Menu) *Menu { _ = "STUB: not implemented"; return nil }

func (m *Menu) ActivateLastSubmenu() bool { _ = "STUB: not implemented"; return false }

func (m *Menu) ActivateNext() { _ = "STUB: not implemented"; return }

func (m *Menu) ActivatePrevious() { _ = "STUB: not implemented"; return }

func (m *Menu) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (m *Menu) DeactivateChild() { _ = "STUB: not implemented"; return }

func (m *Menu) DeactivateLastSubmenu() bool { _ = "STUB: not implemented"; return false }

func (m *Menu) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (m *Menu) Refresh() { _ = "STUB: not implemented"; return }

func (m *Menu) getContainsCheck() bool { _ = "STUB: not implemented"; return false }

func (*Menu) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (m *Menu) TriggerLast() { _ = "STUB: not implemented"; return }

func (m *Menu) Dismiss() { _ = "STUB: not implemented"; return }

func (m *Menu) activateItem(item *menuItem) { _ = "STUB: not implemented"; return }

func (m *Menu) setMenu(menu *fyne.Menu) { _ = "STUB: not implemented"; return }

type menuRenderer struct {
	widget.BaseRenderer
	box    *menuBox
	m      *Menu
	scroll *widget.Scroll
	b      *canvas.Rectangle
}

func (r *menuRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *menuRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *menuRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *menuRenderer) layoutActiveChild() { _ = "STUB: not implemented"; return }

type menuBox struct {
	BaseWidget
	items []fyne.CanvasObject
}

var _ fyne.Widget = (*menuBox)(nil)

func newMenuBox(items []fyne.CanvasObject) *menuBox { _ = "STUB: not implemented"; return nil }

func (b *menuBox) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

type menuBoxRenderer struct {
	widget.BaseRenderer
	b    *menuBox
	cont *fyne.Container
}

var _ fyne.WidgetRenderer = (*menuBoxRenderer)(nil)

func (r *menuBoxRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *menuBoxRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *menuBoxRenderer) Refresh() { _ = "STUB: not implemented"; return }
