//go:build tamago || noos || tinygo

package app

func (s *settings) load() { _ = "STUB: not implemented"; return }

func (s *settings) loadFromFile(_ string) error { _ = "STUB: not implemented"; return nil }

func watchFile(_ string, _ func()) { _ = "STUB: not implemented"; return }

func (s *settings) watchSettings() { _ = "STUB: not implemented"; return }

func (s *settings) stopWatching() { _ = "STUB: not implemented"; return }
