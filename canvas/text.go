package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Text)(nil)

type Text struct {
	baseObject
	Alignment fyne.TextAlign

	Color     color.Color
	Text      string
	TextSize  float32
	TextStyle fyne.TextStyle

	FontSource fyne.Resource
}

func (t *Text) Hide() { _ = "STUB: not implemented"; return }

func (t *Text) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *Text) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (t *Text) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func (*Text) SetMinSize(fyne.Size) { _ = "STUB: not implemented"; return }

func (t *Text) Refresh() { _ = "STUB: not implemented"; return }

func NewText(text string, c color.Color) *Text { _ = "STUB: not implemented"; return nil }
