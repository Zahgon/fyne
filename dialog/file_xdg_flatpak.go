//go:build flatpak && !windows && !android && !ios && !wasm && !js

package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"

	"github.com/rymdport/portal/filechooser"
)

func openFile(parentWindowHandle string, options *filechooser.OpenFileOptions) (fyne.URIReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIReadCloser), nil
}

func openFolder(parentWindowHandle string, options *filechooser.OpenFileOptions) (fyne.ListableURI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.ListableURI), nil
}

func open(parentWindowHandle, title string, options *filechooser.OpenFileOptions) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}

func saveFile(parentWindowHandle string, options *filechooser.SaveFileOptions) (fyne.URIWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URIWriteCloser), nil
}

func fileOpenOSOverride(d *FileDialog) bool { _ = "STUB: not implemented"; return false }

func fileSaveOSOverride(d *FileDialog) bool { _ = "STUB: not implemented"; return false }

func windowHandleForPortal(window fyne.Window) string { _ = "STUB: not implemented"; return "" }

func convertFilterForPortal(fyneFilter storage.FileFilter) (list []*filechooser.Filter, current *filechooser.Filter) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatFilterName(patterns []string, count int) string { _ = "STUB: not implemented"; return "" }
