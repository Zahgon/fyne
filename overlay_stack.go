package fyne

type OverlayStack interface {
	Add(overlay CanvasObject)

	List() []CanvasObject

	Remove(overlay CanvasObject)

	Top() CanvasObject
}
