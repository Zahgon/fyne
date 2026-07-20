package fyne

type Clipboard interface {
	Content() string

	SetContent(content string)
}
