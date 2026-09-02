package widget

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/internal/async"
	"fyne.io/fyne/v2/internal/widget"
)

const (
	columnLetterCount = 26
	noCellMatch       = math.MaxInt
)

var (
	allTableCellsID = TableCellID{-1, -1}

	onlyNewTableCellsID = TableCellID{-2, -2}
)

var (
	_ desktop.Cursorable = (*Table)(nil)
	_ fyne.Draggable     = (*Table)(nil)
	_ fyne.Focusable     = (*Table)(nil)
	_ desktop.Hoverable  = (*Table)(nil)
	_ fyne.Tappable      = (*Table)(nil)
	_ fyne.Widget        = (*Table)(nil)
)

type TableCellID struct {
	Row int
	Col int
}

type Table struct {
	BaseWidget

	Length       func() (rows int, cols int)                      `json:"-"`
	CreateCell   func() fyne.CanvasObject                         `json:"-"`
	UpdateCell   func(id TableCellID, template fyne.CanvasObject) `json:"-"`
	OnSelected   func(id TableCellID)                             `json:"-"`
	OnUnselected func(id TableCellID)                             `json:"-"`

	ShowHeaderRow bool

	ShowHeaderColumn bool

	CreateHeader func() fyne.CanvasObject `json:"-"`

	UpdateHeader func(id TableCellID, template fyne.CanvasObject) `json:"-"`

	StickyRowCount int

	StickyColumnCount int

	HideSeparators bool

	OnHighlighted func(id TableCellID) `json:"-"`

	currentHighlight          TableCellID
	focused                   bool
	selectedCell, hoveredCell *TableCellID
	cells                     *tableCells
	columnWidths, rowHeights  map[int]float32
	moveCallback              func()
	offset                    fyne.Position
	content                   *widget.Scroll

	cellSize, headerSize                                         fyne.Size
	stuckXOff, stuckYOff, stuckWidth, stuckHeight, dragStartSize float32
	top, left, corner, dividerLayer                              *clip
	hoverHeaderRow, hoverHeaderCol, dragCol, dragRow             int
	dragStartPos                                                 fyne.Position
}

func NewTable(length func() (rows int, cols int), create func() fyne.CanvasObject, update func(TableCellID, fyne.CanvasObject)) *Table {
	_ = "STUB: not implemented"
	return nil
}

func NewTableWithHeaders(length func() (rows int, cols int), create func() fyne.CanvasObject, update func(TableCellID, fyne.CanvasObject)) *Table {
	_ = "STUB: not implemented"
	return nil
}

func (t *Table) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *Table) Cursor() desktop.Cursor { _ = "STUB: not implemented"; return *new(desktop.Cursor) }

func (t *Table) Dragged(e *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (t *Table) DragEnd() { _ = "STUB: not implemented"; return }

func (t *Table) FocusGained() { _ = "STUB: not implemented"; return }

func (t *Table) FocusLost() { _ = "STUB: not implemented"; return }

func (t *Table) MouseIn(ev *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (t *Table) MouseDown(e *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (t *Table) MouseMoved(ev *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (t *Table) MouseOut() { _ = "STUB: not implemented"; return }

func (*Table) MouseUp(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (t *Table) RefreshItem(id TableCellID) { _ = "STUB: not implemented"; return }

func (t *Table) Select(id TableCellID) { _ = "STUB: not implemented"; return }

func (t *Table) SetColumnWidth(id int, width float32) { _ = "STUB: not implemented"; return }

func (t *Table) SetRowHeight(id int, height float32) { _ = "STUB: not implemented"; return }

func (t *Table) TouchDown(e *mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (*Table) TouchUp(*mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (*Table) TouchCancel(*mobile.TouchEvent) { _ = "STUB: not implemented"; return }

func (t *Table) TypedKey(event *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (*Table) TypedRune(_ rune) { _ = "STUB: not implemented"; return }

func (t *Table) Unselect(id TableCellID) { _ = "STUB: not implemented"; return }

func (t *Table) UnselectAll() { _ = "STUB: not implemented"; return }

func (t *Table) Highlight(id TableCellID) { _ = "STUB: not implemented"; return }

func (t *Table) ScrollTo(id TableCellID) { _ = "STUB: not implemented"; return }

func (t *Table) ScrollToBottom() { _ = "STUB: not implemented"; return }

func (t *Table) ScrollToLeading() { _ = "STUB: not implemented"; return }

func (t *Table) ScrollToOffset(off fyne.Position) { _ = "STUB: not implemented"; return }

func (t *Table) ScrollToTop() { _ = "STUB: not implemented"; return }

func (t *Table) ScrollToTrailing() { _ = "STUB: not implemented"; return }

func (t *Table) Tapped(e *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (t *Table) columnAt(pos fyne.Position) int { _ = "STUB: not implemented"; return 0 }

func (t *Table) createHeader() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (t *Table) findX(col int) (cellX float32, cellWidth float32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (t *Table) findY(row int) (cellY float32, cellHeight float32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (t *Table) finishScroll() { _ = "STUB: not implemented"; return }

func (t *Table) hoverAt(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (t *Table) hoverOut() { _ = "STUB: not implemented"; return }

func (t *Table) rowAt(pos fyne.Position) int { _ = "STUB: not implemented"; return 0 }

func (t *Table) tapped(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (t *Table) templateSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *Table) updateHeader(id TableCellID, o fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) stickyColumnWidths(colWidth float32, cols int) (visible []float32) {
	_ = "STUB: not implemented"
	return nil
}

func (t *Table) visibleColumnWidths(colWidth float32, cols int) (visible map[int]float32, offX float32, minCol, maxCol int) {
	_ = "STUB: not implemented"
	return nil, 0, 0, 0
}

func (t *Table) stickyRowHeights(rowHeight float32, rows int) (visible []float32) {
	_ = "STUB: not implemented"
	return nil
}

func (t *Table) visibleRowHeights(rowHeight float32, rows int) (visible map[int]float32, offY float32, minRow, maxRow int) {
	_ = "STUB: not implemented"
	return nil, 0, 0, 0
}

var _ fyne.WidgetRenderer = (*tableRenderer)(nil)

type tableRenderer struct {
	widget.BaseRenderer
	t *Table
}

func (t *tableRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (t *tableRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *tableRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (t *tableRenderer) calculateHeaderSizes(th fyne.Theme) { _ = "STUB: not implemented"; return }

var _ fyne.Widget = (*tableCells)(nil)

type tableCells struct {
	BaseWidget
	t *Table

	nextRefreshCellsID TableCellID
}

func newTableCells(t *Table) *tableCells { _ = "STUB: not implemented"; return nil }

func (c *tableCells) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *tableCells) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func (c *tableCells) refreshForID(id TableCellID) { _ = "STUB: not implemented"; return }

func (c *tableCells) Refresh() { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*tableCellsRenderer)(nil)

type tableCellsRenderer struct {
	widget.BaseRenderer

	cells            *tableCells
	pool, headerPool async.Pool[fyne.CanvasObject]
	visible, headers map[TableCellID]fyne.CanvasObject
	hover, marker    *canvas.Rectangle
	dividers         []fyne.CanvasObject

	headColBG, headRowBG, headRowStickyBG, headColStickyBG *canvas.Rectangle
}

func (r *tableCellsRenderer) Layout(fyne.Size) { _ = "STUB: not implemented"; return }

func (r *tableCellsRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *tableCellsRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *tableCellsRenderer) refreshForID(toDraw TableCellID) { _ = "STUB: not implemented"; return }

func (r *tableCellsRenderer) updateCells(toDraw TableCellID, visible, wasVisible map[TableCellID]fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return
}

func (r *tableCellsRenderer) moveIndicators() { _ = "STUB: not implemented"; return }

func (r *tableCellsRenderer) moveMarker(marker fyne.CanvasObject, row, col int, offX, offY float32, minCol, minRow int, widths, heights map[int]float32) {
	_ = "STUB: not implemented"
	return
}

func (r *tableCellsRenderer) refreshHeaders(visibleRowHeights, visibleColWidths map[int]float32, offX, offY float32,
	startRow, maxRow, startCol, maxCol int, separatorThickness float32, th fyne.Theme, v fyne.ThemeVariant,
) []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

type clip struct {
	widget.Scroll

	t *Table
}

func newClip(t *Table, o fyne.CanvasObject) *clip { _ = "STUB: not implemented"; return nil }

func (c *clip) DragEnd() { _ = "STUB: not implemented"; return }

func (c *clip) Dragged(e *fyne.DragEvent) { _ = "STUB: not implemented"; return }
