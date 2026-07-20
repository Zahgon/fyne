package widget

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	_ fyne.Accessible = (*Hyperlink)(nil)
	_ fyne.Focusable  = (*Hyperlink)(nil)
	_ fyne.Widget     = (*Hyperlink)(nil)
)

type Hyperlink struct {
	BaseWidget
	Text      string
	URL       *url.URL
	Alignment fyne.TextAlign
	Wrapping  fyne.TextWrap
	TextStyle fyne.TextStyle

	Truncation fyne.TextTruncation

	SizeName fyne.ThemeSizeName

	OnTapped func() `json:"-"`

	textSize         fyne.Size
	focused, hovered bool
	provider         RichText

	siblings []*Hyperlink
}

func NewHyperlink(text string, url *url.URL) *Hyperlink { _ = "STUB: not implemented"; return nil }

func NewHyperlinkWithStyle(text string, url *url.URL, alignment fyne.TextAlign, style fyne.TextStyle) *Hyperlink {
	_ = "STUB: not implemented"
	return nil
}

func (hl *Hyperlink) AccessibilityLabel() string { _ = "STUB: not implemented"; return "" }

func (hl *Hyperlink) AccessibilityRole() fyne.AccessibleRole {
	_ = "STUB: not implemented"
	return *new(fyne.AccessibleRole)
}

func (hl *Hyperlink) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (hl *Hyperlink) Cursor() desktop.Cursor {
	_ = "STUB: not implemented"
	return *new(desktop.Cursor)
}

func (hl *Hyperlink) FocusGained() { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) FocusLost() { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) MouseIn(e *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) MouseMoved(e *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) MouseOut() { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) setHovered(hovered bool) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) focusWidth() float32 { _ = "STUB: not implemented"; return 0 }

func (hl *Hyperlink) focusXPos() float32 { _ = "STUB: not implemented"; return 0 }

func (hl *Hyperlink) isPosOverText(pos fyne.Position) bool { _ = "STUB: not implemented"; return false }

func (hl *Hyperlink) Refresh() { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (hl *Hyperlink) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) SetText(text string) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) SetURL(url *url.URL) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) SetURLFromString(str string) error { _ = "STUB: not implemented"; return nil }

func (hl *Hyperlink) Tapped(e *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) invokeAction() { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) TypedRune(rune) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) TypedKey(ev *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) openURL() { _ = "STUB: not implemented"; return }

func (hl *Hyperlink) syncSegments() { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*hyperlinkRenderer)(nil)

type hyperlinkRenderer struct {
	hl    *Hyperlink
	focus *canvas.Rectangle
	under *canvas.Rectangle

	objects []fyne.CanvasObject
}

func (r *hyperlinkRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *hyperlinkRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *hyperlinkRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *hyperlinkRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *hyperlinkRenderer) Refresh() { _ = "STUB: not implemented"; return }
