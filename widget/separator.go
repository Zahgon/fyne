package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var _ fyne.Widget = (*Separator)(nil)

type Separator struct {
	BaseWidget

	invert bool
}

func NewSeparator() *Separator { _ = "STUB: not implemented"; return nil }

func (s *Separator) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (s *Separator) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

var _ fyne.WidgetRenderer = (*separatorRenderer)(nil)

type separatorRenderer struct {
	fyne.WidgetRenderer
	bar *canvas.Rectangle
	d   *Separator
}

func (r *separatorRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *separatorRenderer) Refresh() { _ = "STUB: not implemented"; return }
