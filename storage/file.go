package storage

import (
	"fyne.io/fyne/v2"
)

func OpenFileFromURI(uri fyne.URI) (fyne.URIReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIReadCloser), nil
}

func SaveFileToURI(uri fyne.URI) (fyne.URIWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIWriteCloser), nil
}

func ListerForURI(uri fyne.URI) (fyne.ListableURI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.ListableURI), nil
}

type legacyListable struct {
	fyne.URI
}

func (l *legacyListable) List() ([]fyne.URI, error) { _ = "STUB: not implemented"; return nil, nil }
