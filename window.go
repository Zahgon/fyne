package fyne

type Window interface {
	Title() string

	SetTitle(string)

	FullScreen() bool

	SetFullScreen(bool)

	Resize(Size)

	RequestFocus()

	FixedSize() bool

	SetFixedSize(bool)

	CenterOnScreen()

	Padded() bool

	SetPadded(bool)

	Icon() Resource

	SetIcon(Resource)

	SetMaster()

	MainMenu() *MainMenu

	SetMainMenu(*MainMenu)

	SetOnClosed(func())

	SetCloseIntercept(func())

	SetOnDropped(func(Position, []URI))

	Show()

	Hide()

	Close()

	ShowAndRun()

	Content() CanvasObject

	SetContent(CanvasObject)

	Canvas() Canvas

	Clipboard() Clipboard
}
