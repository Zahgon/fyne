package dialog

import (
	"fyne.io/fyne/v2"
)

func driveMask() uint32 { _ = "STUB: not implemented"; return 0 }

func listDrives() []string { _ = "STUB: not implemented"; return nil }

func (f *fileDialog) getPlaces() []favoriteItem { _ = "STUB: not implemented"; return nil }

func isHidden(file fyne.URI) bool { _ = "STUB: not implemented"; return false }

func hideFile(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func fileOpenOSOverride(*FileDialog) bool { _ = "STUB: not implemented"; return false }

func fileSaveOSOverride(*FileDialog) bool { _ = "STUB: not implemented"; return false }

func getFavoriteLocation(homeURI fyne.URI, name string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}
