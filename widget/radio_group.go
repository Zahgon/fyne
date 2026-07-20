package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/widget"
)

type RadioGroup struct {
	DisableableWidget
	Horizontal bool
	Required   bool
	OnChanged  func(string) `json:"-"`
	Options    []string
	Selected   string

	_selIdx int
}

var _ fyne.Widget = (*RadioGroup)(nil)

func NewRadioGroup(options []string, changed func(string)) *RadioGroup {
	_ = "STUB: not implemented"
	return nil
}

func (r *RadioGroup) Append(option string) { _ = "STUB: not implemented"; return }

func (r *RadioGroup) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (r *RadioGroup) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *RadioGroup) SetSelected(option string) { _ = "STUB: not implemented"; return }

func (r *RadioGroup) itemTapped(item *radioItem, idx int) { _ = "STUB: not implemented"; return }

func (r *RadioGroup) Refresh() { _ = "STUB: not implemented"; return }

func (r *RadioGroup) selectedIndex() int { _ = "STUB: not implemented"; return 0 }

func (r *RadioGroup) setSelectedIndex(idx int) { _ = "STUB: not implemented"; return }

func (r *RadioGroup) updateSelectedIndex() { _ = "STUB: not implemented"; return }

type radioGroupRenderer struct {
	widget.BaseRenderer

	items []fyne.CanvasObject
	radio *RadioGroup
}

func (r *radioGroupRenderer) Layout(_ fyne.Size) { _ = "STUB: not implemented"; return }

func (r *radioGroupRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *radioGroupRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *radioGroupRenderer) updateItems(refresh bool) { _ = "STUB: not implemented"; return }
