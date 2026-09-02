//go:build !wasm && !test_web_driver && !tamago && !noos && !tinygo

package app

func (s *settings) load() { _ = "STUB: not implemented"; return }

func (s *settings) loadFromFile(path string) error { _ = "STUB: not implemented"; return nil }
