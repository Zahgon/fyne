//go:build !ci && !android && !ios && !mobile && (wasm || test_web_driver)

package app

import (
	"syscall/js"
	"time"

	"fyne.io/fyne/v2"
)

func (a *fyneApp) ScheduleNotification(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *fyneApp) CancelScheduledNotification(id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *fyneApp) SendNotification(n *fyne.Notification) { _ = "STUB: not implemented"; return }

func (a *fyneApp) showNotification(data *fyne.Notification, notification *js.Value) {
	_ = "STUB: not implemented"
	return
}

var themeChanged = js.FuncOf(func(this js.Value, args []js.Value) any {
	if len(args) > 0 && args[0].Type() == js.TypeObject {
		fyne.CurrentApp().Settings().(*settings).setupTheme()
	}
	return nil
})

func watchTheme(_ *settings) { _ = "STUB: not implemented"; return }

func stopWatchingTheme() { _ = "STUB: not implemented"; return }

func (a *fyneApp) registerRepositories() { _ = "STUB: not implemented"; return }
