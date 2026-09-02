package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type FormDialog struct {
	*dialog
	items   []*widget.FormItem
	confirm *widget.Button
	cancel  *widget.Button
}

func (d *FormDialog) Submit() { _ = "STUB: not implemented"; return }

func (d *FormDialog) setSubmitState(err error) { _ = "STUB: not implemented"; return }

func NewForm(title, confirm, dismiss string, items []*widget.FormItem, callback func(bool), parent fyne.Window) *FormDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowForm(title, confirm, dismiss string, content []*widget.FormItem, callback func(bool), parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}
