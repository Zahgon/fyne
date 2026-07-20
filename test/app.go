package test

import (
	"net/url"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	intapp "fyne.io/fyne/v2/internal/app"
)

func init() {
	NewApp()
}

type app struct {
	driver       *driver
	settings     *testSettings
	prefs        fyne.Preferences
	propertyLock sync.RWMutex
	storage      fyne.Storage
	lifecycle    intapp.Lifecycle
	cache        fyne.Cache
	clip         fyne.Clipboard
	cloud        fyne.CloudProvider

	appliedTheme              fyne.Theme
	lastNotification          *fyne.Notification
	scheduledNotifications    map[string]*fyne.ScheduledNotification
	lastScheduledNotification *fyne.ScheduledNotification
	lastCancelledScheduleID   string
}

func (a *app) CloudProvider() fyne.CloudProvider {
	_ = "STUB: not implemented"
	return *new(fyne.CloudProvider)
}

func (a *app) Icon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func (a *app) SetIcon(fyne.Resource) { _ = "STUB: not implemented"; return }

func (a *app) NewWindow(title string) fyne.Window {
	_ = "STUB: not implemented"
	return *new(fyne.Window)
}

func (a *app) OpenURL(_ *url.URL) error { _ = "STUB: not implemented"; return nil }

func (a *app) Run() { _ = "STUB: not implemented"; return }

func (a *app) Quit() { _ = "STUB: not implemented"; return }

func (a *app) Cache() fyne.Cache { _ = "STUB: not implemented"; return *new(fyne.Cache) }

func (a *app) Clipboard() fyne.Clipboard { _ = "STUB: not implemented"; return *new(fyne.Clipboard) }

func (a *app) UniqueID() string { _ = "STUB: not implemented"; return "" }

func (a *app) Driver() fyne.Driver { _ = "STUB: not implemented"; return *new(fyne.Driver) }

func (a *app) SendNotification(notify *fyne.Notification) { _ = "STUB: not implemented"; return }

func (a *app) ScheduleNotification(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *app) CancelScheduledNotification(id string) error { _ = "STUB: not implemented"; return nil }

func (a *app) SetCloudProvider(p fyne.CloudProvider) { _ = "STUB: not implemented"; return }

func (a *app) Settings() fyne.Settings { _ = "STUB: not implemented"; return *new(fyne.Settings) }

func (a *app) Preferences() fyne.Preferences {
	_ = "STUB: not implemented"
	return *new(fyne.Preferences)
}

func (a *app) Storage() fyne.Storage { _ = "STUB: not implemented"; return *new(fyne.Storage) }

func (a *app) Lifecycle() fyne.Lifecycle { _ = "STUB: not implemented"; return *new(fyne.Lifecycle) }

func (a *app) Metadata() fyne.AppMetadata { _ = "STUB: not implemented"; return *new(fyne.AppMetadata) }

func (a *app) lastAppliedTheme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func (a *app) transitionCloud(p fyne.CloudProvider) { _ = "STUB: not implemented"; return }

func NewApp() fyne.App { _ = "STUB: not implemented"; return *new(fyne.App) }

type testSettings struct {
	primaryColor string
	scale        float32
	theme        fyne.Theme

	listeners       []func(fyne.Settings)
	changeListeners []chan fyne.Settings
	propertyLock    sync.RWMutex
	app             *app
}

func (s *testSettings) AddChangeListener(listener chan fyne.Settings) {
	_ = "STUB: not implemented"
	return
}

func (s *testSettings) AddListener(listener func(fyne.Settings)) { _ = "STUB: not implemented"; return }

func (s *testSettings) BuildType() fyne.BuildType {
	_ = "STUB: not implemented"
	return *new(fyne.BuildType)
}

func (s *testSettings) PrimaryColor() string { _ = "STUB: not implemented"; return "" }

func (s *testSettings) SetTheme(theme fyne.Theme) { _ = "STUB: not implemented"; return }

func (s *testSettings) ShowAnimations() bool { _ = "STUB: not implemented"; return false }

func (s *testSettings) Theme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func (s *testSettings) ThemeVariant() fyne.ThemeVariant {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeVariant)
}

func (s *testSettings) Scale() float32 { _ = "STUB: not implemented"; return 0 }

func (s *testSettings) apply() { _ = "STUB: not implemented"; return }
