package dialog

import (
	"fyne.io/fyne/v2"
)

func createInformationDialog(title, message string, icon fyne.Resource, parent fyne.Window) Dialog {
	_ = "STUB: not implemented"
	return *new(Dialog)
}

func NewInformation(title, message string, parent fyne.Window) Dialog {
	_ = "STUB: not implemented"
	return *new(Dialog)
}

func ShowInformation(title, message string, parent fyne.Window) { _ = "STUB: not implemented"; return }

func NewError(err error, parent fyne.Window) Dialog { _ = "STUB: not implemented"; return *new(Dialog) }

func ShowError(err error, parent fyne.Window) { _ = "STUB: not implemented"; return }
