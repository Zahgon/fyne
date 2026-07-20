package fyne

import "io"

type Cache interface {
	RootURI() URI

	Exists(name string) bool
	Read(name string) (io.ReadCloser, error)
	Write(name string) (io.WriteCloser, error)
	Remove(name string) error
}
