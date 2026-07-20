package desktop

import "image"

type Cursor interface {
	Image() (image.Image, int, int)
}

type StandardCursor int

func (d StandardCursor) Image() (image.Image, int, int) {
	_ = "STUB: not implemented"
	return *new(image.Image), 0, 0
}

const (
	DefaultCursor StandardCursor = iota

	TextCursor

	CrosshairCursor

	PointerCursor

	HResizeCursor

	VResizeCursor

	NESWResizeCursor

	NWSEResizeCursor

	HiddenCursor
)

type Cursorable interface {
	Cursor() Cursor
}
