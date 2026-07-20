package test

import (
	"errors"
	"io"
	"os"

	"fyne.io/fyne/v2"
)

var errUnsupportedURLProtocol = errors.New("unsupported URL protocol")

type file struct {
	*os.File
	path string
}

type directory struct {
	fyne.URI
}

var _ fyne.ListableURI = (*directory)(nil)

func (f *file) Open() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (f *file) Save() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (f *file) ReadOnly() bool { _ = "STUB: not implemented"; return false }

func (f *file) Name() string { _ = "STUB: not implemented"; return "" }

func (f *file) URI() fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func openFile(uri fyne.URI, create bool) (*file, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *driver) FileReaderForURI(uri fyne.URI) (fyne.URIReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIReadCloser), nil
}

func (d *driver) FileWriterForURI(uri fyne.URI) (fyne.URIWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIWriteCloser), nil
}

func (d *driver) ListerForURI(uri fyne.URI) (fyne.ListableURI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.ListableURI), nil
}

func (d *directory) List() ([]fyne.URI, error) { _ = "STUB: not implemented"; return nil, nil }
