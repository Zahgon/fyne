package app

import (
	"time"

	"fyne.io/fyne/v2"
)

func (a *fyneApp) scheduleViaScheduler(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *fyneApp) cancelViaScheduler(id string) error { _ = "STUB: not implemented"; return nil }
