package fyne

type Widget interface {
	CanvasObject

	CreateRenderer() WidgetRenderer
}

type WidgetRenderer interface {
	Destroy()

	Layout(Size)

	MinSize() Size

	Objects() []CanvasObject

	Refresh()
}
