package canvas

import (
	"image"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*Shader)(nil)

type Shader struct {
	baseObject

	Name string

	Source []byte

	SourceES []byte

	Textures map[string]image.Image

	Uniforms map[string]float32
}

func NewShader(name string, source, sourceES []byte) *Shader { _ = "STUB: not implemented"; return nil }

func (s *Shader) Hide() { _ = "STUB: not implemented"; return }

func (s *Shader) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (s *Shader) Refresh() { _ = "STUB: not implemented"; return }

func (s *Shader) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }
