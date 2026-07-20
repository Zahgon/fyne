package desktop

import "fyne.io/fyne/v2"

type App interface {
	SetSystemTrayMenu(menu *fyne.Menu)

	SetSystemTrayIcon(icon fyne.Resource)

	SetSystemTrayWindow(fyne.Window)
}
