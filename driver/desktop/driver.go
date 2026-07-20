package desktop

import "fyne.io/fyne/v2"

type Driver interface {
	CreateSplashWindow() fyne.Window

	CurrentKeyModifiers() fyne.KeyModifier

	HasSecondaryDisplay() bool
}
