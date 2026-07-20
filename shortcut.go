package fyne

import "sync"

type ShortcutHandler struct {
	entry sync.Map
}

func (sh *ShortcutHandler) TypedShortcut(shortcut Shortcut) { _ = "STUB: not implemented"; return }

func (sh *ShortcutHandler) AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut)) {
	_ = "STUB: not implemented"
	return
}

func (sh *ShortcutHandler) RemoveShortcut(shortcut Shortcut) { _ = "STUB: not implemented"; return }

type Shortcut interface {
	ShortcutName() string
}

type KeyboardShortcut interface {
	Shortcut
	Key() KeyName
	Mod() KeyModifier
}

var _ KeyboardShortcut = (*ShortcutPaste)(nil)

type ShortcutPaste struct {
	Clipboard Clipboard

	Secondary bool
}

func (se *ShortcutPaste) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (se *ShortcutPaste) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (se *ShortcutPaste) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutCopy)(nil)

type ShortcutCopy struct {
	Clipboard Clipboard

	Secondary bool
}

func (se *ShortcutCopy) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (se *ShortcutCopy) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (se *ShortcutCopy) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutCut)(nil)

type ShortcutCut struct {
	Clipboard Clipboard

	Secondary bool
}

func (se *ShortcutCut) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (se *ShortcutCut) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (se *ShortcutCut) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutSelectAll)(nil)

type ShortcutSelectAll struct{}

func (se *ShortcutSelectAll) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (se *ShortcutSelectAll) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (se *ShortcutSelectAll) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutUndo)(nil)

type ShortcutUndo struct{}

func (se *ShortcutUndo) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (se *ShortcutUndo) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (se *ShortcutUndo) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutRedo)(nil)

type ShortcutRedo struct{}

func (se *ShortcutRedo) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (se *ShortcutRedo) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (se *ShortcutRedo) ShortcutName() string { _ = "STUB: not implemented"; return "" }
