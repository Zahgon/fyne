//go:build wasm || test_web_driver

package app

func (s *settings) load() { _ = "STUB: not implemented"; return }

func (s *settings) loadFromFile(path string) error { _ = "STUB: not implemented"; return nil }

func watchFile(path string, callback func()) { _ = "STUB: not implemented"; return }

func (s *settings) watchSettings() { _ = "STUB: not implemented"; return }

func (s *settings) stopWatching() { _ = "STUB: not implemented"; return }
