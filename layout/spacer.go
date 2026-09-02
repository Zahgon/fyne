package layout

import "fyne.io/fyne/v2"

type SpacerObject interface {
	ExpandVertical() bool
	ExpandHorizontal() bool
}

type Spacer struct {
	FixHorizontal bool
	FixVertical   bool

	size   fyne.Size
	pos    fyne.Position
	hidden bool
}

func NewSpacer() fyne.CanvasObject { _ = "STUB: not implemented"; return *new(fyne.CanvasObject) }

func (s *Spacer) ExpandVertical() bool { _ = "STUB: not implemented"; return false }

func (s *Spacer) ExpandHorizontal() bool { _ = "STUB: not implemented"; return false }

func (s *Spacer) Size() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *Spacer) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (s *Spacer) Position() fyne.Position { _ = "STUB: not implemented"; return *new(fyne.Position) }

func (s *Spacer) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (*Spacer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *Spacer) Visible() bool { _ = "STUB: not implemented"; return false }

func (s *Spacer) Show() { _ = "STUB: not implemented"; return }

func (s *Spacer) Hide() { _ = "STUB: not implemented"; return }

func (*Spacer) Refresh() { _ = "STUB: not implemented"; return }
