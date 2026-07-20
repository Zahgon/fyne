package binding

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/async"
)

type preferenceItem interface {
	checkForChange()
}

type preferenceBindings struct {
	async.Map[string, preferenceItem]
}

func (b *preferenceBindings) list() []preferenceItem { _ = "STUB: not implemented"; return nil }

type preferencesMap struct {
	prefs async.Map[fyne.Preferences, *preferenceBindings]

	appPrefs fyne.Preferences
	appLock  sync.Mutex
}

func newPreferencesMap() *preferencesMap { _ = "STUB: not implemented"; return nil }

func (m *preferencesMap) ensurePreferencesAttached(p fyne.Preferences) *preferenceBindings {
	_ = "STUB: not implemented"
	return nil
}

func (m *preferencesMap) getBindings(p fyne.Preferences) *preferenceBindings {
	_ = "STUB: not implemented"
	return nil
}

func (m *preferencesMap) preferencesChanged(p fyne.Preferences) { _ = "STUB: not implemented"; return }

func (m *preferencesMap) migratePreferences(src, dst fyne.Preferences) {
	_ = "STUB: not implemented"
	return
}
