package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/widget"
)

var _ fyne.Widget = (*Accordion)(nil)

type Accordion struct {
	BaseWidget
	Items     []*AccordionItem
	MultiOpen bool
}

func NewAccordion(items ...*AccordionItem) *Accordion { _ = "STUB: not implemented"; return nil }

func (a *Accordion) Append(item *AccordionItem) { _ = "STUB: not implemented"; return }

func (a *Accordion) Close(index int) { _ = "STUB: not implemented"; return }

func (a *Accordion) CloseAll() { _ = "STUB: not implemented"; return }

func (a *Accordion) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (a *Accordion) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (a *Accordion) Open(index int) { _ = "STUB: not implemented"; return }

func (a *Accordion) OpenAll() { _ = "STUB: not implemented"; return }

func (a *Accordion) Prepend(item *AccordionItem) { _ = "STUB: not implemented"; return }

func (a *Accordion) Remove(item *AccordionItem) { _ = "STUB: not implemented"; return }

func (a *Accordion) RemoveIndex(index int) { _ = "STUB: not implemented"; return }

type accordionRenderer struct {
	widget.BaseRenderer
	container    *Accordion
	headers      []*Button
	dividers     []fyne.CanvasObject
	minSizeCache fyne.Size
}

func (r *accordionRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *accordionRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *accordionRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *accordionRenderer) updateObjects() { _ = "STUB: not implemented"; return }

type AccordionItem struct {
	Title  string
	Detail fyne.CanvasObject
	Open   bool
}

func NewAccordionItem(title string, detail fyne.CanvasObject) *AccordionItem {
	_ = "STUB: not implemented"
	return nil
}
