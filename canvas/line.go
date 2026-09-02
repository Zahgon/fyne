package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Line)(nil)

type Line struct {
	Position1 fyne.Position
	Position2 fyne.Position
	Hidden    bool

	StrokeColor color.Color
	StrokeWidth float32
}

func (l *Line) Size() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *Line) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (l *Line) Position() fyne.Position { _ = "STUB: not implemented"; return *new(fyne.Position) }

func (l *Line) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (*Line) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (l *Line) Visible() bool { _ = "STUB: not implemented"; return false }

func (l *Line) Show() { _ = "STUB: not implemented"; return }

func (l *Line) Hide() { _ = "STUB: not implemented"; return }

func (l *Line) Refresh() { _ = "STUB: not implemented"; return }

func NewLine(c color.Color) *Line { _ = "STUB: not implemented"; return nil }
