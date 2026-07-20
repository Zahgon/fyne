package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Circle)(nil)

type Circle struct {
	Position1 fyne.Position
	Position2 fyne.Position
	Hidden    bool

	FillColor   color.Color
	StrokeColor color.Color
	StrokeWidth float32

	Shadow Shadow
}

func NewCircle(color color.Color) *Circle { _ = "STUB: not implemented"; return nil }

func (c *Circle) Hide() { _ = "STUB: not implemented"; return }

func (c *Circle) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *Circle) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (c *Circle) Position() fyne.Position { _ = "STUB: not implemented"; return *new(fyne.Position) }

func (c *Circle) Refresh() { _ = "STUB: not implemented"; return }

func (c *Circle) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (c *Circle) Show() { _ = "STUB: not implemented"; return }

func (c *Circle) Size() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (c *Circle) Visible() bool { _ = "STUB: not implemented"; return false }
