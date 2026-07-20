package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
)

type iconRenderer struct {
	widget.BaseRenderer
	raster *canvas.Image

	image *Icon
}

func (i *iconRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (i *iconRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (i *iconRenderer) Refresh() { _ = "STUB: not implemented"; return }

type Icon struct {
	BaseWidget

	Resource  fyne.Resource
	cachedRes fyne.Resource
}

func (i *Icon) SetResource(res fyne.Resource) { _ = "STUB: not implemented"; return }

func (i *Icon) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (i *Icon) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func NewIcon(res fyne.Resource) *Icon { _ = "STUB: not implemented"; return nil }
