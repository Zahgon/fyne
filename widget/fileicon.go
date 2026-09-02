package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
)

const (
	ratioDown     = 0.45
	ratioTextSize = 0.22
)

type FileIcon struct {
	BaseWidget

	Selected bool
	URI      fyne.URI

	resource  fyne.Resource
	extension string
}

func NewFileIcon(uri fyne.URI) *FileIcon { _ = "STUB: not implemented"; return nil }

func (i *FileIcon) SetURI(uri fyne.URI) { _ = "STUB: not implemented"; return }

func (i *FileIcon) setURI(uri fyne.URI) { _ = "STUB: not implemented"; return }

func (i *FileIcon) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (i *FileIcon) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (i *FileIcon) SetSelected(selected bool) { _ = "STUB: not implemented"; return }

func (i *FileIcon) lookupIcon(uri fyne.URI) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (i *FileIcon) isDir(uri fyne.URI) bool { _ = "STUB: not implemented"; return false }

type fileIconRenderer struct {
	widget.BaseRenderer

	file *FileIcon

	background *canvas.Rectangle
	ext        *canvas.Text
	img        *canvas.Image
}

func (s *fileIconRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *fileIconRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (s *fileIconRenderer) Refresh() { _ = "STUB: not implemented"; return }

func trimmedExtension(uri fyne.URI) string { _ = "STUB: not implemented"; return "" }
