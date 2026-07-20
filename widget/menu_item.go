package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/widget"
)

var (
	_ fyne.Widget       = (*menuItem)(nil)
	_ desktop.Hoverable = (*menuItem)(nil)
	_ fyne.Tappable     = (*menuItem)(nil)
)

type menuItem struct {
	widget.Base
	Item *fyne.MenuItem

	alignment     fyne.TextAlign
	child, parent *Menu
}

func newMenuItem(item *fyne.MenuItem, parent *Menu) *menuItem {
	_ = "STUB: not implemented"
	return nil
}

func (i *menuItem) Child() *Menu { _ = "STUB: not implemented"; return nil }

func (i *menuItem) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (i *menuItem) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (i *menuItem) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (i *menuItem) MouseOut() { _ = "STUB: not implemented"; return }

func (i *menuItem) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (i *menuItem) activate() { _ = "STUB: not implemented"; return }

func (i *menuItem) activateLastSubmenu() bool { _ = "STUB: not implemented"; return false }

func (i *menuItem) deactivate() { _ = "STUB: not implemented"; return }

func (i *menuItem) deactivateLastSubmenu() bool { _ = "STUB: not implemented"; return false }

func (i *menuItem) isActive() bool { _ = "STUB: not implemented"; return false }

func (i *menuItem) isSubmenuOpen() bool { _ = "STUB: not implemented"; return false }

func (i *menuItem) trigger() { _ = "STUB: not implemented"; return }

func (i *menuItem) triggerLast() { _ = "STUB: not implemented"; return }

type menuItemRenderer struct {
	widget.BaseRenderer
	i                *menuItem
	background       *canvas.Rectangle
	checkIcon        *canvas.Image
	expandIcon       *canvas.Image
	icon             *canvas.Image
	lastThemePadding float32
	minSize          fyne.Size
	shortcutTexts    []*canvas.Text
	text             *canvas.Text
}

func (r *menuItemRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *menuItemRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *menuItemRenderer) updateVisuals() { _ = "STUB: not implemented"; return }

func (r *menuItemRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *menuItemRenderer) checkSpace() float32 { _ = "STUB: not implemented"; return 0 }

func (r *menuItemRenderer) minSizeUnchanged() bool { _ = "STUB: not implemented"; return false }

func (r *menuItemRenderer) updateIcon(img *canvas.Image, rsc fyne.Resource) {
	_ = "STUB: not implemented"
	return
}

func (r *menuItemRenderer) refreshText(text *canvas.Text, shortcut bool) {
	_ = "STUB: not implemented"
	return
}

func shortcutColor(th fyne.Theme) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func textsForShortcut(sc fyne.KeyboardShortcut, th fyne.Theme) (texts []*canvas.Text) {
	_ = "STUB: not implemented"
	return nil
}
