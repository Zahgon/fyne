package test

import "fyne.io/fyne/v2"

type clipboard struct {
	content string
}

func (c *clipboard) Content() string { _ = "STUB: not implemented"; return "" }

func (c *clipboard) SetContent(content string) { _ = "STUB: not implemented"; return }

func NewClipboard() fyne.Clipboard { _ = "STUB: not implemented"; return *new(fyne.Clipboard) }
