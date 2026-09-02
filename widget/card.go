package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
)

type Card struct {
	BaseWidget
	Title, Subtitle string
	Image           *canvas.Image
	Content         fyne.CanvasObject
}

func NewCard(title, subtitle string, content fyne.CanvasObject) *Card {
	_ = "STUB: not implemented"
	return nil
}

func (c *Card) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *Card) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *Card) SetContent(obj fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (c *Card) SetImage(img *canvas.Image) { _ = "STUB: not implemented"; return }

func (c *Card) SetSubTitle(text string) { _ = "STUB: not implemented"; return }

func (c *Card) SetTitle(text string) { _ = "STUB: not implemented"; return }

type cardRenderer struct {
	widget.BaseRenderer

	header, subHeader *canvas.Text

	card *Card

	background *canvas.Rectangle
}

const (
	cardMediaHeight = 128
)

func (c *cardRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (c *cardRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *cardRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (c *cardRenderer) applyTheme() { _ = "STUB: not implemented"; return }
