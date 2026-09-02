package desktop

import "image"

type Cursor interface {
	Image() (img image.Image, hotspotX, hotspotY int)
}

type StandardCursor int

func (StandardCursor) Image() (img image.Image, hotspotX, hotspotY int) {
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
