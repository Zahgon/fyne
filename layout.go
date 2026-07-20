package fyne

type Layout interface {
	Layout([]CanvasObject, Size)

	MinSize(objects []CanvasObject) Size
}
