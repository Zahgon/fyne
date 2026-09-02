package driver

import (
	"image"

	"fyne.io/fyne/v2"
)

type Painter interface {
	Paint(fyne.Canvas) image.Image
}
