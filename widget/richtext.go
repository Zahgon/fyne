package widget

import (
	"golang.org/x/image/math/fixed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
)

const passwordChar = "•"

var _ fyne.Widget = (*RichText)(nil)

type RichText struct {
	BaseWidget
	Segments []RichTextSegment
	Wrapping fyne.TextWrap
	Scroll   fyne.ScrollDirection

	Truncation fyne.TextTruncation

	inset     fyne.Size
	rowBounds []rowBoundary
	scr       *widget.Scroll
	prop      *canvas.Rectangle

	visualCache    map[RichTextSegment]visualCacheEntry
	visualCacheGen int64
	minCache       fyne.Size
}

type visualCacheEntry struct {
	gen int64
	obj []fyne.CanvasObject
}

func NewRichText(segments ...RichTextSegment) *RichText { _ = "STUB: not implemented"; return nil }

func NewRichTextWithText(text string) *RichText { _ = "STUB: not implemented"; return nil }

func (t *RichText) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *RichText) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *RichText) Refresh() { _ = "STUB: not implemented"; return }

func (t *RichText) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (t *RichText) String() string { _ = "STUB: not implemented"; return "" }

func (*RichText) charMinSize(concealed bool, style fyne.TextStyle, textSize float32) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (t *RichText) deleteFromTo(lowBound int, highBound int) []rune {
	_ = "STUB: not implemented"
	return nil
}

func (t *RichText) cachedSegmentVisual(seg RichTextSegment, offset int) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (t *RichText) cleanVisualCache() { _ = "STUB: not implemented"; return }

func (t *RichText) insertAt(pos int, runes []rune) { _ = "STUB: not implemented"; return }

func (t *RichText) len() int { _ = "STUB: not implemented"; return 0 }

func (t *RichText) lineSizeToColumn(col, row int, textSize, innerPad float32) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (t *RichText) row(row int) []rune { _ = "STUB: not implemented"; return nil }

func (t *RichText) rowBoundary(row int) *rowBoundary { _ = "STUB: not implemented"; return nil }

func (t *RichText) rowLength(row int) int { _ = "STUB: not implemented"; return 0 }

func (t *RichText) rows() int { _ = "STUB: not implemented"; return 0 }

func (t *RichText) updateRowBounds() { _ = "STUB: not implemented"; return }

type RichTextBlock interface {
	Segments() []RichTextSegment
}

type textRenderer struct {
	widget.BaseRenderer
	obj *RichText
}

func codeInlineText(obj fyne.CanvasObject) (*canvas.Text, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *textRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *textRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *textRenderer) calculateMin(bounds []rowBoundary, wrap fyne.TextWrap, objs []fyne.CanvasObject,
	charMinSize fyne.Size, th fyne.Theme,
) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *textRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *textRenderer) associateSiblings(hl *Hyperlink, hlSeg *HyperlinkSegment, reuse int) {
	_ = "STUB: not implemented"
	return
}

func (r *textRenderer) layoutRow(texts []fyne.CanvasObject, align fyne.TextAlign, xPos, yPos, lineWidth float32) (x, height float32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func isEmptyScroll(o *widget.Scroll) bool { _ = "STUB: not implemented"; return false }

func howManyRunesFit(runes []rune, availableWidth float32, charWidth float32, measurer func([]rune) fyne.Size) int {
	_ = "STUB: not implemented"
	return 0
}

func concealed(seg RichTextSegment) bool { _ = "STUB: not implemented"; return false }

func ellipsisPriorBound(bounds []rowBoundary, trunc fyne.TextTruncation, width float32, charWidth float32, measurer func([]rune) fyne.Size) []rowBoundary {
	_ = "STUB: not implemented"
	return nil
}

//revive:disable-line:add-constant

func findSpaceIndex(text []rune, curIndex int) int { _ = "STUB: not implemented"; return 0 }

func float32ToFixed266(f float32) fixed.Int26_6 {
	_ = "STUB: not implemented"
	return *new(fixed.Int26_6)
}

func lineBounds(t *RichText, seg RichTextSegment, firstWidth float32, maxSize fyne.Size, measurer func([]rune) fyne.Size) ([]rowBoundary, float32) {
	_ = "STUB: not implemented"
	return nil, 0
}

func wrapBreakLines(seg RichTextSegment, trunc fyne.TextTruncation, measureWidth float32, maxSize fyne.Size, measurer func([]rune) fyne.Size, lines []rowBoundary) ([]rowBoundary, float32) {
	_ = "STUB: not implemented"
	return nil, 0
}

func wrapWordLines(seg RichTextSegment, trunc fyne.TextTruncation, measureWidth float32, maxSize fyne.Size, measurer func([]rune) fyne.Size, lines []rowBoundary) ([]rowBoundary, float32) {
	_ = "STUB: not implemented"
	return nil, 0
}

func truncateLines(t *RichText, seg RichTextSegment, trunc fyne.TextTruncation, measureWidth float32, measurer func([]rune) fyne.Size, lines []rowBoundary) ([]rowBoundary, float32) {
	_ = "STUB: not implemented"
	return nil, 0
}

//revive:disable-line:add-constant -- TODO: clarify whether we want to define a common letter constant for approximate character sizes

func setAlign(obj fyne.CanvasObject, align fyne.TextAlign) { _ = "STUB: not implemented"; return }

func rowPaddingAndAlign(bound rowBoundary, lineSpacing float32, currentAlign fyne.TextAlign) (float32, fyne.TextAlign) {
	_ = "STUB: not implemented"
	return 0, *new(fyne.TextAlign)
}

func splitLines(seg RichTextSegment) []rowBoundary { _ = "STUB: not implemented"; return nil }

func truncateLimit(s string, text *canvas.Text, limit int, ellipsis []rune) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type rowBoundary struct {
	segments          []RichTextSegment
	firstSegmentReuse int
	begin, end        int
	ellipsis          bool
	indent            float32
}
