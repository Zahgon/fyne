//go:build !windows && !android && !ios && !wasm && !js

package dialog

import (
	"fyne.io/fyne/v2"
)

func (*fileDialog) getPlaces() []favoriteItem { _ = "STUB: not implemented"; return nil }

func isHidden(file fyne.URI) bool { _ = "STUB: not implemented"; return false }

func hideFile(_ string) error { _ = "STUB: not implemented"; return nil }
