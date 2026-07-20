package dialog

import (
	"fyne.io/fyne/v2"
)

var _ Dialog = (*CustomDialog)(nil)

type CustomDialog struct {
	*dialog
}

func NewCustom(title, dismiss string, content fyne.CanvasObject, parent fyne.Window) *CustomDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowCustom(title, dismiss string, content fyne.CanvasObject, parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}

func NewCustomWithoutButtons(title string, content fyne.CanvasObject, parent fyne.Window) *CustomDialog {
	_ = "STUB: not implemented"
	return nil
}

func (d *CustomDialog) SetButtons(buttons []fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (d *CustomDialog) SetIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func ShowCustomWithoutButtons(title string, content fyne.CanvasObject, parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}

func NewCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window,
) *ConfirmDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window,
) {
	_ = "STUB: not implemented"
	return
}
