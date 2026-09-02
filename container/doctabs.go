package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*DocTabs)(nil)

type DocTabs struct {
	widget.BaseWidget

	Items []*TabItem

	CreateTab      func() *TabItem `json:"-"`
	CloseIntercept func(*TabItem)  `json:"-"`
	OnClosed       func(*TabItem)  `json:"-"`
	OnSelected     func(*TabItem)  `json:"-"`
	OnUnselected   func(*TabItem)  `json:"-"`

	current         int
	location        TabLocation
	isTransitioning bool

	popUpMenu *widget.PopUpMenu
}

func NewDocTabs(items ...*TabItem) *DocTabs { _ = "STUB: not implemented"; return nil }

func (t *DocTabs) Append(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *DocTabs) DisableIndex(i int) { _ = "STUB: not implemented"; return }

func (t *DocTabs) DisableItem(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) EnableIndex(i int) { _ = "STUB: not implemented"; return }

func (t *DocTabs) EnableItem(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) Hide() { _ = "STUB: not implemented"; return }

func (t *DocTabs) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *DocTabs) Remove(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) RemoveIndex(index int) { _ = "STUB: not implemented"; return }

func (t *DocTabs) Select(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) SelectIndex(index int) { _ = "STUB: not implemented"; return }

func (t *DocTabs) Selected() *TabItem { _ = "STUB: not implemented"; return nil }

func (t *DocTabs) SelectedIndex() int { _ = "STUB: not implemented"; return 0 }

func (t *DocTabs) SetItems(items []*TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) SetTabLocation(l TabLocation) { _ = "STUB: not implemented"; return }

func (t *DocTabs) Show() { _ = "STUB: not implemented"; return }

func (t *DocTabs) close(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) onUnselected() func(*TabItem) { _ = "STUB: not implemented"; return nil }

func (t *DocTabs) onSelected() func(*TabItem) { _ = "STUB: not implemented"; return nil }

func (t *DocTabs) items() []*TabItem { _ = "STUB: not implemented"; return nil }

func (t *DocTabs) selected() int { _ = "STUB: not implemented"; return 0 }

func (t *DocTabs) setItems(items []*TabItem) { _ = "STUB: not implemented"; return }

func (t *DocTabs) setSelected(selected int) { _ = "STUB: not implemented"; return }

func (t *DocTabs) setTransitioning(transitioning bool) { _ = "STUB: not implemented"; return }

func (t *DocTabs) tabLocation() TabLocation { _ = "STUB: not implemented"; return *new(TabLocation) }

func (t *DocTabs) transitioning() bool { _ = "STUB: not implemented"; return false }

var _ fyne.WidgetRenderer = (*docTabsRenderer)(nil)

type docTabsRenderer struct {
	baseTabsRenderer
	docTabs      *DocTabs
	scroller     *Scroll
	box          *fyne.Container
	create       *widget.Button
	lastSelected int
}

func (r *docTabsRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *docTabsRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *docTabsRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *docTabsRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *docTabsRenderer) buildAllTabsButton() (all *widget.Button) {
	_ = "STUB: not implemented"
	return nil
}

func (r *docTabsRenderer) buildCreateTabsButton() *widget.Button {
	_ = "STUB: not implemented"
	return nil
}

func (r *docTabsRenderer) buildTabButtons(count int, buttons *fyne.Container) {
	_ = "STUB: not implemented"
	return
}

func (r *docTabsRenderer) scrollToSelected() { _ = "STUB: not implemented"; return }

func (r *docTabsRenderer) updateIndicator(animate bool) { _ = "STUB: not implemented"; return }

func (r *docTabsRenderer) updateAllTabs() { _ = "STUB: not implemented"; return }

func (r *docTabsRenderer) updateCreateTab() { _ = "STUB: not implemented"; return }

func (r *docTabsRenderer) updateTabs() { _ = "STUB: not implemented"; return }
