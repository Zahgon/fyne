//go:build ios || android

package dialog

import (
	"fyne.io/fyne/v2"
)

func (f *fileDialog) getPlaces() []favoriteItem { _ = "STUB: not implemented"; return nil }

func isHidden(file fyne.URI) bool { _ = "STUB: not implemented"; return false }

func hideFile(filename string) error { _ = "STUB: not implemented"; return nil }

func fileOpenOSOverride(f *FileDialog) bool { _ = "STUB: not implemented"; return false }

func fileSaveOSOverride(f *FileDialog) bool { _ = "STUB: not implemented"; return false }

func getFavoriteLocation(homeURI fyne.URI, name string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}
