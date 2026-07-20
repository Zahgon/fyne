//go:build !ci && !wasm && !test_web_driver && !android && !ios && !mobile && (linux || openbsd || freebsd || netbsd) && !tinygo && !noos && !tamago

package app

import (
	"net/url"
	"sync/atomic"
	"time"

	"github.com/rymdport/portal/settings/appearance"

	"fyne.io/fyne/v2"
)

const systemTheme = fyne.ThemeVariant(99)

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

//gosec:disable G204 -- It’s the callers responsibility to validate the input.

func findFreedesktopColorScheme() fyne.ThemeVariant {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeVariant)
}

func colorSchemeToThemeVariant(colorScheme appearance.ColorScheme) fyne.ThemeVariant {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeVariant)
}

func (a *fyneApp) SendNotification(n *fyne.Notification) { _ = "STUB: not implemented"; return }

func (a *fyneApp) ScheduleNotification(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *fyneApp) CancelScheduledNotification(id string) error {
	_ = "STUB: not implemented"
	return nil
}

var notificationID atomic.Uint64

func (a *fyneApp) sendNotificationThroughPortal(n *fyne.Notification) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *fyneApp) SetSystemTrayMenu(menu *fyne.Menu) { _ = "STUB: not implemented"; return }

func (a *fyneApp) SetSystemTrayIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func (a *fyneApp) SetSystemTrayWindow(w fyne.Window) { _ = "STUB: not implemented"; return }

func watchTheme(s *settings) { _ = "STUB: not implemented"; return }

//gosec:disable G115 -- Probably okay to cast uint32 to uint8 here.

func (a *fyneApp) registerRepositories() { _ = "STUB: not implemented"; return }

func (s *settings) applyVariant(variant fyne.ThemeVariant) { _ = "STUB: not implemented"; return }
