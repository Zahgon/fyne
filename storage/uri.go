package storage

import (
	"fyne.io/fyne/v2"
)

func EqualURI(t1, t2 fyne.URI) bool { _ = "STUB: not implemented"; return false }

func NewFileURI(fpath string) fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func NewURI(s string) fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func ParseURI(s string) (fyne.URI, error) { _ = "STUB: not implemented"; return *new(fyne.URI), nil }

func Parent(u fyne.URI) (fyne.URI, error) { _ = "STUB: not implemented"; return *new(fyne.URI), nil }

func Child(u fyne.URI, component string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}

func Exists(u fyne.URI) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func Delete(u fyne.URI) error { _ = "STUB: not implemented"; return nil }

func DeleteAll(u fyne.URI) error { _ = "STUB: not implemented"; return nil }

func Reader(u fyne.URI) (fyne.URIReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIReadCloser), nil
}

func CanRead(u fyne.URI) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func Writer(u fyne.URI) (fyne.URIWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIWriteCloser), nil
}

func Appender(u fyne.URI) (fyne.URIWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIWriteCloser), nil
}

func CanWrite(u fyne.URI) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func Copy(source fyne.URI, destination fyne.URI) error { _ = "STUB: not implemented"; return nil }

func Move(source fyne.URI, destination fyne.URI) error { _ = "STUB: not implemented"; return nil }

func CanList(u fyne.URI) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func List(u fyne.URI) ([]fyne.URI, error) { _ = "STUB: not implemented"; return nil, nil }

func CreateListable(u fyne.URI) error { _ = "STUB: not implemented"; return nil }
