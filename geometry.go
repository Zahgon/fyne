package fyne

var (
	_ Vector2 = (*Delta)(nil)
	_ Vector2 = (*Position)(nil)
	_ Vector2 = (*Size)(nil)
)

type Vector2 interface {
	Components() (float32, float32)
	IsZero() bool
}

type Delta struct {
	DX, DY float32
}

func NewDelta(dx float32, dy float32) Delta { _ = "STUB: not implemented"; return *new(Delta) }

func (v Delta) Components() (float32, float32) { _ = "STUB: not implemented"; return 0, 0 }

func (v Delta) IsZero() bool { _ = "STUB: not implemented"; return false }

type Position struct {
	X float32
	Y float32
}

func NewPos(x float32, y float32) Position { _ = "STUB: not implemented"; return *new(Position) }

func NewSquareOffsetPos(length float32) Position { _ = "STUB: not implemented"; return *new(Position) }

func (p Position) Add(v Vector2) Position { _ = "STUB: not implemented"; return *new(Position) }

func (p Position) AddXY(x, y float32) Position { _ = "STUB: not implemented"; return *new(Position) }

func (p Position) Components() (float32, float32) { _ = "STUB: not implemented"; return 0, 0 }

func (p Position) IsZero() bool { _ = "STUB: not implemented"; return false }

func (p Position) Subtract(v Vector2) Position { _ = "STUB: not implemented"; return *new(Position) }

func (p Position) SubtractXY(x, y float32) Position {
	_ = "STUB: not implemented"
	return *new(Position)
}

type Size struct {
	Width  float32
	Height float32
}

func NewSize(w float32, h float32) Size { _ = "STUB: not implemented"; return *new(Size) }

func NewSquareSize(side float32) Size { _ = "STUB: not implemented"; return *new(Size) }

func (s Size) Add(v Vector2) Size { _ = "STUB: not implemented"; return *new(Size) }

func (s Size) AddWidthHeight(width, height float32) Size {
	_ = "STUB: not implemented"
	return *new(Size)
}

func (s Size) IsZero() bool { _ = "STUB: not implemented"; return false }

func (s Size) Max(v Vector2) Size { _ = "STUB: not implemented"; return *new(Size) }

func (s Size) Min(v Vector2) Size { _ = "STUB: not implemented"; return *new(Size) }

func (s Size) Components() (float32, float32) { _ = "STUB: not implemented"; return 0, 0 }

func (s Size) Subtract(v Vector2) Size { _ = "STUB: not implemented"; return *new(Size) }

func (s Size) SubtractWidthHeight(width, height float32) Size {
	_ = "STUB: not implemented"
	return *new(Size)
}
