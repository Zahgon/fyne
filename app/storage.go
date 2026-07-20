package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal"
)

type store struct {
	*internal.Docs
	a *fyneApp
}

func (s *store) RootURI() fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func (s *store) docRootURI() (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}
