package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/widget"
)

type ToolbarItem interface {
	ToolbarObject() fyne.CanvasObject
}

type ToolbarAction struct {
	Icon        fyne.Resource
	OnActivated func() `json:"-"`
	button      Button
}

func (t *ToolbarAction) ToolbarObject() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (t *ToolbarAction) SetIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func (t *ToolbarAction) Enable() { _ = "STUB: not implemented"; return }

func (t *ToolbarAction) Disable() { _ = "STUB: not implemented"; return }

func (t *ToolbarAction) Disabled() bool { _ = "STUB: not implemented"; return false }

func NewToolbarAction(icon fyne.Resource, onActivated func()) *ToolbarAction {
	_ = "STUB: not implemented"
	return nil
}

type ToolbarSpacer struct{}

func (*ToolbarSpacer) ToolbarObject() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func NewToolbarSpacer() *ToolbarSpacer { _ = "STUB: not implemented"; return nil }

type ToolbarSeparator struct{}

func (*ToolbarSeparator) ToolbarObject() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func NewToolbarSeparator() *ToolbarSeparator { _ = "STUB: not implemented"; return nil }

type Toolbar struct {
	BaseWidget
	Items []ToolbarItem
}

func (t *Toolbar) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *Toolbar) Append(item ToolbarItem) { _ = "STUB: not implemented"; return }

func (t *Toolbar) Prepend(item ToolbarItem) { _ = "STUB: not implemented"; return }

func (t *Toolbar) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func NewToolbar(items ...ToolbarItem) *Toolbar { _ = "STUB: not implemented"; return nil }

type toolbarRenderer struct {
	widget.BaseRenderer
	layout  fyne.Layout
	items   []fyne.CanvasObject
	toolbar *Toolbar
}

func (r *toolbarRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *toolbarRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *toolbarRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *toolbarRenderer) resetObjects() { _ = "STUB: not implemented"; return }
