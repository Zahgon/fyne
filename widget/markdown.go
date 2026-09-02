package widget

import (
	"io"

	"github.com/yuin/goldmark/ast"
	ast2 "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"

	"fyne.io/fyne/v2"
)

func NewRichTextFromMarkdown(content string) *RichText { _ = "STUB: not implemented"; return nil }

func (t *RichText) ParseMarkdown(content string) { _ = "STUB: not implemented"; return }

func (t *RichText) AppendMarkdown(content string) { _ = "STUB: not implemented"; return }

type markdownRenderer []RichTextSegment

func (*markdownRenderer) AddOptions(...renderer.Option) { _ = "STUB: not implemented"; return }

func (m *markdownRenderer) Render(_ io.Writer, source []byte, n ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func renderNode(source []byte, n ast.Node, quotingDepth int, listDepth int) ([]RichTextSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renderTable(source []byte, n *ast2.Table) *TableSegment { _ = "STUB: not implemented"; return nil }

func renderCodeBlock(source []byte, n ast.Node, quotingDepth int) []RichTextSegment {
	_ = "STUB: not implemented"
	return nil
}

func tableAlignment(a ast2.Alignment) fyne.TextAlign {
	_ = "STUB: not implemented"
	return *new(fyne.TextAlign)
}

func renderChildren(source []byte, n ast.Node, quotingDepth int, listDepth int) ([]RichTextSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renderEmphasis(source []byte, n ast.Node, quotingDepth int, strength, listDepth int) ([]RichTextSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renderHeading(source []byte, n ast.Node, quotingDepth int, listDepth int) ([]RichTextSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func forceIntoText(source []byte, n ast.Node) string { _ = "STUB: not implemented"; return "" }

//revive:disable-line:add-constant

func parseMarkdown(content string) []RichTextSegment { _ = "STUB: not implemented"; return nil }

func decodeText(text string) string { _ = "STUB: not implemented"; return "" }
