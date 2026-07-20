package software

import (
	"image"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver"
	"fyne.io/fyne/v2/internal"
	"fyne.io/fyne/v2/internal/app"
)

const canvasDefaultSize = 100

type WindowlessCanvas interface {
	fyne.Canvas

	Padded() bool
	Resize(fyne.Size)
	SetPadded(bool)
	SetScale(float32)
}

func NewCanvas() WindowlessCanvas { _ = "STUB: not implemented"; return *new(WindowlessCanvas) }

func NewCanvasWithPainter(painter driver.Painter) WindowlessCanvas {
	_ = "STUB: not implemented"
	return *new(WindowlessCanvas)
}

func NewTransparentCanvas() WindowlessCanvas {
	_ = "STUB: not implemented"
	return *new(WindowlessCanvas)
}

func NewTransparentCanvasWithPainter(painter driver.Painter) WindowlessCanvas {
	_ = "STUB: not implemented"
	return *new(WindowlessCanvas)
}

func newCanvas(painter driver.Painter, transparent bool) WindowlessCanvas {
	_ = "STUB: not implemented"
	return *new(WindowlessCanvas)
}

type canvas struct {
	size    fyne.Size
	resized bool
	scale   float32

	content     fyne.CanvasObject
	overlays    internal.OverlayStack
	focusMgr    *app.FocusManager
	padded      bool
	transparent bool

	onTypedRune func(rune)
	onTypedKey  func(*fyne.KeyEvent)

	fyne.ShortcutHandler
	painter      driver.Painter
	propertyLock sync.RWMutex
}

func (c *canvas) Capture() image.Image { _ = "STUB: not implemented"; return *new(image.Image) }

func (c *canvas) Content() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (c *canvas) Focus(obj fyne.Focusable) { _ = "STUB: not implemented"; return }

func (c *canvas) FocusNext() { _ = "STUB: not implemented"; return }

func (c *canvas) FocusPrevious() { _ = "STUB: not implemented"; return }

func (c *canvas) Focused() fyne.Focusable { _ = "STUB: not implemented"; return *new(fyne.Focusable) }

func (c *canvas) InteractiveArea() (fyne.Position, fyne.Size) {
	_ = "STUB: not implemented"
	return *new(fyne.Position), *new(fyne.Size)
}

func (c *canvas) OnTypedKey() func(*fyne.KeyEvent) { _ = "STUB: not implemented"; return nil }

func (c *canvas) OnTypedRune() func(rune) { _ = "STUB: not implemented"; return nil }

func (c *canvas) Overlays() fyne.OverlayStack {
	_ = "STUB: not implemented"
	return *new(fyne.OverlayStack)
}

func (c *canvas) Padded() bool { _ = "STUB: not implemented"; return false }

func (c *canvas) PixelCoordinateForPosition(pos fyne.Position) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (c *canvas) Refresh(fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (c *canvas) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (c *canvas) Scale() float32 { _ = "STUB: not implemented"; return 0 }

func (c *canvas) SetContent(content fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (c *canvas) SetOnTypedKey(handler func(*fyne.KeyEvent)) { _ = "STUB: not implemented"; return }

func (c *canvas) SetOnTypedRune(handler func(rune)) { _ = "STUB: not implemented"; return }

func (c *canvas) SetPadded(padded bool) { _ = "STUB: not implemented"; return }

func (c *canvas) SetScale(scale float32) { _ = "STUB: not implemented"; return }

func (c *canvas) Size() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *canvas) Unfocus() { _ = "STUB: not implemented"; return }

func (c *canvas) doResize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (c *canvas) focusManager() *app.FocusManager { _ = "STUB: not implemented"; return nil }
