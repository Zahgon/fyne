//go:build !ci && !android && !ios && !mobile && (wasm || test_web_driver)

package widget

import (
	"github.com/yuin/goldmark/ast"
)

func parseMarkdownImage(t *ast.Image) []RichTextSegment { _ = "STUB: not implemented"; return nil }
