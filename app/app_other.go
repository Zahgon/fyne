//go:build ci || (!ios && !android && !linux && !darwin && !windows && !freebsd && !openbsd && !netbsd && !wasm && !test_web_driver) || tamago || noos || tinygo

package app

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
)

func (a *fyneApp) OpenURL(_ *url.URL) error { _ = "STUB: not implemented"; return nil }

func (a *fyneApp) SendNotification(_ *fyne.Notification) { _ = "STUB: not implemented"; return }

func (a *fyneApp) ScheduleNotification(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *fyneApp) CancelScheduledNotification(id string) error {
	_ = "STUB: not implemented"
	return nil
}

func watchTheme(_ *settings) { _ = "STUB: not implemented"; return }

func (a *fyneApp) registerRepositories() { _ = "STUB: not implemented"; return }
