package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ConfirmDialog struct {
	*dialog

	confirm *widget.Button
}

func (d *ConfirmDialog) Confirm() { _ = "STUB: not implemented"; return }

func (d *ConfirmDialog) SetConfirmText(label string) { _ = "STUB: not implemented"; return }

func (d *ConfirmDialog) SetConfirmImportance(importance widget.Importance) {
	_ = "STUB: not implemented"
	return
}

func NewConfirm(title, message string, callback func(bool), parent fyne.Window) *ConfirmDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowConfirm(title, message string, callback func(bool), parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}
