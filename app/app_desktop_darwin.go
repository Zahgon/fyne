//go:build !ci && !ios && !wasm && !test_web_driver && !mobile

package app

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#include <AppKit/AppKit.h>

bool isBundled();
void watchTheme();
*/
import "C"

import (
	"net/url"

	"fyne.io/fyne/v2"
)

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (a *fyneApp) SetSystemTrayIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func (a *fyneApp) SetSystemTrayMenu(menu *fyne.Menu) { _ = "STUB: not implemented"; return }

func (a *fyneApp) SetSystemTrayWindow(w fyne.Window) { _ = "STUB: not implemented"; return }

func themeChanged() { _ = "STUB: not implemented"; return }

func watchTheme(_ *settings) { _ = "STUB: not implemented"; return }

func (a *fyneApp) registerRepositories() { _ = "STUB: not implemented"; return }
