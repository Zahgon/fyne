package container

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	intWidget "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

type titleBarButtonMode int

const (
	modeClose titleBarButtonMode = iota
	modeMinimize
	modeMaximize
	modeIcon

	sizeDraggableCorner = 16
)

var _ fyne.Widget = (*InnerWindow)(nil)

type InnerWindow struct {
	widget.BaseWidget

	CloseIntercept                                      func()                `json:"-"`
	OnDragged, OnResized                                func(*fyne.DragEvent) `json:"-"`
	OnMinimized, OnMaximized, OnTappedBar, OnTappedIcon func()                `json:"-"`
	Icon                                                fyne.Resource

	Alignment widget.ButtonAlign

	Title string

	Content *fyne.Container

	maximized, inactive bool
}

func NewInnerWindow(title string, content fyne.CanvasObject) *InnerWindow {
	_ = "STUB: not implemented"
	return nil
}

func (w *InnerWindow) Close() { _ = "STUB: not implemented"; return }

func (w *InnerWindow) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (w *InnerWindow) SetActive(active bool) { _ = "STUB: not implemented"; return }

func (w *InnerWindow) SetContent(obj fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (w *InnerWindow) SetMaximized(max bool) { _ = "STUB: not implemented"; return }

func (w *InnerWindow) SetPadded(pad bool) { _ = "STUB: not implemented"; return }

func (w *InnerWindow) SetTitle(title string) { _ = "STUB: not implemented"; return }

func (w *InnerWindow) buttonPosition() widget.ButtonAlign {
	_ = "STUB: not implemented"
	return *new(widget.ButtonAlign)
}

var _ fyne.WidgetRenderer = (*innerWindowRenderer)(nil)

type innerWindowRenderer struct {
	intWidget.BaseRenderer

	win            *InnerWindow
	bar, buttonBox *fyne.Container
	buttons        []*borderButton
	icon           *borderButton
	bg, contentBG  *canvas.Rectangle
	corner         fyne.CanvasObject
}

func (i *innerWindowRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (i *innerWindowRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (i *innerWindowRenderer) Refresh() { _ = "STUB: not implemented"; return }

type draggableLabel struct {
	widget.Label
	win *InnerWindow
}

func newDraggableLabel(title string, win *InnerWindow) *draggableLabel {
	_ = "STUB: not implemented"
	return nil
}

func (d *draggableLabel) Dragged(ev *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (d *draggableLabel) DragEnd() { _ = "STUB: not implemented"; return }

func (d *draggableLabel) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (d *draggableLabel) Tapped(_ *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (d *draggableLabel) labelMinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

type draggableCorner struct {
	widget.BaseWidget
	win *InnerWindow
}

func newDraggableCorner(w *InnerWindow) *draggableCorner { _ = "STUB: not implemented"; return nil }

func (c *draggableCorner) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *draggableCorner) Cursor() desktop.Cursor {
	_ = "STUB: not implemented"
	return *new(desktop.Cursor)
}

func (c *draggableCorner) Dragged(ev *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (c *draggableCorner) DragEnd() { _ = "STUB: not implemented"; return }

type borderButton struct {
	widget.BaseWidget

	b    *widget.Button
	c    *ThemeOverride
	mode titleBarButtonMode
}

func newBorderButton(icon fyne.Resource, mode titleBarButtonMode, th fyne.Theme, fn func()) *borderButton {
	_ = "STUB: not implemented"
	return nil
}

func (b *borderButton) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (b *borderButton) Disable() { _ = "STUB: not implemented"; return }

func (b *borderButton) Enable() { _ = "STUB: not implemented"; return }

func (b *borderButton) SetOnTapped(fn func()) { _ = "STUB: not implemented"; return }

func (b *borderButton) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (b *borderButton) setTheme(th fyne.Theme) { _ = "STUB: not implemented"; return }

type buttonTheme struct {
	fyne.Theme
	mode titleBarButtonMode
}

func (b *buttonTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (b *buttonTheme) Size(n fyne.ThemeSizeName) float32 { _ = "STUB: not implemented"; return 0 }

type titleBarLayout struct {
	win                  *InnerWindow
	buttons, icon, title fyne.CanvasObject
}

func (t *titleBarLayout) Layout(_ []fyne.CanvasObject, s fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (t *titleBarLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
