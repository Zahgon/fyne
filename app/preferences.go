package app

import (
	"errors"
	"io"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal"
)

type preferences struct {
	*internal.InMemoryPreferences

	prefLock            sync.RWMutex
	savedRecently       bool
	changedDuringSaving bool

	app                 *fyneApp
	needsSaveBeforeExit bool
}

var _ fyne.Preferences = (*preferences)(nil)

var errEmptyPreferencesStore = errors.New("empty preferences store")

type writeSyncCloser interface {
	io.WriteCloser
	Sync() error
}

func (p *preferences) forceImmediateSave() { _ = "STUB: not implemented"; return }

func (p *preferences) resetSavedRecently() { _ = "STUB: not implemented"; return }

func (p *preferences) save() error { _ = "STUB: not implemented"; return nil }

func (p *preferences) saveToStorage(writer writeSyncCloser) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *preferences) load() { _ = "STUB: not implemented"; return }

func (p *preferences) loadFromStorage(storage io.ReadCloser) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func newPreferences(app *fyneApp) *preferences { _ = "STUB: not implemented"; return nil }

func convertLists(values map[string]any) { _ = "STUB: not implemented"; return }
