package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/async"
	"fyne.io/fyne/v2/internal/widget"
)

const (
	textAreaSpaceSymbol   = '·'
	textAreaTabSymbol     = '→'
	textAreaNewLineSymbol = '↵'
)

var (
	TextGridStyleDefault TextGridStyle

	TextGridStyleWhitespace TextGridStyle
)

type TextGridCell struct {
	Rune  rune
	Style TextGridStyle
}

type TextGridRow struct {
	Cells []TextGridCell
	Style TextGridStyle
}

type TextGridStyle interface {
	Style() fyne.TextStyle
	TextColor() color.Color
	BackgroundColor() color.Color
}

type CustomTextGridStyle struct {
	TextStyle        fyne.TextStyle
	FGColor, BGColor color.Color
}

func (c *CustomTextGridStyle) TextColor() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (c *CustomTextGridStyle) BackgroundColor() color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (c *CustomTextGridStyle) Style() fyne.TextStyle {
	_ = "STUB: not implemented"
	return *new(fyne.TextStyle)
}

type TextGrid struct {
	BaseWidget
	Rows []TextGridRow

	scroll  *widget.Scroll
	content *textGridContent

	ShowLineNumbers bool
	ShowWhitespace  bool
	TabWidth        int

	Scroll fyne.ScrollDirection
}

func (t *TextGrid) Append(text string) { _ = "STUB: not implemented"; return }

func (t *TextGrid) CursorLocationForPosition(p fyne.Position) (row, col int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (t *TextGrid) ScrollToTop() { _ = "STUB: not implemented"; return }

func (t *TextGrid) ScrollToBottom() { _ = "STUB: not implemented"; return }

func (t *TextGrid) PositionForCursorLocation(row, col int) fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func (t *TextGrid) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *TextGrid) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (t *TextGrid) SetText(text string) { _ = "STUB: not implemented"; return }

func (t *TextGrid) Text() string { _ = "STUB: not implemented"; return "" }

func (t *TextGrid) Row(row int) TextGridRow { _ = "STUB: not implemented"; return *new(TextGridRow) }

func (t *TextGrid) RowText(row int) string { _ = "STUB: not implemented"; return "" }

func (t *TextGrid) SetRow(row int, content TextGridRow) { _ = "STUB: not implemented"; return }

func (t *TextGrid) SetRowStyle(row int, style TextGridStyle) { _ = "STUB: not implemented"; return }

func (t *TextGrid) SetCell(row, col int, cell TextGridCell) { _ = "STUB: not implemented"; return }

func (t *TextGrid) SetRune(row, col int, r rune) { _ = "STUB: not implemented"; return }

func (t *TextGrid) SetStyle(row, col int, style TextGridStyle) { _ = "STUB: not implemented"; return }

func (t *TextGrid) SetStyleRange(startRow, startCol, endRow, endCol int, style TextGridStyle) {
	_ = "STUB: not implemented"
	return
}

func (t *TextGrid) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *TextGrid) ensureCells(row, col int) { _ = "STUB: not implemented"; return }

func (t *TextGrid) parseRows(text string) []TextGridRow { _ = "STUB: not implemented"; return nil }

func (t *TextGrid) refreshCell(row, col int) { _ = "STUB: not implemented"; return }

func NewTextGrid() *TextGrid { _ = "STUB: not implemented"; return nil }

func NewTextGridFromString(content string) *TextGrid { _ = "STUB: not implemented"; return nil }

func nextTab(column int, tabWidth int) int { _ = "STUB: not implemented"; return 0 }

type textGridRenderer struct {
	widget.BaseRenderer

	text   *textGridContent
	scroll *widget.Scroll
}

func (t *textGridRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (t *textGridRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *textGridRenderer) Refresh() { _ = "STUB: not implemented"; return }

type textGridContent struct {
	BaseWidget
	text *TextGrid

	rows     int
	cellSize fyne.Size

	visible []fyne.CanvasObject
}

func newTextGridContent(t *TextGrid) *textGridContent { _ = "STUB: not implemented"; return nil }

func (t *textGridContent) lineNumberWidth() int { _ = "STUB: not implemented"; return 0 }

func (t *textGridContent) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *textGridContent) refreshCell(row, col int) { _ = "STUB: not implemented"; return }

type textGridContentRenderer struct {
	text     *textGridContent
	itemPool async.Pool[*textGridRow]
}

func (t *textGridContentRenderer) updateGridSize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (*textGridContentRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (t *textGridContentRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (t *textGridContentRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (t *textGridContentRenderer) Objects() []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func (t *textGridContentRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (t *textGridContentRenderer) addRowsIfRequired() { _ = "STUB: not implemented"; return }

func (t *textGridContentRenderer) updateCellSize() { _ = "STUB: not implemented"; return }

type textGridRow struct {
	BaseWidget
	text *textGridContent

	objects []fyne.CanvasObject
	row     int
	cols    int

	cachedFGColor  color.Color
	cachedTextSize float32
}

func newTextGridRow(t *textGridContent, row int) *textGridRow {
	_ = "STUB: not implemented"
	return nil
}

func (t *textGridRow) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *textGridRow) setRow(row int) { _ = "STUB: not implemented"; return }

func (t *textGridRow) appendTextCell(str rune) { _ = "STUB: not implemented"; return }

func (t *textGridRow) refreshCell(col int) { _ = "STUB: not implemented"; return }

func (t *textGridRow) setCellRune(str rune, pos int, style, rowStyle TextGridStyle) {
	_ = "STUB: not implemented"
	return
}

func (t *textGridRow) addCellsIfRequired() { _ = "STUB: not implemented"; return }

func (t *textGridRow) refreshCells() { _ = "STUB: not implemented"; return }

func (t *TextGrid) tabWidth() int { _ = "STUB: not implemented"; return 0 }

func (t *textGridRow) updateGridSize(size fyne.Size) { _ = "STUB: not implemented"; return }

type textGridRowRenderer struct {
	obj *textGridRow
}

func (t *textGridRowRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (t *textGridRowRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (t *textGridRowRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (*textGridRowRenderer) ApplyTheme() { _ = "STUB: not implemented"; return }

func (t *textGridRowRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (*textGridRowRenderer) Destroy() { _ = "STUB: not implemented"; return }
