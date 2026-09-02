package app

import (
	"sync"
)

var once sync.Once

func (a *fyneApp) cachedIconPath() string { _ = "STUB: not implemented"; return "" }

func (a *fyneApp) saveIconToCache(dirPath, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}
