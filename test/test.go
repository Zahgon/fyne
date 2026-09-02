package test

import (
	"fyne.io/fyne/v2"
)

func DoubleTap(obj fyne.DoubleTappable) { _ = "STUB: not implemented"; return }

func Drag(c fyne.Canvas, pos fyne.Position, deltaX, deltaY float32) {
	_ = "STUB: not implemented"
	return
}

func FocusNext(c fyne.Canvas) { _ = "STUB: not implemented"; return }

func FocusPrevious(c fyne.Canvas) { _ = "STUB: not implemented"; return }

func LaidOutObjects(o fyne.CanvasObject) (objects []fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return nil
}

func MoveMouse(c fyne.Canvas, pos fyne.Position) { _ = "STUB: not implemented"; return }

func RenderObjectToMarkup(o fyne.CanvasObject) string { _ = "STUB: not implemented"; return "" }

func RenderToMarkup(c fyne.Canvas) string { _ = "STUB: not implemented"; return "" }

func Scroll(c fyne.Canvas, pos fyne.Position, deltaX, deltaY float32) {
	_ = "STUB: not implemented"
	return
}

func Tap(obj fyne.Tappable) { _ = "STUB: not implemented"; return }

func TapAt(obj fyne.Tappable, pos fyne.Position) { _ = "STUB: not implemented"; return }

func TapCanvas(c fyne.Canvas, pos fyne.Position) { _ = "STUB: not implemented"; return }

func TapSecondary(obj fyne.SecondaryTappable) { _ = "STUB: not implemented"; return }

func TapSecondaryAt(obj fyne.SecondaryTappable, pos fyne.Position) {
	_ = "STUB: not implemented"
	return
}

func Type(obj fyne.Focusable, chars string) { _ = "STUB: not implemented"; return }

func TypeOnCanvas(c fyne.Canvas, chars string) { _ = "STUB: not implemented"; return }

func WidgetRenderer(wid fyne.Widget) fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func findTappable(c fyne.Canvas, pos fyne.Position) (o fyne.CanvasObject, p fyne.Position) {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject), *new(fyne.Position)
}

func handleFocusOnTap(c fyne.Canvas, obj any) { _ = "STUB: not implemented"; return }

func layoutAndCollect(objects []fyne.CanvasObject, o fyne.CanvasObject, size fyne.Size) []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func prepareTap(obj any, pos fyne.Position) (*fyne.PointEvent, fyne.Canvas) {
	_ = "STUB: not implemented"
	return nil, *new(fyne.Canvas)
}

func tap(c fyne.Canvas, obj fyne.Tappable, ev *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func typeChars(chars []rune, keyDown func(rune)) { _ = "STUB: not implemented"; return }

func writeMarkup(path string, markup string) error { _ = "STUB: not implemented"; return nil }
