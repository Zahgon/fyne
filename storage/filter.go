package storage

import (
	"fyne.io/fyne/v2"
)

type FileFilter interface {
	Matches(fyne.URI) bool
}

type ExtensionFileFilter struct {
	Extensions []string
}

type MimeTypeFileFilter struct {
	MimeTypes []string
}

func (e *ExtensionFileFilter) Matches(uri fyne.URI) bool { _ = "STUB: not implemented"; return false }

func NewExtensionFileFilter(extensions []string) FileFilter {
	_ = "STUB: not implemented"
	return *new(FileFilter)
}

func (mt *MimeTypeFileFilter) Matches(uri fyne.URI) bool { _ = "STUB: not implemented"; return false }

func NewMimeTypeFileFilter(mimeTypes []string) FileFilter {
	_ = "STUB: not implemented"
	return *new(FileFilter)
}
