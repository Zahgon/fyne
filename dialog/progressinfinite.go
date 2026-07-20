package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ProgressInfiniteDialog struct {
	*dialog

	bar *widget.ProgressBarInfinite
}

func NewProgressInfinite(title, message string, parent fyne.Window) *ProgressInfiniteDialog {
	_ = "STUB: not implemented"
	return nil
}

//revive:disable-line:add-constant

func (d *ProgressInfiniteDialog) Hide() { _ = "STUB: not implemented"; return }
