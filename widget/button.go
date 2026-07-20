package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/widget"
)

type ButtonAlign int

type ButtonIconPlacement int

type ButtonImportance = Importance

type ButtonStyle int

const (
	ButtonAlignCenter ButtonAlign = iota

	ButtonAlignLeading

	ButtonAlignTrailing
)

const (
	ButtonIconLeadingText ButtonIconPlacement = iota

	ButtonIconTrailingText
)

var (
	_ fyne.Focusable  = (*Button)(nil)
	_ fyne.Accessible = (*Button)(nil)
)

type Button struct {
	DisableableWidget
	Text string
	Icon fyne.Resource

	Importance    Importance
	Alignment     ButtonAlign
	IconPlacement ButtonIconPlacement

	OnTapped func() `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewButton(label string, tapped func()) *Button { _ = "STUB: not implemented"; return nil }

func NewButtonWithIcon(label string, icon fyne.Resource, tapped func()) *Button {
	_ = "STUB: not implemented"
	return nil
}

func (b *Button) AccessibilityLabel() string { _ = "STUB: not implemented"; return "" }

func (b *Button) AccessibilityRole() fyne.AccessibleRole {
	_ = "STUB: not implemented"
	return *new(fyne.AccessibleRole)
}

func (b *Button) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (b *Button) Cursor() desktop.Cursor { _ = "STUB: not implemented"; return *new(desktop.Cursor) }

func (b *Button) FocusGained() { _ = "STUB: not implemented"; return }

func (b *Button) FocusLost() { _ = "STUB: not implemented"; return }

func (b *Button) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (b *Button) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *Button) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *Button) MouseOut() { _ = "STUB: not implemented"; return }

func (b *Button) SetIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func (b *Button) SetText(text string) { _ = "STUB: not implemented"; return }

func (b *Button) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (b *Button) TypedRune(rune) { _ = "STUB: not implemented"; return }

func (b *Button) TypedKey(ev *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (b *Button) tapAnimation() { _ = "STUB: not implemented"; return }

type buttonRenderer struct {
	widget.BaseRenderer

	icon       *canvas.Image
	label      *RichText
	background *canvas.Rectangle
	tapBG      *canvas.Rectangle
	button     *Button
	layout     fyne.Layout
}

func (r *buttonRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *buttonRenderer) MinSize() (size fyne.Size) {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *buttonRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *buttonRenderer) applyTheme() { _ = "STUB: not implemented"; return }

func (r *buttonRenderer) buttonColorNames() (foreground, background, backgroundBlend fyne.ThemeColorName) {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeColorName), *new(fyne.ThemeColorName), *new(fyne.ThemeColorName)
}

func (r *buttonRenderer) padding(th fyne.Theme) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *buttonRenderer) updateIconAndText() { _ = "STUB: not implemented"; return }

func alignedPosition(align ButtonAlign, padding, objectSize, layoutSize fyne.Size) (pos fyne.Position) {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func blendColor(under, over color.Color) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func newButtonTapAnimation(bg *canvas.Rectangle, w fyne.Widget, th fyne.Theme) *fyne.Animation {
	_ = "STUB: not implemented"
	return nil
}
