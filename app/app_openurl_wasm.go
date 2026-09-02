//go:build !ci && wasm

package app

import (
	"net/url"
)

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }
