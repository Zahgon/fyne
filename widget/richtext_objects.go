package widget

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var (
	RichTextStyleBlockquote = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Italic: true},
	}

	RichTextStyleCodeBlock = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    false,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Monospace: true},
	}

	RichTextStyleCodeInline = RichTextStyle{
		ColorName:  theme.ColorNameForeground,
		Inline:     true,
		SizeName:   theme.SizeNameText,
		TextStyle:  fyne.TextStyle{Monospace: true},
		codeInline: true,
	}

	RichTextStyleEmphasis = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Italic: true},
	}

	RichTextStyleHeading = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameHeadingText,
		TextStyle: fyne.TextStyle{Bold: true},
	}

	RichTextStyleInline = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
	}

	RichTextStyleParagraph = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    false,
		SizeName:  theme.SizeNameText,
	}

	RichTextStylePassword = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
		concealed: true,
	}

	RichTextStyleStrong = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Bold: true},
	}

	RichTextStyleSubHeading = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameSubHeadingText,
		TextStyle: fyne.TextStyle{Bold: true},
	}
)

type HyperlinkSegment struct {
	Alignment fyne.TextAlign
	Text      string
	URL       *url.URL

	OnTapped func() `json:"-"`

	TextStyle fyne.TextStyle

	SizeName     fyne.ThemeSizeName
	quotingLevel int
}

func (h *HyperlinkSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (h *HyperlinkSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (h *HyperlinkSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (h *HyperlinkSegment) Update(o fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (h *HyperlinkSegment) Select(begin, end fyne.Position) { _ = "STUB: not implemented"; return }

func (h *HyperlinkSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (h *HyperlinkSegment) Unselect() { _ = "STUB: not implemented"; return }

type ImageSegment struct {
	Source fyne.URI
	Title  string

	Alignment fyne.TextAlign
}

func (i *ImageSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (i *ImageSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (i *ImageSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (i *ImageSegment) Update(o fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (i *ImageSegment) Select(begin, end fyne.Position) { _ = "STUB: not implemented"; return }

func (i *ImageSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (i *ImageSegment) Unselect() { _ = "STUB: not implemented"; return }

type ListSegment struct {
	Items   []RichTextSegment
	Ordered bool

	startIndex       int
	indentationLevel int
	quotingLevel     int
}

func (l *ListSegment) SetStartNumber(s int) { _ = "STUB: not implemented"; return }

func (l *ListSegment) StartNumber() int { _ = "STUB: not implemented"; return 0 }

func (l *ListSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (l *ListSegment) Segments() []RichTextSegment { _ = "STUB: not implemented"; return nil }

func (l *ListSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (l *ListSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (l *ListSegment) Update(fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (l *ListSegment) Select(_, _ fyne.Position) { _ = "STUB: not implemented"; return }

func (l *ListSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (l *ListSegment) Unselect() { _ = "STUB: not implemented"; return }

type ParagraphSegment struct {
	Texts []RichTextSegment
}

func (p *ParagraphSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (p *ParagraphSegment) Segments() []RichTextSegment { _ = "STUB: not implemented"; return nil }

func (p *ParagraphSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (p *ParagraphSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (p *ParagraphSegment) Update(fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (p *ParagraphSegment) Select(_, _ fyne.Position) { _ = "STUB: not implemented"; return }

func (p *ParagraphSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (p *ParagraphSegment) Unselect() { _ = "STUB: not implemented"; return }

type SeparatorSegment struct {
	_ bool
}

func (s *SeparatorSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (s *SeparatorSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (s *SeparatorSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (s *SeparatorSegment) Update(fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (s *SeparatorSegment) Select(_, _ fyne.Position) { _ = "STUB: not implemented"; return }

func (s *SeparatorSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (s *SeparatorSegment) Unselect() { _ = "STUB: not implemented"; return }

type CodeBlockSegment struct {
	Text         string
	quotingLevel int
}

func (c *CodeBlockSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (c *CodeBlockSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (c *CodeBlockSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (c *CodeBlockSegment) Update(o fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (c *CodeBlockSegment) Select(_, _ fyne.Position) { _ = "STUB: not implemented"; return }

func (c *CodeBlockSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (c *CodeBlockSegment) Unselect() { _ = "STUB: not implemented"; return }

type richCodeBlock struct {
	BaseWidget
	text  string
	bg    *canvas.Rectangle
	label *Label
}

func newRichCodeBlock(text string) *richCodeBlock { _ = "STUB: not implemented"; return nil }

func (c *richCodeBlock) setText(text string) { _ = "STUB: not implemented"; return }

func (c *richCodeBlock) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

type richCodeBlockLayout struct{}

func (l *richCodeBlockLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (l *richCodeBlockLayout) Layout(objects []fyne.CanvasObject, s fyne.Size) {
	_ = "STUB: not implemented"
	return
}

type CheckBoxSegment struct {
	Checked bool
	Text    string
}

func (c *CheckBoxSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (c *CheckBoxSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (c *CheckBoxSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (c *CheckBoxSegment) Update(fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (c *CheckBoxSegment) Select(_, _ fyne.Position) { _ = "STUB: not implemented"; return }

func (c *CheckBoxSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (c *CheckBoxSegment) Unselect() { _ = "STUB: not implemented"; return }

type TableSegment struct {
	Headers [][]RichTextSegment

	Rows       [][][]RichTextSegment
	Alignments []fyne.TextAlign
}

func (t *TableSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (t *TableSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (t *TableSegment) columns() int { _ = "STUB: not implemented"; return 0 }

func (t *TableSegment) alignFor(col int) fyne.TextAlign {
	_ = "STUB: not implemented"
	return *new(fyne.TextAlign)
}

func (t *TableSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (t *TableSegment) Update(fyne.CanvasObject) { _ = "STUB: not implemented"; return }

func (t *TableSegment) Select(_, _ fyne.Position) { _ = "STUB: not implemented"; return }

func (t *TableSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (t *TableSegment) Unselect() { _ = "STUB: not implemented"; return }

func newTableCell(segs []RichTextSegment, align fyne.TextAlign, header bool) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

type tableSegmentLayout struct {
	cols int
}

func (l *tableSegmentLayout) measure(objects []fyne.CanvasObject) (colWidths, rowHeights []float32) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *tableSegmentLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (l *tableSegmentLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

type RichTextStyle struct {
	Alignment fyne.TextAlign
	ColorName fyne.ThemeColorName
	Inline    bool
	SizeName  fyne.ThemeSizeName
	TextStyle fyne.TextStyle

	QuotingDepth int

	concealed bool

	codeInline bool
}

type RichTextSegment interface {
	Inline() bool
	Textual() string
	Update(fyne.CanvasObject)
	Visual() fyne.CanvasObject

	Select(pos1, pos2 fyne.Position)
	SelectedText() string
	Unselect()
}

type TextSegment struct {
	Style RichTextStyle
	Text  string

	parent *RichText
}

func (t *TextSegment) Inline() bool { _ = "STUB: not implemented"; return false }

func (t *TextSegment) Textual() string { _ = "STUB: not implemented"; return "" }

func (t *TextSegment) Visual() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (t *TextSegment) Update(o fyne.CanvasObject) { _ = "STUB: not implemented"; return }

type codeInlineLayout struct{}

func (codeInlineLayout) MinSize(o []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (codeInlineLayout) Layout(o []fyne.CanvasObject, _ fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (t *TextSegment) Select(begin, end fyne.Position) { _ = "STUB: not implemented"; return }

func (t *TextSegment) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (t *TextSegment) Unselect() { _ = "STUB: not implemented"; return }

func (t *TextSegment) color() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (t *TextSegment) size() float32 { _ = "STUB: not implemented"; return 0 }

type richImage struct {
	BaseWidget
	align  fyne.TextAlign
	img    *canvas.Image
	oldMin fyne.Size
	layout *fyne.Container
	min    fyne.Size
}

func newRichImage(u fyne.URI, align fyne.TextAlign) *richImage {
	_ = "STUB: not implemented"
	return nil
}

func (r *richImage) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (r *richImage) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *richImage) setAlign(a fyne.TextAlign) { _ = "STUB: not implemented"; return }

type richImageLayout struct {
	r *richImage
}

func (r *richImageLayout) Layout(_ []fyne.CanvasObject, s fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (r *richImageLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

type unpadTextWidgetLayout struct {
	parent fyne.Widget
}

func (u *unpadTextWidgetLayout) Layout(o []fyne.CanvasObject, s fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (u *unpadTextWidgetLayout) MinSize(o []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
