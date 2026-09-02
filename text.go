package fyne

type TextAlign int

const (
	TextAlignLeading TextAlign = iota

	TextAlignCenter

	TextAlignTrailing
)

type TextTruncation int

const (
	TextTruncateOff TextTruncation = iota

	TextTruncateClip

	TextTruncateEllipsis
)

type TextWrap int

const (
	TextWrapOff TextWrap = iota

	TextTruncate

	TextWrapBreak

	TextWrapWord
)

type TextStyle struct {
	Bold      bool
	Italic    bool
	Monospace bool

	Symbol bool

	TabWidth int

	Underline bool

	Strikethrough bool
}

func MeasureText(text string, size float32, style TextStyle) Size {
	_ = "STUB: not implemented"
	return *new(Size)
}
