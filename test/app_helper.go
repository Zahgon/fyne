//go:build !tamago && !noos

package test

import (
	"testing"

	"fyne.io/fyne/v2"
)

func NewTempApp(t testing.TB) fyne.App { _ = "STUB: not implemented"; return *new(fyne.App) }
