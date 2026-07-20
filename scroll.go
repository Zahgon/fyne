package fyne

type ScrollDirection int

const (
	ScrollBoth ScrollDirection = iota

	ScrollHorizontalOnly

	ScrollVerticalOnly

	ScrollNone
)
