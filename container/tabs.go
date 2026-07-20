package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type TabItem struct {
	Text    string
	Icon    fyne.Resource
	Content fyne.CanvasObject

	button *tabButton

	disabled bool
}

func (ti *TabItem) Disabled() bool { _ = "STUB: not implemented"; return false }

func (ti *TabItem) disable() { _ = "STUB: not implemented"; return }

func (ti *TabItem) enable() { _ = "STUB: not implemented"; return }

type TabLocation int

const (
	TabLocationTop TabLocation = iota
	TabLocationLeading
	TabLocationBottom
	TabLocationTrailing
)

func NewTabItem(text string, content fyne.CanvasObject) *TabItem {
	_ = "STUB: not implemented"
	return nil
}

func NewTabItemWithIcon(text string, icon fyne.Resource, content fyne.CanvasObject) *TabItem {
	_ = "STUB: not implemented"
	return nil
}

type baseTabs interface {
	fyne.Widget

	onUnselected() func(*TabItem)
	onSelected() func(*TabItem)

	items() []*TabItem
	setItems([]*TabItem)

	selected() int
	setSelected(int)

	tabLocation() TabLocation

	transitioning() bool
	setTransitioning(bool)
}

func isMobile(b baseTabs) bool { _ = "STUB: not implemented"; return false }

func tabsAdjustedLocation(l TabLocation, b baseTabs) TabLocation {
	_ = "STUB: not implemented"
	return *new(TabLocation)
}

func buildPopUpMenu(t baseTabs, button *widget.Button, items []*fyne.MenuItem) *widget.PopUpMenu {
	_ = "STUB: not implemented"
	return nil
}

func removeIndex(t baseTabs, index int) { _ = "STUB: not implemented"; return }

func removeItem(t baseTabs, item *TabItem) { _ = "STUB: not implemented"; return }

func selected(t baseTabs) *TabItem { _ = "STUB: not implemented"; return nil }

func selectIndex(t baseTabs, index int) { _ = "STUB: not implemented"; return }

func selectItem(t baseTabs, item *TabItem) { _ = "STUB: not implemented"; return }

func setItems(t baseTabs, items []*TabItem) { _ = "STUB: not implemented"; return }

func disableIndex(t baseTabs, index int) { _ = "STUB: not implemented"; return }

func disableItem(t baseTabs, item *TabItem) { _ = "STUB: not implemented"; return }

func enableIndex(t baseTabs, index int) { _ = "STUB: not implemented"; return }

func enableItem(t baseTabs, item *TabItem) { _ = "STUB: not implemented"; return }

type baseTabsRenderer struct {
	positionAnimation, sizeAnimation *fyne.Animation

	lastIndicatorPos    fyne.Position
	lastIndicatorSize   fyne.Size
	lastIndicatorHidden bool

	action             *widget.Button
	bar                *fyne.Container
	divider, indicator *canvas.Rectangle

	tabs baseTabs
}

func (r *baseTabsRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *baseTabsRenderer) applyTheme(t baseTabs) { _ = "STUB: not implemented"; return }

func (r *baseTabsRenderer) layout(t baseTabs, size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *baseTabsRenderer) minSize(t baseTabs) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *baseTabsRenderer) moveIndicator(pos fyne.Position, siz fyne.Size, th fyne.Theme, animate bool) {
	_ = "STUB: not implemented"
	return
}

func (r *baseTabsRenderer) objects(t baseTabs) []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func (r *baseTabsRenderer) refresh(t baseTabs) { _ = "STUB: not implemented"; return }

type buttonIconPosition int

const (
	buttonIconInline buttonIconPosition = iota
	buttonIconTop
)

var (
	_ fyne.Widget       = (*tabButton)(nil)
	_ fyne.Tappable     = (*tabButton)(nil)
	_ desktop.Hoverable = (*tabButton)(nil)
)

type tabButton struct {
	widget.DisableableWidget
	hovered       bool
	icon          fyne.Resource
	iconPosition  buttonIconPosition
	importance    widget.Importance
	onTapped      func()
	onClosed      func()
	text          string
	textAlignment fyne.TextAlign

	tabs baseTabs
}

func (b *tabButton) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (b *tabButton) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (b *tabButton) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *tabButton) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *tabButton) MouseOut() { _ = "STUB: not implemented"; return }

func (b *tabButton) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

type tabButtonRenderer struct {
	button     *tabButton
	background *canvas.Rectangle
	icon       *canvas.Image
	label      *canvas.Text
	close      *tabCloseButton
	objects    []fyne.CanvasObject
}

func (r *tabButtonRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *tabButtonRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *tabButtonRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *tabButtonRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *tabButtonRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *tabButtonRenderer) iconSize() float32 { _ = "STUB: not implemented"; return 0 }

//revive:disable-line:add-constant

func (r *tabButtonRenderer) padding() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

var (
	_ fyne.Widget       = (*tabCloseButton)(nil)
	_ fyne.Tappable     = (*tabCloseButton)(nil)
	_ desktop.Hoverable = (*tabCloseButton)(nil)
)

type tabCloseButton struct {
	widget.BaseWidget
	parent   *tabButton
	hovered  bool
	onTapped func()
}

func (b *tabCloseButton) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (b *tabCloseButton) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (b *tabCloseButton) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *tabCloseButton) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (b *tabCloseButton) MouseOut() { _ = "STUB: not implemented"; return }

func (b *tabCloseButton) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

type tabCloseButtonRenderer struct {
	button     *tabCloseButton
	background *canvas.Rectangle
	icon       *canvas.Image
	objects    []fyne.CanvasObject
}

func (r *tabCloseButtonRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *tabCloseButtonRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *tabCloseButtonRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *tabCloseButtonRenderer) Objects() []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func (r *tabCloseButtonRenderer) Refresh() { _ = "STUB: not implemented"; return }

func mismatchedTabItems(items []*TabItem) bool { _ = "STUB: not implemented"; return false }

func moreIcon(t baseTabs) fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }
