//go:build !tamago && !noos

package test

import (
	"testing"

	"fyne.io/fyne/v2"
)

func AssertNotificationSent(t *testing.T, n *fyne.Notification, f func()) {
	_ = "STUB: not implemented"
	return
}

func AssertNotificationScheduled(t *testing.T, n *fyne.Notification, f func()) {
	_ = "STUB: not implemented"
	return
}
