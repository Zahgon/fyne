//go:build wasm

package app

import (
	"bytes"
	"io"
)

const preferencesLocalStorageKey = "fyne-preferences.json"

func (a *fyneApp) storageRoot() string { _ = "STUB: not implemented"; return "" }

func (p *preferences) storageReader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (p *preferences) storageWriter() (writeSyncCloser, error) {
	_ = "STUB: not implemented"
	return *new(writeSyncCloser), nil
}

func (p *preferences) watch() { _ = "STUB: not implemented"; return }

type readerNopCloser struct {
	reader io.Reader
}

func (r readerNopCloser) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r readerNopCloser) Close() error { _ = "STUB: not implemented"; return nil }

type localStorageWriter struct {
	bytes.Buffer
	key string
}

func (s *localStorageWriter) Sync() error { _ = "STUB: not implemented"; return nil }

func (s *localStorageWriter) Close() error { _ = "STUB: not implemented"; return nil }
