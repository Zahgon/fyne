package desktop

import (
	"fyne.io/fyne/v2"
)

const (
	KeyNone fyne.KeyName = ""

	KeyShiftLeft fyne.KeyName = "LeftShift"

	KeyShiftRight fyne.KeyName = "RightShift"

	KeyControlLeft fyne.KeyName = "LeftControl"

	KeyControlRight fyne.KeyName = "RightControl"

	KeyAltLeft fyne.KeyName = "LeftAlt"

	KeyAltRight fyne.KeyName = "RightAlt"

	KeySuperLeft fyne.KeyName = "LeftSuper"

	KeySuperRight fyne.KeyName = "RightSuper"

	KeyMenu fyne.KeyName = "Menu"

	KeyPrintScreen fyne.KeyName = "PrintScreen"

	KeyCapsLock fyne.KeyName = "CapsLock"
)

type Modifier = fyne.KeyModifier

const (
	ShiftModifier = fyne.KeyModifierShift

	ControlModifier = fyne.KeyModifierControl

	AltModifier = fyne.KeyModifierAlt

	SuperModifier = fyne.KeyModifierSuper
)

type Keyable interface {
	fyne.Focusable

	KeyDown(*fyne.KeyEvent)
	KeyUp(*fyne.KeyEvent)
}
