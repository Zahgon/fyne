package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/driver/mobile"
)

type selectable struct {
	BaseWidget
	cursorRow, cursorColumn int

	selectRow, selectColumn int

	focussed, selecting, selectEnded, password bool
	sizeName                                   fyne.ThemeSizeName
	style                                      fyne.TextStyle

	provider *RichText
	theme    fyne.Theme
	focus    fyne.Focusable

	doubleTappedAtUnixMillis int64
}

func (s *selectable) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (s *selectable) Cursor() desktop.Cursor {
	_ = "STUB: not implemented"
	return *new(desktop.Cursor)
}

func (s *selectable) DoubleTapped(p *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) DragEnd() { _ = "STUB: not implemented"; return }

func (s *selectable) Dragged(d *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) dragged(d *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) MouseDown(m *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) MouseUp(ev *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (s *selectable) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) TappedSecondary(ev *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) TouchCancel(m *mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) TouchDown(m *mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) TouchUp(*mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (s *selectable) TypedShortcut(sh fyne.Shortcut) { _ = "STUB: not implemented"; return }

func (s *selectable) cursorColAt(text []rune, pos fyne.Position) int {
	_ = "STUB: not implemented"
	return 0
}

func (s *selectable) getRowCol(p fyne.Position) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *selectable) selectCurrentRow(focus bool) { _ = "STUB: not implemented"; return }

func (s *selectable) selection() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func textPosFromRowCol(row, col int, prov *RichText) int { _ = "STUB: not implemented"; return 0 }

func (s *selectable) updateMousePointer(p fyne.Position) { _ = "STUB: not implemented"; return }

func (s *selectable) getSizeName() fyne.ThemeSizeName {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeSizeName)
}

type selectableRenderer struct {
	sel *selectable

	selections []fyne.CanvasObject
}

func (r *selectableRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *selectableRenderer) Layout(fyne.Size) { _ = "STUB: not implemented"; return }

func (r *selectableRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *selectableRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *selectableRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *selectableRenderer) buildSelection() { _ = "STUB: not implemented"; return }

func (s *selectable) grabFocus() { _ = "STUB: not implemented"; return }

func isTripleTap(double, nowMilli int64) bool { _ = "STUB: not implemented"; return false }
