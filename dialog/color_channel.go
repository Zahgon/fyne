package dialog

import (
	"fyne.io/fyne/v2"
	internalwidget "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*colorChannel)(nil)

type colorChannel struct {
	widget.BaseWidget
	name      string
	min, max  int
	value     int
	onChanged func(int)
}

func newColorChannel(name string, min, max, value int, onChanged func(int)) *colorChannel {
	_ = "STUB: not implemented"
	return nil
}

func (c *colorChannel) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *colorChannel) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *colorChannel) SetValue(value int) { _ = "STUB: not implemented"; return }

type colorChannelRenderer struct {
	internalwidget.BaseRenderer
	control *colorChannel
	label   *widget.Label
	entry   *colorChannelEntry
	slider  *widget.Slider
}

func (r *colorChannelRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *colorChannelRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *colorChannelRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *colorChannelRenderer) updateObjects() { _ = "STUB: not implemented"; return }

type colorChannelEntry struct {
	userChangeEntry
}

func newColorChannelEntry(c *colorChannel) *colorChannelEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *colorChannelEntry) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

type userChangeEntry struct {
	widget.Entry
	userTyped bool
}

func newUserChangeEntry(text string) *userChangeEntry { _ = "STUB: not implemented"; return nil }

func (e *userChangeEntry) setOnChanged(onChanged func(s string)) { _ = "STUB: not implemented"; return }

func (e *userChangeEntry) TypedRune(r rune) { _ = "STUB: not implemented"; return }

func (e *userChangeEntry) TypedKey(ev *fyne.KeyEvent) { _ = "STUB: not implemented"; return }
