package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MultipleWindows struct {
	widget.BaseWidget

	Windows []*InnerWindow

	content *fyne.Container
}

func NewMultipleWindows(wins ...*InnerWindow) *MultipleWindows {
	_ = "STUB: not implemented"
	return nil
}

func (m *MultipleWindows) Add(w *InnerWindow) { _ = "STUB: not implemented"; return }

func (m *MultipleWindows) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (m *MultipleWindows) Refresh() { _ = "STUB: not implemented"; return }

func (m *MultipleWindows) RaiseToTop(w *InnerWindow) { _ = "STUB: not implemented"; return }

func (m *MultipleWindows) Top() *InnerWindow { _ = "STUB: not implemented"; return nil }

func (m *MultipleWindows) refreshChildren() { _ = "STUB: not implemented"; return }

func (m *MultipleWindows) setupChild(w *InnerWindow) { _ = "STUB: not implemented"; return }

type multiWinLayout struct{}

func (*multiWinLayout) Layout(objects []fyne.CanvasObject, _ fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (*multiWinLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
