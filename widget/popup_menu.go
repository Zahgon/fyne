package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/widget"
)

var (
	_ fyne.Widget    = (*PopUpMenu)(nil)
	_ fyne.Focusable = (*PopUpMenu)(nil)
)

type PopUpMenu struct {
	*Menu
	canvas  fyne.Canvas
	overlay *widget.OverlayContainer
}

func NewPopUpMenu(menu *fyne.Menu, c fyne.Canvas) *PopUpMenu { _ = "STUB: not implemented"; return nil }

//revive:disable-line:add-constant // non-zero pos to get manual overlay, fixed on show

func ShowPopUpMenuAtPosition(menu *fyne.Menu, c fyne.Canvas, pos fyne.Position) {
	_ = "STUB: not implemented"
	return
}

func ShowPopUpMenuAtRelativePosition(menu *fyne.Menu, c fyne.Canvas, rel fyne.Position, to fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return
}

func (p *PopUpMenu) FocusGained() { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) FocusLost() { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) Hide() { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) SetCanvas(c fyne.Canvas) { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) Show() { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) ShowAtPosition(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) ShowAtRelativePosition(rel fyne.Position, to fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return
}

func (p *PopUpMenu) TypedKey(e *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) TypedRune(rune) { _ = "STUB: not implemented"; return }

func (p *PopUpMenu) adjustedPosition(pos fyne.Position, size fyne.Size) fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}
