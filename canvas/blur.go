package canvas

import "fyne.io/fyne/v2"

var _ fyne.CanvasObject = (*Blur)(nil)

type Blur struct {
	baseObject

	Radius float32

	CornerRadius float32
}

func (b *Blur) Hide() { _ = "STUB: not implemented"; return }

func (b *Blur) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (b *Blur) Refresh() { _ = "STUB: not implemented"; return }

func (b *Blur) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewBlur(radius float32) *Blur { _ = "STUB: not implemented"; return nil }
