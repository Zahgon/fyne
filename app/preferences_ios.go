//go:build ios

package app

import "C"

func (p *preferences) storagePath() string { _ = "STUB: not implemented"; return "" }

func (a *fyneApp) storageRoot() string { _ = "STUB: not implemented"; return "" }

func (p *preferences) watch() { _ = "STUB: not implemented"; return }
