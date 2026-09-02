package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

var (
	_ fyne.Widget     = (*Label)(nil)
	_ fyne.Accessible = (*Label)(nil)
)

type Label struct {
	BaseWidget
	Text      string
	Alignment fyne.TextAlign
	Wrapping  fyne.TextWrap
	TextStyle fyne.TextStyle

	Truncation fyne.TextTruncation

	Importance Importance

	SizeName fyne.ThemeSizeName

	Selectable bool

	provider  *RichText
	binder    basicBinder
	selection *focusSelectable
}

func NewLabel(text string) *Label { _ = "STUB: not implemented"; return nil }

func NewLabelWithData(data binding.String) *Label { _ = "STUB: not implemented"; return nil }

func NewLabelWithStyle(text string, alignment fyne.TextAlign, style fyne.TextStyle) *Label {
	_ = "STUB: not implemented"
	return nil
}

func (l *Label) AccessibilityLabel() string { _ = "STUB: not implemented"; return "" }

func (*Label) AccessibilityRole() fyne.AccessibleRole {
	_ = "STUB: not implemented"
	return *new(fyne.AccessibleRole)
}

func (l *Label) Bind(data binding.String) { _ = "STUB: not implemented"; return }

func (l *Label) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (l *Label) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *Label) Refresh() { _ = "STUB: not implemented"; return }

func (l *Label) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (l *Label) ClearSelection() { _ = "STUB: not implemented"; return }

func (l *Label) SetText(text string) { _ = "STUB: not implemented"; return }

func (l *Label) Unbind() { _ = "STUB: not implemented"; return }

func (l *Label) syncSegments() { _ = "STUB: not implemented"; return }

func (l *Label) updateFromData(data binding.DataItem) { _ = "STUB: not implemented"; return }

type labelRenderer struct {
	l       *Label
	objects []fyne.CanvasObject
}

func (*labelRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *labelRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *labelRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *labelRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *labelRenderer) Refresh() { _ = "STUB: not implemented"; return }

type focusSelectable struct {
	selectable
}

func (f *focusSelectable) FocusGained() { _ = "STUB: not implemented"; return }

func (f *focusSelectable) FocusLost() { _ = "STUB: not implemented"; return }

func (*focusSelectable) TypedKey(*fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (*focusSelectable) TypedRune(rune) { _ = "STUB: not implemented"; return }
