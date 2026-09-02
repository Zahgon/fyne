//go:build !ci && !wasm && test_web_driver

package app

import (
	"net/url"
)

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }
