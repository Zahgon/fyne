//go:build !flatpak && !windows && !android && !ios && !wasm && !js

package dialog

func fileOpenOSOverride(_ *FileDialog) bool { _ = "STUB: not implemented"; return false }

func fileSaveOSOverride(_ *FileDialog) bool { _ = "STUB: not implemented"; return false }
