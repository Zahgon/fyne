//go:build !ci && android

package app

/*
#cgo LDFLAGS: -landroid -llog

#include <stdbool.h>
#include <stdlib.h>

void openURL(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx, char *url);
void sendNotification(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx, char *title, char *content);
bool scheduleNotification(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx,
	char *id, char *title, char *body, long long deliveryMillis);
void cancelScheduledNotification(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx, char *id);
*/
import "C"

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
)

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (a *fyneApp) SendNotification(n *fyne.Notification) { _ = "STUB: not implemented"; return }

func (a *fyneApp) ScheduleNotification(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *fyneApp) CancelScheduledNotification(id string) error {
	_ = "STUB: not implemented"
	return nil
}
