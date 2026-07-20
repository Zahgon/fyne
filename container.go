package fyne

var (
	_ CanvasObject = (*Container)(nil)

	_ Accessible = (*Container)(nil)
)

type Container struct {
	size     Size
	position Position
	Hidden   bool

	Layout  Layout
	Objects []CanvasObject
}

func NewContainer(objects ...CanvasObject) *Container { _ = "STUB: not implemented"; return nil }

func NewContainerWithoutLayout(objects ...CanvasObject) *Container {
	_ = "STUB: not implemented"
	return nil
}

func NewContainerWithLayout(layout Layout, objects ...CanvasObject) *Container {
	_ = "STUB: not implemented"
	return nil
}

func (c *Container) AccessibilityLabel() string { _ = "STUB: not implemented"; return "" }

func (c *Container) AccessibilityRole() AccessibleRole {
	_ = "STUB: not implemented"
	return *new(AccessibleRole)
}

func (c *Container) Add(add CanvasObject) { _ = "STUB: not implemented"; return }

func (c *Container) AddObject(o CanvasObject) { _ = "STUB: not implemented"; return }

func (c *Container) Hide() { _ = "STUB: not implemented"; return }

func (c *Container) MinSize() Size { _ = "STUB: not implemented"; return *new(Size) }

func (c *Container) Move(pos Position) { _ = "STUB: not implemented"; return }

func (c *Container) Position() Position { _ = "STUB: not implemented"; return *new(Position) }

func (c *Container) Refresh() { _ = "STUB: not implemented"; return }

func (c *Container) Remove(rem CanvasObject) { _ = "STUB: not implemented"; return }

func (c *Container) RemoveAll() { _ = "STUB: not implemented"; return }

func (c *Container) Resize(size Size) { _ = "STUB: not implemented"; return }

func (c *Container) Show() { _ = "STUB: not implemented"; return }

func (c *Container) Size() Size { _ = "STUB: not implemented"; return *new(Size) }

func (c *Container) Visible() bool { _ = "STUB: not implemented"; return false }

func (c *Container) layout() { _ = "STUB: not implemented"; return }

func repaint(obj *Container) { _ = "STUB: not implemented"; return }
