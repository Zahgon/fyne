package canvas

import (
	"fyne.io/fyne/v2"
)

type baseObject struct {
	size     fyne.Size
	position fyne.Position
	Hidden   bool

	min fyne.Size
}

func (o *baseObject) Hide() { _ = "STUB: not implemented"; return }

func (o *baseObject) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (o *baseObject) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (o *baseObject) Position() fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func (o *baseObject) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (o *baseObject) SetMinSize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (o *baseObject) Show() { _ = "STUB: not implemented"; return }

func (o *baseObject) Size() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (o *baseObject) Visible() bool { _ = "STUB: not implemented"; return false }
