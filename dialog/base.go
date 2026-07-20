package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	padWidth  = 32
	padHeight = 16
)

type Dialog interface {
	Show()
	Hide()
	SetDismissText(label string)
	SetOnClosed(closed func())
	Refresh()
	Resize(size fyne.Size)

	MinSize() fyne.Size

	Dismiss()
}

var _ Dialog = (*dialog)(nil)

type dialog struct {
	callback    func(bool)
	title       string
	icon        fyne.Resource
	desiredSize fyne.Size

	win     *widget.PopUp
	content fyne.CanvasObject
	dismiss *widget.Button
	parent  fyne.Window

	beforeShowHook func()
}

func (d *dialog) Dismiss() { _ = "STUB: not implemented"; return }

func (d *dialog) Hide() { _ = "STUB: not implemented"; return }

func (d *dialog) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (d *dialog) Show() { _ = "STUB: not implemented"; return }

func (d *dialog) Refresh() { _ = "STUB: not implemented"; return }

func (d *dialog) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (d *dialog) SetDismissText(label string) { _ = "STUB: not implemented"; return }

func (d *dialog) SetOnClosed(closed func()) { _ = "STUB: not implemented"; return }

func (d *dialog) hideWithResponse(resp bool) { _ = "STUB: not implemented"; return }

func (d *dialog) create(buttons fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (d *dialog) setButtons(buttons fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (d *dialog) setIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func newDialog(title, message string, icon fyne.Resource, callback func(bool), parent fyne.Window) *dialog {
	_ = "STUB: not implemented"
	return nil
}

type themedBackground struct {
	widget.BaseWidget
}

func newThemedBackground() *themedBackground { _ = "STUB: not implemented"; return nil }

func (t *themedBackground) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

type themedBackgroundRenderer struct {
	rect    *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (renderer *themedBackgroundRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (renderer *themedBackgroundRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (renderer *themedBackgroundRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (renderer *themedBackgroundRenderer) Objects() []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func (renderer *themedBackgroundRenderer) Refresh() { _ = "STUB: not implemented"; return }

//revive:disable-line:add-constant

type dialogLayout struct {
	d *dialog
}

func (l *dialogLayout) Layout(obj []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (l *dialogLayout) MinSize(obj []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
