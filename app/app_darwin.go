//go:build !ci && !wasm && !test_web_driver && !mobile && !tinygo

package app

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#include <stdbool.h>
#include <stdlib.h>

bool isBundled();
void sendNotification(char *title, char *content);
bool scheduleNotification(char *id, char *title, char *content, double seconds);
void cancelScheduledNotification(char *id);
*/
import "C"

import (
	"time"

	"fyne.io/fyne/v2"
)

func (*fyneApp) SendNotification(n *fyne.Notification) { _ = "STUB: not implemented"; return }

func (a *fyneApp) ScheduleNotification(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *fyneApp) CancelScheduledNotification(id string) error {
	_ = "STUB: not implemented"
	return nil
}

func newDarwinNotificationID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func escapeNotificationString(in string) string { _ = "STUB: not implemented"; return "" }

func fallbackSend(cTitle, cContent *C.char) { _ = "STUB: not implemented"; return }

func fallbackNotification(title, content string) { _ = "STUB: not implemented"; return }
