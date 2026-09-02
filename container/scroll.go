package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/widget"
)

type Scroll = widget.Scroll

type ScrollDirection = fyne.ScrollDirection

const (
	ScrollBoth ScrollDirection = fyne.ScrollBoth

	ScrollHorizontalOnly = fyne.ScrollHorizontalOnly

	ScrollVerticalOnly = fyne.ScrollVerticalOnly

	ScrollNone = fyne.ScrollNone
)

func NewScroll(content fyne.CanvasObject) *Scroll { _ = "STUB: not implemented"; return nil }

func NewHScroll(content fyne.CanvasObject) *Scroll { _ = "STUB: not implemented"; return nil }

func NewVScroll(content fyne.CanvasObject) *Scroll { _ = "STUB: not implemented"; return nil }
