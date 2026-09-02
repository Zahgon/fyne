//go:build tamago || noos || tinygo

package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/embedded"
)

func NewWithID(id string) fyne.App { _ = "STUB: not implemented"; return *new(fyne.App) }

func SetDriverDetails(a fyne.App, d embedded.Driver) { _ = "STUB: not implemented"; return }
