package test

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/software"
)

type window struct {
	title              string
	fullScreen         bool
	fixedSize          bool
	focused            bool
	onClosed           func()
	onCloseIntercepted func()

	canvas software.WindowlessCanvas
	driver *driver
	menu   *fyne.MainMenu
}

func NewWindow(content fyne.CanvasObject) fyne.Window {
	_ = "STUB: not implemented"
	return *new(fyne.Window)
}

func (w *window) Canvas() fyne.Canvas { _ = "STUB: not implemented"; return *new(fyne.Canvas) }

func (*window) CenterOnScreen() { _ = "STUB: not implemented"; return }

func (*window) Clipboard() fyne.Clipboard { _ = "STUB: not implemented"; return *new(fyne.Clipboard) }

func (w *window) Close() { _ = "STUB: not implemented"; return }

func (w *window) Content() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (w *window) FixedSize() bool { _ = "STUB: not implemented"; return false }

func (w *window) FullScreen() bool { _ = "STUB: not implemented"; return false }

func (w *window) Hide() { _ = "STUB: not implemented"; return }

func (*window) Icon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func (w *window) MainMenu() *fyne.MainMenu { _ = "STUB: not implemented"; return nil }

func (w *window) Padded() bool { _ = "STUB: not implemented"; return false }

func (w *window) RequestFocus() { _ = "STUB: not implemented"; return }

func (w *window) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (w *window) SetContent(obj fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (w *window) SetFixedSize(fixed bool) { _ = "STUB: not implemented"; return }

func (*window) SetIcon(_ fyne.Resource) { _ = "STUB: not implemented"; return }

func (w *window) SetFullScreen(fullScreen bool) { _ = "STUB: not implemented"; return }

func (w *window) SetMainMenu(menu *fyne.MainMenu) { _ = "STUB: not implemented"; return }

func (*window) SetMaster() { _ = "STUB: not implemented"; return }

func (w *window) SetOnClosed(closed func()) { _ = "STUB: not implemented"; return }

func (w *window) SetCloseIntercept(callback func()) { _ = "STUB: not implemented"; return }

func (*window) SetOnDropped(func(fyne.Position, []fyne.URI)) { _ = "STUB: not implemented"; return }

func (w *window) SetPadded(padded bool) { _ = "STUB: not implemented"; return }

func (w *window) SetTitle(title string) { _ = "STUB: not implemented"; return }

func (w *window) Show() { _ = "STUB: not implemented"; return }

func (w *window) ShowAndRun() { _ = "STUB: not implemented"; return }

func (w *window) Title() string { _ = "STUB: not implemented"; return "" }
