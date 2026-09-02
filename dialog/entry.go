package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type EntryDialog struct {
	*FormDialog

	entry *widget.Entry

	onClosed func()
}

func (i *EntryDialog) SetText(s string) { _ = "STUB: not implemented"; return }

func (i *EntryDialog) SetPlaceholder(s string) { _ = "STUB: not implemented"; return }

func (i *EntryDialog) SetOnClosed(callback func()) { _ = "STUB: not implemented"; return }

func NewEntryDialog(title, message string, onConfirm func(string), parent fyne.Window) *EntryDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowEntryDialog(title, message string, onConfirm func(string), parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}
