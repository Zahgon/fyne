package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ProgressDialog struct {
	*dialog

	bar *widget.ProgressBar
}

func (p *ProgressDialog) SetValue(v float64) { _ = "STUB: not implemented"; return }

func NewProgress(title, message string, parent fyne.Window) *ProgressDialog {
	_ = "STUB: not implemented"
	return nil
}

//revive:disable-line:add-constant
