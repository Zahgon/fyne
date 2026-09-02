//go:build (linux || openbsd || freebsd || netbsd) && !android && !wasm && !js && !tamago && !noos

package dialog

import (
	"fyne.io/fyne/v2"
)

func getFavoriteLocation(homeURI fyne.URI, name string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}

//gosec:disable G204 - Probably okay to allow arbitrary input for xdg-user-dir. Also, the input is trustworthy as of 2026-06-24.
