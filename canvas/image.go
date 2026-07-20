package canvas

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/svg"
)

type ImageFill int

const (
	ImageFillStretch ImageFill = iota

	ImageFillContain

	ImageFillOriginal

	ImageFillCover
)

type ImageScale int32

const (
	ImageScaleSmooth ImageScale = iota

	ImageScalePixels

	ImageScaleFastest
)

var _ fyne.CanvasObject = (*Image)(nil)

type Image struct {
	baseObject

	aspect float32
	icon   *svg.Decoder
	isSVG  bool

	File     string
	Resource fyne.Resource
	Image    image.Image

	Translucency float64
	FillMode     ImageFill
	ScaleMode    ImageScale

	CornerRadius float32

	previousRender bool
}

func (i *Image) Alpha() float64 { _ = "STUB: not implemented"; return 0 }

func (i *Image) Aspect() float32 { _ = "STUB: not implemented"; return 0 }

func (i *Image) Hide() { _ = "STUB: not implemented"; return }

func (i *Image) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (i *Image) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (i *Image) Refresh() { _ = "STUB: not implemented"; return }

func (i *Image) Resize(s fyne.Size) { _ = "STUB: not implemented"; return }

func NewImageFromFile(file string) *Image { _ = "STUB: not implemented"; return nil }

func NewImageFromURI(uri fyne.URI) *Image { _ = "STUB: not implemented"; return nil }

func NewImageFromReader(read io.Reader, name string) *Image { _ = "STUB: not implemented"; return nil }

func NewImageFromResource(res fyne.Resource) *Image { _ = "STUB: not implemented"; return nil }

func NewImageFromImage(img image.Image) *Image { _ = "STUB: not implemented"; return nil }

func (i *Image) name() string { _ = "STUB: not implemented"; return "" }

func (i *Image) updateReader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (i *Image) updateAspectAndMinSize(reader io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (i *Image) imageDetailsFromReader(source io.Reader) (reader io.Reader, width, height int, aspect float32, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), 0, 0, 0, nil
}

func (i *Image) renderSVG(width, height float32) (image.Image, error) {
	_ = "STUB: not implemented"
	return *new(image.Image), nil
}
