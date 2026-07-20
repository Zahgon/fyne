package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	_ desktop.Cursorable = (*passwordRevealer)(nil)
	_ fyne.Tappable      = (*passwordRevealer)(nil)
	_ fyne.Widget        = (*passwordRevealer)(nil)
)

type passwordRevealer struct {
	BaseWidget

	icon  *canvas.Image
	entry *Entry
}

func newPasswordRevealer(e *Entry) *passwordRevealer { _ = "STUB: not implemented"; return nil }

func (r *passwordRevealer) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (r *passwordRevealer) Cursor() desktop.Cursor {
	_ = "STUB: not implemented"
	return *new(desktop.Cursor)
}

func (r *passwordRevealer) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*passwordRevealerRenderer)(nil)

type passwordRevealerRenderer struct {
	fyne.WidgetRenderer
	entry *Entry
	icon  *canvas.Image
}

func (r *passwordRevealerRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *passwordRevealerRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *passwordRevealerRenderer) Refresh() { _ = "STUB: not implemented"; return }
