//go:build wasm || js

package dialog

import (
	"fyne.io/fyne/v2"
)

func (f *fileDialog) loadPlaces() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func isHidden(file fyne.URI) bool { _ = "STUB: not implemented"; return false }

func fileOpenOSOverride(f *FileDialog) bool { _ = "STUB: not implemented"; return false }

func fileSaveOSOverride(f *FileDialog) bool { _ = "STUB: not implemented"; return false }

func (f *fileDialog) getPlaces() []favoriteItem { _ = "STUB: not implemented"; return nil }

func getFavoriteLocation(homeURI fyne.URI, name string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}
