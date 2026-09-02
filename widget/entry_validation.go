package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var (
	_ fyne.Requireable = (*Entry)(nil)
	_ fyne.Validatable = (*Entry)(nil)
)

func (e *Entry) HasValue() bool { _ = "STUB: not implemented"; return false }

func (e *Entry) SetOnRequiredChanged(callback func(bool)) { _ = "STUB: not implemented"; return }

func (e *Entry) Validate() (err error) { _ = "STUB: not implemented"; return nil }

func (e *Entry) validate() { _ = "STUB: not implemented"; return }

func (e *Entry) SetOnValidationChanged(callback func(error)) { _ = "STUB: not implemented"; return }

func (e *Entry) SetValidationError(err error) { _ = "STUB: not implemented"; return }

func (e *Entry) setValidationError(err error) bool { _ = "STUB: not implemented"; return false }

var _ fyne.Widget = (*validationStatus)(nil)

type validationStatus struct {
	BaseWidget
	entry *Entry
}

func newValidationStatus(e *Entry) *validationStatus { _ = "STUB: not implemented"; return nil }

func (r *validationStatus) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

var _ fyne.WidgetRenderer = (*validationStatusRenderer)(nil)

type validationStatusRenderer struct {
	fyne.WidgetRenderer
	entry *Entry
	icon  *canvas.Image
}

func (r *validationStatusRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *validationStatusRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *validationStatusRenderer) Refresh() { _ = "STUB: not implemented"; return }
