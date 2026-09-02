package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type ShadowVariant int

const (
	DropShadow ShadowVariant = iota

	BoxShadow
)

type Shadow struct {
	Color      color.Color
	BlurRadius float32
	Spread     float32
	Offset     fyne.Position
	Variant    ShadowVariant
}
