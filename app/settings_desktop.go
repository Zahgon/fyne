//go:build !android && !ios && !mobile && !wasm && !test_web_driver && !tamago && !noos && !tinygo

package app

import (
	"github.com/fsnotify/fsnotify"
)

func watchFileAddTarget(watcher *fsnotify.Watcher, path string) { _ = "STUB: not implemented"; return }

func ensureDirExists(dir string) { _ = "STUB: not implemented"; return }

func watchFile(path string, callback func()) *fsnotify.Watcher {
	_ = "STUB: not implemented"
	return nil
}

func (s *settings) watchSettings() { _ = "STUB: not implemented"; return }

func (s *settings) stopWatching() { _ = "STUB: not implemented"; return }
