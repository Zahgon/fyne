package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/async"
	"fyne.io/fyne/v2/internal/widget"
)

type ListItemID = int

var (
	_ fyne.Widget    = (*List)(nil)
	_ fyne.Focusable = (*List)(nil)
)

type listBind struct {
	listener annotatedListener

	oldLength func() int                                  `json:"-"`
	oldUpdate func(id ListItemID, item fyne.CanvasObject) `json:"-"`
}

type List struct {
	BaseWidget

	Length func() int `json:"-"`

	CreateItem func() fyne.CanvasObject `json:"-"`

	UpdateItem func(id ListItemID, item fyne.CanvasObject) `json:"-"`

	OnSelected func(id ListItemID) `json:"-"`

	OnUnselected func(id ListItemID) `json:"-"`

	HideSeparators bool

	OnHighlighted func(id ListItemID) `json:"-"`

	currentHighlight ListItemID
	focused          bool
	scroller         *widget.Scroll
	selected         []ListItemID
	itemMin          fyne.Size
	itemHeights      map[ListItemID]float32
	offsetY          float32
	offsetUpdated    func(fyne.Position)
	minSizeCache     fyne.Size

	lastBind *listBind
}

func NewList(length func() int, createItem func() fyne.CanvasObject, updateItem func(ListItemID, fyne.CanvasObject)) *List {
	_ = "STUB: not implemented"
	return nil
}

func NewListWithData(data binding.DataList, createItem func() fyne.CanvasObject, updateItem func(binding.DataItem, fyne.CanvasObject)) *List {
	_ = "STUB: not implemented"
	return nil
}

func (l *List) Bind(data binding.DataList, update func(di binding.DataItem, o fyne.CanvasObject)) {
	_ = "STUB: not implemented"
	return
}

func (l *List) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (l *List) FocusGained() { _ = "STUB: not implemented"; return }

func (l *List) FocusLost() { _ = "STUB: not implemented"; return }

func (l *List) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *List) RefreshItem(id ListItemID) { _ = "STUB: not implemented"; return }

func (l *List) SetItemHeight(id ListItemID, height float32) { _ = "STUB: not implemented"; return }

func (l *List) Unbind() { _ = "STUB: not implemented"; return }

func (l *List) scrollTo(id ListItemID) { _ = "STUB: not implemented"; return }

func (l *List) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func (l *List) Highlight(id ListItemID) { _ = "STUB: not implemented"; return }

func (l *List) Select(id ListItemID) { _ = "STUB: not implemented"; return }

func (l *List) ScrollTo(id ListItemID) { _ = "STUB: not implemented"; return }

func (l *List) ScrollToBottom() { _ = "STUB: not implemented"; return }

func (l *List) ScrollToTop() { _ = "STUB: not implemented"; return }

func (l *List) ScrollToOffset(offset float32) { _ = "STUB: not implemented"; return }

func (l *List) GetScrollOffset() float32 { _ = "STUB: not implemented"; return 0 }

func (l *List) TypedKey(event *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (l *List) TypedRune(_ rune) { _ = "STUB: not implemented"; return }

func (l *List) Unselect(id ListItemID) { _ = "STUB: not implemented"; return }

func (l *List) UnselectAll() { _ = "STUB: not implemented"; return }

func (l *List) Refresh() { _ = "STUB: not implemented"; return }

func (l *List) contentMinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *listLayout) calculateVisibleRowHeights(itemHeight float32, length int, th fyne.Theme) (offY float32, minRow int) {
	_ = "STUB: not implemented"
	return 0, 0
}

var _ fyne.WidgetRenderer = (*listRenderer)(nil)

type listRenderer struct {
	widget.BaseRenderer

	list     *List
	scroller *widget.Scroll
	layout   *fyne.Container
}

func newListRenderer(objects []fyne.CanvasObject, l *List, scroller *widget.Scroll, layout *fyne.Container) *listRenderer {
	_ = "STUB: not implemented"
	return nil
}

func (l *listRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (l *listRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *listRenderer) Refresh() { _ = "STUB: not implemented"; return }

var (
	_ fyne.Widget       = (*listItem)(nil)
	_ fyne.Tappable     = (*listItem)(nil)
	_ desktop.Hoverable = (*listItem)(nil)
)

type listItem struct {
	BaseWidget

	onTapped          func()
	onHovered         func()
	background        *canvas.Rectangle
	child             fyne.CanvasObject
	hovered, selected bool
}

func newListItem(child fyne.CanvasObject, tapped func()) *listItem {
	_ = "STUB: not implemented"
	return nil
}

func (li *listItem) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (li *listItem) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (li *listItem) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (li *listItem) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (li *listItem) MouseOut() { _ = "STUB: not implemented"; return }

func (li *listItem) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*listItemRenderer)(nil)

type listItemRenderer struct {
	widget.BaseRenderer

	item *listItem
}

func (li *listItemRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (li *listItemRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (li *listItemRenderer) Refresh() { _ = "STUB: not implemented"; return }

var _ fyne.Layout = (*listLayout)(nil)

type listItemAndID struct {
	item *listItem
	id   ListItemID
}

type listLayout struct {
	list       *List
	separators []fyne.CanvasObject
	children   []fyne.CanvasObject

	itemPool          async.Pool[fyne.CanvasObject]
	visible           []listItemAndID
	wasVisible        []listItemAndID
	visibleRowHeights []float32
}

func newListLayout(list *List) fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func (l *listLayout) Layout([]fyne.CanvasObject, fyne.Size) { _ = "STUB: not implemented"; return }

func (l *listLayout) MinSize([]fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (l *listLayout) getItem() *listItem { _ = "STUB: not implemented"; return nil }

func (l *listLayout) offsetUpdated(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (l *listLayout) setupListItem(li *listItem, id ListItemID, focus bool) {
	_ = "STUB: not implemented"
	return
}

func (l *listLayout) updateList(newOnly bool) { _ = "STUB: not implemented"; return }

func (l *listLayout) updateSeparators() { _ = "STUB: not implemented"; return }

func (l *listLayout) searchVisible(visible []listItemAndID, id ListItemID) (*listItem, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (l *listLayout) nilOldSliceData(objs []fyne.CanvasObject, len, oldLen int) {
	_ = "STUB: not implemented"
	return
}

func createItemAndApplyThemeScope(f func() fyne.CanvasObject, scope fyne.Widget) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}
