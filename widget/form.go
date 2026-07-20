package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type FormItem struct {
	Text   string
	Widget fyne.CanvasObject

	HintText string

	Required bool

	validationError        error
	invalid, meetsRequired bool
	helperOutput           *canvas.Text
	wasFocused             bool
}

func NewFormItem(text string, widget fyne.CanvasObject) *FormItem {
	_ = "STUB: not implemented"
	return nil
}

var _ fyne.Validatable = (*Form)(nil)

type Form struct {
	BaseWidget

	Items      []*FormItem
	OnSubmit   func() `json:"-"`
	OnCancel   func() `json:"-"`
	SubmitText string
	CancelText string

	Orientation Orientation

	Validator func() error `json:"-"`

	itemGrid     *fyne.Container
	buttonBox    *fyne.Container
	validateText *canvas.Text
	cancelButton *Button
	submitButton *Button

	disabled bool

	onValidationChanged func(error)
	validationError     error
}

func (f *Form) Append(text string, widget fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (f *Form) AppendItem(item *FormItem) { _ = "STUB: not implemented"; return }

func (f *Form) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (f *Form) Refresh() { _ = "STUB: not implemented"; return }

func (f *Form) Enable() { _ = "STUB: not implemented"; return }

func (f *Form) Disable() { _ = "STUB: not implemented"; return }

func (f *Form) Disabled() bool { _ = "STUB: not implemented"; return false }

func (f *Form) RemoveItem(item *FormItem) { _ = "STUB: not implemented"; return }

func (f *Form) SetOnValidationChanged(callback func(error)) { _ = "STUB: not implemented"; return }

func (f *Form) Validate() error { _ = "STUB: not implemented"; return nil }

func (f *Form) createInput(item *FormItem) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (f *Form) itemWidgetHasValidator(w fyne.CanvasObject) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *Form) createLabel(item *FormItem) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (f *Form) updateButtons() { _ = "STUB: not implemented"; return }

func (f *Form) checkValidation(err error) { _ = "STUB: not implemented"; return }

func (f *Form) ensureRenderItems() { _ = "STUB: not implemented"; return }

func (f *Form) isVertical() bool { _ = "STUB: not implemented"; return false }

func (f *Form) setUpValidation(widget fyne.CanvasObject, i int) { _ = "STUB: not implemented"; return }

func (f *Form) setValidationError(err error) { _ = "STUB: not implemented"; return }

func (f *Form) updateHelperText(item *FormItem) { _ = "STUB: not implemented"; return }

func (f *Form) updateLabels() { _ = "STUB: not implemented"; return }

func (f *Form) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func NewForm(items ...*FormItem) *Form { _ = "STUB: not implemented"; return nil }

type formItemLayout struct {
	form *Form
}

func (f formItemLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (f formItemLayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
