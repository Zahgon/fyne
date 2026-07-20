package test

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal"
)

type testStorage struct {
	*internal.Docs
}

func (s *testStorage) RootURI() fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func (s *testStorage) docRootURI() (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}
