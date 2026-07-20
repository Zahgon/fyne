package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/async"
	"fyne.io/fyne/v2/internal/widget"
)

var (
	_ fyne.Widget    = (*GridWrap)(nil)
	_ fyne.Focusable = (*GridWrap)(nil)
)

type GridWrapItemID = int

type GridWrap struct {
	BaseWidget

	Length func() int `json:"-"`

	CreateItem func() fyne.CanvasObject `json:"-"`

	UpdateItem func(id GridWrapItemID, item fyne.CanvasObject) `json:"-"`

	OnSelected func(id GridWrapItemID) `json:"-"`

	OnUnselected func(id GridWrapItemID) `json:"-"`

	OnHighlighted func(id GridWrapItemID) `json:"-"`

	currentHighlight ListItemID
	focused          bool
	scroller         *widget.Scroll
	selected         []GridWrapItemID
	itemMin          fyne.Size
	offsetY          float32
	offsetUpdated    func(fyne.Position)
	colCountCache    int
	minSizeCache     fyne.Size
}

func NewGridWrap(length func() int, createItem func() fyne.CanvasObject, updateItem func(GridWrapItemID, fyne.CanvasObject)) *GridWrap {
	_ = "STUB: not implemented"
	return nil
}

func NewGridWrapWithData(data binding.DataList, createItem func() fyne.CanvasObject, updateItem func(binding.DataItem, fyne.CanvasObject)) *GridWrap {
	_ = "STUB: not implemented"
	return nil
}

func (l *GridWrap) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (l *GridWrap) FocusGained() { _ = "STUB: not implemented"; return }

func (l *GridWrap) FocusLost() { _ = "STUB: not implemented"; return }

func (l *GridWrap) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *GridWrap) scrollTo(id GridWrapItemID) { _ = "STUB: not implemented"; return }

func (l *GridWrap) RefreshItem(id GridWrapItemID) { _ = "STUB: not implemented"; return }

func (l *GridWrap) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func (l *GridWrap) Highlight(id GridWrapItemID) { _ = "STUB: not implemented"; return }

func (l *GridWrap) Select(id GridWrapItemID) { _ = "STUB: not implemented"; return }

func (l *GridWrap) ScrollTo(id GridWrapItemID) { _ = "STUB: not implemented"; return }

func (l *GridWrap) ScrollToBottom() { _ = "STUB: not implemented"; return }

func (l *GridWrap) ScrollToTop() { _ = "STUB: not implemented"; return }

func (l *GridWrap) ScrollToOffset(offset float32) { _ = "STUB: not implemented"; return }

func (l *GridWrap) TypedKey(event *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (l *GridWrap) TypedRune(_ rune) { _ = "STUB: not implemented"; return }

func (l *GridWrap) GetScrollOffset() float32 { _ = "STUB: not implemented"; return 0 }

func (l *GridWrap) Unselect(id GridWrapItemID) { _ = "STUB: not implemented"; return }

func (l *GridWrap) UnselectAll() { _ = "STUB: not implemented"; return }

func (l *GridWrap) Refresh() { _ = "STUB: not implemented"; return }

func (l *GridWrap) contentMinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

var _ fyne.WidgetRenderer = (*gridWrapRenderer)(nil)

type gridWrapRenderer struct {
	objects []fyne.CanvasObject

	list     *GridWrap
	scroller *widget.Scroll
	layout   *fyne.Container
}

func newGridWrapRenderer(objects []fyne.CanvasObject, l *GridWrap, scroller *widget.Scroll, layout *fyne.Container) *gridWrapRenderer {
	_ = "STUB: not implemented"
	return nil
}

func (l *gridWrapRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (l *gridWrapRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *gridWrapRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (l *gridWrapRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (l *gridWrapRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

var (
	_ fyne.Widget       = (*gridWrapItem)(nil)
	_ fyne.Tappable     = (*gridWrapItem)(nil)
	_ desktop.Hoverable = (*gridWrapItem)(nil)
)

type gridWrapItem struct {
	BaseWidget

	onTapped          func()
	onHovered         func()
	background        *canvas.Rectangle
	child             fyne.CanvasObject
	hovered, selected bool
}

func newGridWrapItem(child fyne.CanvasObject, tapped func()) *gridWrapItem {
	_ = "STUB: not implemented"
	return nil
}

func (gw *gridWrapItem) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (gw *gridWrapItem) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (gw *gridWrapItem) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (gw *gridWrapItem) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (gw *gridWrapItem) MouseOut() { _ = "STUB: not implemented"; return }

func (gw *gridWrapItem) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*gridWrapItemRenderer)(nil)

type gridWrapItemRenderer struct {
	widget.BaseRenderer

	item *gridWrapItem
}

func (gw *gridWrapItemRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (gw *gridWrapItemRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (gw *gridWrapItemRenderer) Refresh() { _ = "STUB: not implemented"; return }

var _ fyne.Layout = (*gridWrapLayout)(nil)

type gridItemAndID struct {
	item *gridWrapItem
	id   GridWrapItemID
}

type gridWrapLayout struct {
	gw *GridWrap

	itemPool   async.Pool[fyne.CanvasObject]
	visible    []gridItemAndID
	wasVisible []gridItemAndID
}

func newGridWrapLayout(gw *GridWrap) fyne.Layout {
	_ = "STUB: not implemented"
	return *new(fyne.Layout)
}

func (l *gridWrapLayout) Layout(_ []fyne.CanvasObject, _ fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (l *gridWrapLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (l *gridWrapLayout) getItem() *gridWrapItem { _ = "STUB: not implemented"; return nil }

func (l *gridWrapLayout) offsetUpdated(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (l *gridWrapLayout) setupGridItem(li *gridWrapItem, id GridWrapItemID, focus bool) {
	_ = "STUB: not implemented"
	return
}

func (l *GridWrap) ColumnCount() int { _ = "STUB: not implemented"; return 0 }

func (l *gridWrapLayout) updateGrid(newOnly bool) { _ = "STUB: not implemented"; return }

func (l *gridWrapLayout) searchVisible(visible []gridItemAndID, id GridWrapItemID) (*gridWrapItem, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (l *gridWrapLayout) nilOldSliceData(objs []fyne.CanvasObject, len, oldLen int) {
	_ = "STUB: not implemented"
	return
}
