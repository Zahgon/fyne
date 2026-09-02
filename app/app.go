package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal"
	"fyne.io/fyne/v2/internal/app"
	"fyne.io/fyne/v2/internal/scheduler"
)

var _ fyne.App = (*fyneApp)(nil)

type fyneApp struct {
	driver    fyne.Driver
	clipboard fyne.Clipboard
	icon      fyne.Resource
	uniqueID  string
	missingID bool

	cache     fyne.Cache
	cloud     fyne.CloudProvider
	lifecycle app.Lifecycle
	settings  *settings
	storage   fyne.Storage
	prefs     fyne.Preferences
	scheduler *scheduler.Scheduler
}

func (a *fyneApp) CloudProvider() fyne.CloudProvider {
	_ = "STUB: not implemented"
	return *new(fyne.CloudProvider)
}

func (a *fyneApp) Icon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func (a *fyneApp) SetIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func (a *fyneApp) UniqueID() string { _ = "STUB: not implemented"; return "" }

func (a *fyneApp) NewWindow(title string) fyne.Window {
	_ = "STUB: not implemented"
	return *new(fyne.Window)
}

func (a *fyneApp) Run() { _ = "STUB: not implemented"; return }

func (a *fyneApp) Quit() { _ = "STUB: not implemented"; return }

func (a *fyneApp) Driver() fyne.Driver { _ = "STUB: not implemented"; return *new(fyne.Driver) }

func (a *fyneApp) Settings() fyne.Settings { _ = "STUB: not implemented"; return *new(fyne.Settings) }

func (a *fyneApp) Storage() fyne.Storage { _ = "STUB: not implemented"; return *new(fyne.Storage) }

func (a *fyneApp) Preferences() fyne.Preferences {
	_ = "STUB: not implemented"
	return *new(fyne.Preferences)
}

func (a *fyneApp) Lifecycle() fyne.Lifecycle {
	_ = "STUB: not implemented"
	return *new(fyne.Lifecycle)
}

func (a *fyneApp) newDefaultPreferences() *preferences { _ = "STUB: not implemented"; return nil }

func (a *fyneApp) Cache() fyne.Cache { _ = "STUB: not implemented"; return *new(fyne.Cache) }

func (a *fyneApp) Clipboard() fyne.Clipboard {
	_ = "STUB: not implemented"
	return *new(fyne.Clipboard)
}

func New() fyne.App { _ = "STUB: not implemented"; return *new(fyne.App) }

func makeStoreDocs(id string, s *store) *internal.Docs { _ = "STUB: not implemented"; return nil }

func newAppWithDriver(d fyne.Driver, clipboard fyne.Clipboard, id string) fyne.App {
	_ = "STUB: not implemented"
	return *new(fyne.App)
}

type systrayDriver interface {
	SetSystemTrayMenu(*fyne.Menu)
	SetSystemTrayIcon(fyne.Resource)
	SetSystemTrayWindow(fyne.Window)
}
