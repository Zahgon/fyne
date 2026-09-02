package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/widget"
)

type Check struct {
	DisableableWidget
	Text    string
	Checked bool

	Partial bool

	OnChanged func(bool) `json:"-"`

	focused bool
	hovered bool

	binder basicBinder

	minSize fyne.Size
}

func NewCheck(label string, changed func(bool)) *Check { _ = "STUB: not implemented"; return nil }

func NewCheckWithData(label string, data binding.Bool) *Check {
	_ = "STUB: not implemented"
	return nil
}

func (c *Check) Bind(data binding.Bool) { _ = "STUB: not implemented"; return }

func (c *Check) SetChecked(checked bool) { _ = "STUB: not implemented"; return }

func (c *Check) Hide() { _ = "STUB: not implemented"; return }

func (c *Check) MouseIn(me *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (c *Check) MouseOut() { _ = "STUB: not implemented"; return }

func (c *Check) MouseMoved(me *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (c *Check) Tapped(pe *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (c *Check) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *Check) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *Check) FocusGained() { _ = "STUB: not implemented"; return }

func (c *Check) FocusLost() { _ = "STUB: not implemented"; return }

func (c *Check) TypedRune(r rune) { _ = "STUB: not implemented"; return }

func (*Check) TypedKey(*fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (c *Check) SetText(text string) { _ = "STUB: not implemented"; return }

func (c *Check) Unbind() { _ = "STUB: not implemented"; return }

func (c *Check) updateFromData(data binding.DataItem) { _ = "STUB: not implemented"; return }

func (c *Check) writeData(data binding.DataItem) { _ = "STUB: not implemented"; return }

type checkRenderer struct {
	widget.BaseRenderer
	bg, icon       *canvas.Image
	label          *canvas.Text
	focusIndicator *canvas.Circle
	check          *Check
}

func (c *checkRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *checkRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (c *checkRenderer) applyTheme(th fyne.Theme, v fyne.ThemeVariant) {
	_ = "STUB: not implemented"
	return
}

func (c *checkRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (c *checkRenderer) updateLabel() { _ = "STUB: not implemented"; return }

func (c *checkRenderer) updateResource(th fyne.Theme) { _ = "STUB: not implemented"; return }

func (c *checkRenderer) updateFocusIndicator(th fyne.Theme, v fyne.ThemeVariant) {
	_ = "STUB: not implemented"
	return
}

func focusIfNotMobile(w fyne.Widget) { _ = "STUB: not implemented"; return }
