package app

import (
	"fyne.io/fyne/v2"
)

var meta = fyne.AppMetadata{
	ID:         "",
	Name:       "",
	Version:    "0.0.1",
	Build:      1,
	Release:    false,
	Custom:     map[string]string{},
	Migrations: map[string]bool{},
}

func SetMetadata(m fyne.AppMetadata) { _ = "STUB: not implemented"; return }

func (a *fyneApp) Metadata() fyne.AppMetadata {
	_ = "STUB: not implemented"
	return *new(fyne.AppMetadata)
}
