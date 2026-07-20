package fyne

import (
	"net/url"
	"sync/atomic"
	"time"
)

type App interface {
	NewWindow(title string) Window

	OpenURL(url *url.URL) error

	Icon() Resource

	SetIcon(Resource)

	Run()

	Quit()

	Driver() Driver

	UniqueID() string

	SendNotification(*Notification)

	ScheduleNotification(n *Notification, deliverAt time.Time) (*ScheduledNotification, error)

	CancelScheduledNotification(id string) error

	Settings() Settings

	Preferences() Preferences

	Storage() Storage

	Lifecycle() Lifecycle

	Metadata() AppMetadata

	CloudProvider() CloudProvider

	SetCloudProvider(CloudProvider)

	Clipboard() Clipboard

	Cache() Cache
}

var app atomic.Pointer[App]

func SetCurrentApp(current App) { _ = "STUB: not implemented"; return }

func CurrentApp() App { _ = "STUB: not implemented"; return *new(App) }

type AppMetadata struct {
	ID string

	Name string

	Version string

	Build int

	Icon Resource

	Release bool

	Custom map[string]string

	Migrations map[string]bool
}

type Lifecycle interface {
	SetOnEnteredForeground(func())

	SetOnExitedForeground(func())

	SetOnStarted(func())

	SetOnStopped(func())
}
