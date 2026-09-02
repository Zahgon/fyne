package mobile

type Device interface {
	ShowVirtualKeyboard()

	ShowVirtualKeyboardType(KeyboardType)

	HideVirtualKeyboard()
}
