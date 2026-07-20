//go:build !ci && ios && !mobile

package app

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework UIKit -framework UserNotifications

#include <stdlib.h>

void openURL(char *urlStr);
void sendNotification(char *title, char *content);
*/
import "C"

import (
	"net/url"
)

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }
