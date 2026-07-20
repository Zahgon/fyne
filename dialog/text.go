package dialog

import (
	"fyne.io/fyne/v2"
)

const (
	maxTextDialogAbsoluteWidth float32 = 600

	maxTextDialogWinPcntWidth float32 = .9
)

func newTextDialog(title, message string, icon fyne.Resource, parent fyne.Window) *dialog {
	_ = "STUB: not implemented"
	return nil
}

func createBeforeShowHook(d *dialog, message string) func() { _ = "STUB: not implemented"; return nil }

func newCenterWrappedLabel(message string) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}
