package app

import (
	"encoding/base64"
	"io"

	"fyne.io/fyne/v2"
)

type cache struct {
	a fyne.App

	enc base64.Encoding
}

func makeCache(a fyne.App) fyne.Cache { _ = "STUB: not implemented"; return *new(fyne.Cache) }

func (c *cache) RootURI() fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func (c *cache) Exists(name string) bool { _ = "STUB: not implemented"; return false }

func (c *cache) Read(name string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (c *cache) Write(name string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (c *cache) Remove(name string) error { _ = "STUB: not implemented"; return nil }

func (c *cache) encodePath(badName string) fyne.URI {
	_ = "STUB: not implemented"
	return *new(fyne.URI)
}
