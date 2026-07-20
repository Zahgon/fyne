package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*AppTabs)(nil)

type AppTabs struct {
	widget.BaseWidget

	Items []*TabItem

	OnChanged    func(*TabItem) `json:"-"`
	OnSelected   func(*TabItem) `json:"-"`
	OnUnselected func(*TabItem) `json:"-"`

	current         int
	location        TabLocation
	isTransitioning bool

	popUpMenu *widget.PopUpMenu
}

func NewAppTabs(items ...*TabItem) *AppTabs { _ = "STUB: not implemented"; return nil }

func (t *AppTabs) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *AppTabs) Append(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) CurrentTab() *TabItem { _ = "STUB: not implemented"; return nil }

func (t *AppTabs) CurrentTabIndex() int { _ = "STUB: not implemented"; return 0 }

func (t *AppTabs) DisableIndex(i int) { _ = "STUB: not implemented"; return }

func (t *AppTabs) DisableItem(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) EnableIndex(i int) { _ = "STUB: not implemented"; return }

func (t *AppTabs) EnableItem(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) ExtendBaseWidget(wid fyne.Widget) { _ = "STUB: not implemented"; return }

func (t *AppTabs) Hide() { _ = "STUB: not implemented"; return }

func (t *AppTabs) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *AppTabs) Remove(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) RemoveIndex(index int) { _ = "STUB: not implemented"; return }

func (t *AppTabs) Select(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) SelectIndex(index int) { _ = "STUB: not implemented"; return }

func (t *AppTabs) SelectTab(item *TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) SelectTabIndex(index int) { _ = "STUB: not implemented"; return }

func (t *AppTabs) Selected() *TabItem { _ = "STUB: not implemented"; return nil }

func (t *AppTabs) SelectedIndex() int { _ = "STUB: not implemented"; return 0 }

func (t *AppTabs) SetItems(items []*TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) SetTabLocation(l TabLocation) { _ = "STUB: not implemented"; return }

func (t *AppTabs) Show() { _ = "STUB: not implemented"; return }

func (t *AppTabs) onUnselected() func(*TabItem) { _ = "STUB: not implemented"; return nil }

func (t *AppTabs) onSelected() func(*TabItem) { _ = "STUB: not implemented"; return nil }

func (t *AppTabs) items() []*TabItem { _ = "STUB: not implemented"; return nil }

func (t *AppTabs) selected() int { _ = "STUB: not implemented"; return 0 }

func (t *AppTabs) setItems(items []*TabItem) { _ = "STUB: not implemented"; return }

func (t *AppTabs) setSelected(selected int) { _ = "STUB: not implemented"; return }

func (t *AppTabs) setTransitioning(transitioning bool) { _ = "STUB: not implemented"; return }

func (t *AppTabs) tabLocation() TabLocation { _ = "STUB: not implemented"; return *new(TabLocation) }

func (t *AppTabs) transitioning() bool { _ = "STUB: not implemented"; return false }

var _ fyne.WidgetRenderer = (*appTabsRenderer)(nil)

type appTabsRenderer struct {
	baseTabsRenderer
	appTabs *AppTabs
}

func (r *appTabsRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *appTabsRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *appTabsRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *appTabsRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *appTabsRenderer) buildOverflowTabsButton() (overflow *widget.Button) {
	_ = "STUB: not implemented"
	return nil
}

func (r *appTabsRenderer) buildTabButtons(count int) *fyne.Container {
	_ = "STUB: not implemented"
	return nil
}

func (r *appTabsRenderer) updateIndicator(animate bool) { _ = "STUB: not implemented"; return }

func (r *appTabsRenderer) updateTabs(max int) { _ = "STUB: not implemented"; return }
