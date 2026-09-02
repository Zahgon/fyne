package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/widget"
)

var _ fyne.Widget = (*CheckGroup)(nil)

type CheckGroup struct {
	DisableableWidget
	Horizontal bool
	Required   bool
	OnChanged  func([]string) `json:"-"`
	Options    []string
	Selected   []string

	items []*Check
}

func NewCheckGroup(options []string, changed func([]string)) *CheckGroup {
	_ = "STUB: not implemented"
	return nil
}

func (r *CheckGroup) Append(option string) { _ = "STUB: not implemented"; return }

func (r *CheckGroup) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (r *CheckGroup) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *CheckGroup) Refresh() { _ = "STUB: not implemented"; return }

func (r *CheckGroup) Remove(option string) bool { _ = "STUB: not implemented"; return false }

func (r *CheckGroup) SetSelected(options []string) { _ = "STUB: not implemented"; return }

func (r *CheckGroup) itemTapped(item *Check) { _ = "STUB: not implemented"; return }

func (r *CheckGroup) update() { _ = "STUB: not implemented"; return }

type checkGroupRenderer struct {
	widget.BaseRenderer
	items  []*Check
	checks *CheckGroup
}

func (r *checkGroupRenderer) Layout(_ fyne.Size) { _ = "STUB: not implemented"; return }

func (r *checkGroupRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *checkGroupRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *checkGroupRenderer) updateItems() { _ = "STUB: not implemented"; return }

func removeDuplicates(options []string) []string { _ = "STUB: not implemented"; return nil }
