package test

import (
	"io"

	"fyne.io/fyne/v2"
)

type memCache struct {
	memStore map[string][]byte
}

func makeCache() fyne.Cache { _ = "STUB: not implemented"; return *new(fyne.Cache) }

func (*memCache) RootURI() fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func (c *memCache) Exists(name string) bool { _ = "STUB: not implemented"; return false }

func (c *memCache) Read(name string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (c *memCache) Write(name string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (c *memCache) Remove(name string) error { _ = "STUB: not implemented"; return nil }

type writeCloser struct {
	io.Writer

	onClose func()
}

func (c writeCloser) Close() error { _ = "STUB: not implemented"; return nil }
