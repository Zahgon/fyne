//go:build !ci && (android || ios || mobile) && !windows

package app

import (
	"fyne.io/fyne/v2"
)

func NewWithID(id string) fyne.App { _ = "STUB: not implemented"; return *new(fyne.App) }

func (a *fyneApp) registerRepositories() { _ = "STUB: not implemented"; return }
