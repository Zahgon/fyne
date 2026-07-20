//go:build windows && mobile

package app

import (
	"fyne.io/fyne/v2"
)

func NewWithID(_ string) fyne.App { _ = "STUB: not implemented"; return *new(fyne.App) }
