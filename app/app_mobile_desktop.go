//go:build !ci && mobile && !android && !ios && !windows

package app

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
)

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

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
