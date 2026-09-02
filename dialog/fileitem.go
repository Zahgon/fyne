package dialog

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	cellWidthIcon = iconSize * 1.25

	folderInsetBottom       = 25
	folderInsetBottomInline = 6
	folderInsetTop          = 20
	folderInsetTopInline    = 6
	folderInsetX            = 10
	folderInsetXInline      = 4

	iconSize       = 64
	iconSizeInline = 24
)

type fileDialogItem struct {
	widget.BaseWidget
	picker *fileDialog

	name     string
	id       int
	choose   func(id int)
	open     func()
	location fyne.URI
	dir      bool

	lastClick time.Time
}

func (i *fileDialogItem) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (i *fileDialogItem) setLocation(l fyne.URI, dir, up bool) { _ = "STUB: not implemented"; return }

func (i *fileDialogItem) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (f *fileDialog) newFileItem(location fyne.URI, dir, up bool) *fileDialogItem {
	_ = "STUB: not implemented"
	return nil
}

type fileItemRenderer struct {
	item         *fileDialogItem
	fileTextSize float32

	icon    *widget.FileIcon
	text    *widget.Label
	over    *canvas.Image
	objects []fyne.CanvasObject
}

func (s *fileItemRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (s *fileItemRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *fileItemRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (s *fileItemRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (*fileItemRenderer) Destroy() { _ = "STUB: not implemented"; return }
