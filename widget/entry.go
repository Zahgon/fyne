package widget

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/internal/widget"
)

const (
	bindIgnoreDelay = time.Millisecond * 100
	multiLineRows   = 3
)

var (
	_ fyne.Disableable    = (*Entry)(nil)
	_ fyne.Draggable      = (*Entry)(nil)
	_ fyne.Focusable      = (*Entry)(nil)
	_ fyne.Tappable       = (*Entry)(nil)
	_ fyne.Widget         = (*Entry)(nil)
	_ desktop.Mouseable   = (*Entry)(nil)
	_ desktop.Keyable     = (*Entry)(nil)
	_ mobile.Keyboardable = (*Entry)(nil)
	_ mobile.Touchable    = (*Entry)(nil)
	_ fyne.Tabbable       = (*Entry)(nil)
)

type Entry struct {
	DisableableWidget
	shortcut fyne.ShortcutHandler
	Text     string

	TextStyle   fyne.TextStyle
	PlaceHolder string
	OnChanged   func(string) `json:"-"`

	OnSubmitted func(string) `json:"-"`
	Password    bool
	MultiLine   bool
	Wrapping    fyne.TextWrap

	Scroll fyne.ScrollDirection

	Validator           fyne.StringValidator `json:"-"`
	validationStatus    *validationStatus
	onValidationChanged func(error)
	validationError     error
	onRequiredChanged   func(bool)

	AlwaysShowValidationError bool

	CursorRow, CursorColumn int
	OnCursorChanged         func() `json:"-"`

	Icon fyne.Resource `json:"-"`

	cursorAnim *entryCursorAnimation

	dirty               bool
	focused, hasFocused bool
	text                RichText
	placeholder         RichText
	content             *entryContent
	scroll              *widget.Scroll

	onFocusChanged func(bool)

	selectKeyDown bool

	sel   *selectable
	popUp *PopUpMenu

	ActionItem      fyne.CanvasObject `json:"-"`
	binder          basicBinder
	conversionError error
	minCache        fyne.Size
	multiLineRows   int

	undoStack entryUndoStack
}

func NewEntry() *Entry { _ = "STUB: not implemented"; return nil }

func NewEntryWithData(data binding.String) *Entry { _ = "STUB: not implemented"; return nil }

func NewMultiLineEntry() *Entry { _ = "STUB: not implemented"; return nil }

func NewPasswordEntry() *Entry { _ = "STUB: not implemented"; return nil }

func (e *Entry) AcceptsTab() bool { _ = "STUB: not implemented"; return false }

func (e *Entry) Bind(data binding.String) { _ = "STUB: not implemented"; return }

func (e *Entry) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (e *Entry) CursorPosition() fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func (e *Entry) CursorTextOffset() (pos int) { _ = "STUB: not implemented"; return 0 }

func (*Entry) Cursor() desktop.Cursor { _ = "STUB: not implemented"; return *new(desktop.Cursor) }

func (e *Entry) DoubleTapped(_ *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) DragEnd() { _ = "STUB: not implemented"; return }

func (e *Entry) Dragged(d *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) ExtendBaseWidget(wid fyne.Widget) { _ = "STUB: not implemented"; return }

func (e *Entry) FocusGained() { _ = "STUB: not implemented"; return }

func (e *Entry) FocusLost() { _ = "STUB: not implemented"; return }

func (e *Entry) Hide() { _ = "STUB: not implemented"; return }

func (e *Entry) Keyboard() mobile.KeyboardType {
	_ = "STUB: not implemented"
	return *new(mobile.KeyboardType)
}

func (e *Entry) KeyDown(key *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) KeyUp(key *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (e *Entry) MouseDown(m *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) MouseUp(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) Redo() { _ = "STUB: not implemented"; return }

func (e *Entry) Refresh() { _ = "STUB: not implemented"; return }

func (e *Entry) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (e *Entry) ClearSelection() { _ = "STUB: not implemented"; return }

func (e *Entry) SetIcon(res fyne.Resource) { _ = "STUB: not implemented"; return }

func (e *Entry) SetMinRowsVisible(count int) { _ = "STUB: not implemented"; return }

func (e *Entry) SetPlaceHolder(text string) { _ = "STUB: not implemented"; return }

func (e *Entry) SetText(text string) { _ = "STUB: not implemented"; return }

func (e *Entry) setText(text string, fromBinding bool) { _ = "STUB: not implemented"; return }

func (e *Entry) Append(text string) { _ = "STUB: not implemented"; return }

func (*Entry) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) TappedSecondary(pe *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) TouchDown(ev *mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (*Entry) TouchUp(*mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (*Entry) TouchCancel(*mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) TypedKey(key *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (e *Entry) Undo() { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyUp(provider *RichText) { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyDown(provider *RichText) { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyLeft(provider *RichText) { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyRight(provider *RichText) { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyHome() { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyEnd(provider *RichText) { _ = "STUB: not implemented"; return }

func (e *Entry) deleteWord(right bool) { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyTab() { _ = "STUB: not implemented"; return }

func (e *Entry) TypedRune(r rune) { _ = "STUB: not implemented"; return }

func (e *Entry) TypedShortcut(shortcut fyne.Shortcut) { _ = "STUB: not implemented"; return }

func (e *Entry) Unbind() { _ = "STUB: not implemented"; return }

func (e *Entry) copyToClipboard(clipboard fyne.Clipboard) { _ = "STUB: not implemented"; return }

func (e *Entry) cutToClipboard(clipboard fyne.Clipboard) { _ = "STUB: not implemented"; return }

func (e *Entry) eraseSelection() bool { _ = "STUB: not implemented"; return false }

func (e *Entry) eraseSelectionAndUpdate() { _ = "STUB: not implemented"; return }

func (e *Entry) pasteFromClipboard(clipboard fyne.Clipboard) { _ = "STUB: not implemented"; return }

func (e *Entry) placeholderProvider() *RichText { _ = "STUB: not implemented"; return nil }

func (e *Entry) registerShortcut() { _ = "STUB: not implemented"; return }

func (e *Entry) requestFocus() { _ = "STUB: not implemented"; return }

func (e *Entry) rowColFromTextPos(pos int) (row int, col int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (e *Entry) selectAll() { _ = "STUB: not implemented"; return }

func (e *Entry) selectingKeyHandler(key *fyne.KeyEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Entry) syncSegments() { _ = "STUB: not implemented"; return }

func (e *Entry) syncSelectable() { _ = "STUB: not implemented"; return }

func (e *Entry) textProvider() *RichText { _ = "STUB: not implemented"; return nil }

func (e *Entry) textWrap() fyne.TextWrap { _ = "STUB: not implemented"; return *new(fyne.TextWrap) }

func (e *Entry) updateCursorAndSelection() { _ = "STUB: not implemented"; return }

func (e *Entry) updateFromData(data binding.DataItem) { _ = "STUB: not implemented"; return }

func (e *Entry) truncatePosition(row, col int) (newRow, newCol int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (e *Entry) updateMousePointer(p fyne.Position, rightClick bool) {
	_ = "STUB: not implemented"
	return
}

func (e *Entry) updateText(text string, fromBinding bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Entry) updateTextAndRefresh(text string, fromBinding bool) {
	_ = "STUB: not implemented"
	return
}

func (e *Entry) writeData(data binding.DataItem) { _ = "STUB: not implemented"; return }

func (e *Entry) typedKeyReturn(provider *RichText, multiLine bool) {
	_ = "STUB: not implemented"
	return
}

func (e *Entry) setFieldsAndRefresh(f func()) { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*entryRenderer)(nil)

type entryRenderer struct {
	box, border *canvas.Rectangle
	scroll      *widget.Scroll
	icon        *canvas.Image

	objects []fyne.CanvasObject
	entry   *Entry
}

func (*entryRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *entryRenderer) trailingInset() float32 { _ = "STUB: not implemented"; return 0 }

func (r *entryRenderer) leadingInset() float32 { _ = "STUB: not implemented"; return 0 }

func (r *entryRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *entryRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *entryRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *entryRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *entryRenderer) ensureValidationSetup() { _ = "STUB: not implemented"; return }

var _ fyne.Widget = (*entryContent)(nil)

type entryContent struct {
	BaseWidget

	entry  *Entry
	scroll *widget.Scroll
}

func (e *entryContent) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (e *entryContent) DragEnd() { _ = "STUB: not implemented"; return }

func (e *entryContent) Dragged(d *fyne.DragEvent) { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*entryContentRenderer)(nil)

type entryContentRenderer struct {
	cursor  *canvas.Rectangle
	objects []fyne.CanvasObject

	provider, placeholder *RichText
	content               *entryContent
}

func (r *entryContentRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *entryContentRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *entryContentRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *entryContentRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *entryContentRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *entryContentRenderer) ensureCursorVisible() { _ = "STUB: not implemented"; return }

func (r *entryContentRenderer) moveCursor() { _ = "STUB: not implemented"; return }

func (r *entryContentRenderer) updateScrollDirections() { _ = "STUB: not implemented"; return }

func getTextWhitespaceRegion(row []rune, col int, expand bool) (start, end int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func isWordSeparator(r rune) bool { _ = "STUB: not implemented"; return false }

type entryUndoAction interface {
	Undo(string) string
	Redo(string) string
}

type entryMergeableUndoAction interface {
	entryUndoAction

	TryMerge(next entryMergeableUndoAction) bool
}

var _ entryMergeableUndoAction = (*entryModifyAction)(nil)

type entryModifyAction struct {
	Delete bool

	Position int

	Text []rune
}

func (i *entryModifyAction) Undo(s string) string { _ = "STUB: not implemented"; return "" }

func (i *entryModifyAction) Redo(s string) string { _ = "STUB: not implemented"; return "" }

func (i *entryModifyAction) add(s string) string { _ = "STUB: not implemented"; return "" }

func (i *entryModifyAction) sub(s string) string { _ = "STUB: not implemented"; return "" }

func (i *entryModifyAction) TryMerge(other entryMergeableUndoAction) bool {
	_ = "STUB: not implemented"
	return false
}

type entryUndoStack struct {
	items []entryUndoAction

	index int
}

func (u *entryUndoStack) Undo(s string) (newS string, action entryUndoAction) {
	_ = "STUB: not implemented"
	return "", *new(entryUndoAction)
}

func (u *entryUndoStack) Redo(s string) (newS string, action entryUndoAction) {
	_ = "STUB: not implemented"
	return "", *new(entryUndoAction)
}

func (u *entryUndoStack) CanUndo() bool { _ = "STUB: not implemented"; return false }

func (u *entryUndoStack) CanRedo() bool { _ = "STUB: not implemented"; return false }

func (u *entryUndoStack) Add(a entryUndoAction) { _ = "STUB: not implemented"; return }

func (u *entryUndoStack) MergeOrAdd(a entryUndoAction) { _ = "STUB: not implemented"; return }

func (u *entryUndoStack) Clear() { _ = "STUB: not implemented"; return }
