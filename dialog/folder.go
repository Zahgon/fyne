package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

var folderFilter = storage.NewMimeTypeFileFilter([]string{"application/x-directory"})

func NewFolderOpen(callback func(fyne.ListableURI, error), parent fyne.Window) *FileDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowFolderOpen(callback func(fyne.ListableURI, error), parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}

func (f *FileDialog) isDirectory() bool { _ = "STUB: not implemented"; return false }
