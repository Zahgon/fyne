package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
)

const defaultPlaceHolder string = "(Select one)"

var (
	_ fyne.Widget       = (*Select)(nil)
	_ desktop.Hoverable = (*Select)(nil)
	_ fyne.Tappable     = (*Select)(nil)
	_ fyne.Focusable    = (*Select)(nil)
	_ fyne.Disableable  = (*Select)(nil)
)

type Select struct {
	DisableableWidget

	Alignment   fyne.TextAlign
	Selected    string
	Options     []string
	PlaceHolder string
	OnChanged   func(string) `json:"-"`

	binder basicBinder

	focused bool
	hovered bool
	popUp   *PopUpMenu
	tapAnim *fyne.Animation
}

func NewSelect(options []string, changed func(string)) *Select {
	_ = "STUB: not implemented"
	return nil
}

func NewSelectWithData(options []string, data binding.String) *Select {
	_ = "STUB: not implemented"
	return nil
}

func (s *Select) Bind(data binding.String) { _ = "STUB: not implemented"; return }

func (s *Select) ClearSelected() { _ = "STUB: not implemented"; return }

func (s *Select) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (s *Select) FocusGained() { _ = "STUB: not implemented"; return }

func (s *Select) FocusLost() { _ = "STUB: not implemented"; return }

func (s *Select) Hide() { _ = "STUB: not implemented"; return }

func (s *Select) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *Select) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (s *Select) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (s *Select) MouseOut() { _ = "STUB: not implemented"; return }

func (s *Select) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (s *Select) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (s *Select) SelectedIndex() int { _ = "STUB: not implemented"; return 0 }

func (s *Select) SetOptions(options []string) { _ = "STUB: not implemented"; return }

func (s *Select) SetSelected(text string) { _ = "STUB: not implemented"; return }

func (s *Select) SetSelectedIndex(index int) { _ = "STUB: not implemented"; return }

func (s *Select) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (s *Select) TypedKey(event *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (s *Select) TypedRune(_ rune) { _ = "STUB: not implemented"; return }

func (s *Select) Unbind() { _ = "STUB: not implemented"; return }

func (s *Select) popUpPos() fyne.Position { _ = "STUB: not implemented"; return *new(fyne.Position) }

func (s *Select) showPopUp() { _ = "STUB: not implemented"; return }

func (s *Select) tapAnimation() { _ = "STUB: not implemented"; return }

func (s *Select) updateFromData(data binding.DataItem) { _ = "STUB: not implemented"; return }

func (s *Select) updateSelected(text string) { _ = "STUB: not implemented"; return }

func (s *Select) writeData(data binding.DataItem) { _ = "STUB: not implemented"; return }

type selectRenderer struct {
	icon       *Icon
	label      *RichText
	background *canvas.Rectangle

	objects []fyne.CanvasObject
	combo   *Select
}

func (s *selectRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (s *selectRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (s *selectRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (s *selectRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *selectRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (s *selectRenderer) bgColor(th fyne.Theme, v fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (s *selectRenderer) updateIcon(th fyne.Theme) { _ = "STUB: not implemented"; return }

func (s *selectRenderer) updateLabel() { _ = "STUB: not implemented"; return }
