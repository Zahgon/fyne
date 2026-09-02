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

func (*ShortcutPaste) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (*ShortcutPaste) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (*ShortcutPaste) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutCopy)(nil)

type ShortcutCopy struct {
	Clipboard Clipboard

	Secondary bool
}

func (*ShortcutCopy) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (*ShortcutCopy) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (*ShortcutCopy) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutCut)(nil)

type ShortcutCut struct {
	Clipboard Clipboard

	Secondary bool
}

func (*ShortcutCut) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (*ShortcutCut) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (*ShortcutCut) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutSelectAll)(nil)

type ShortcutSelectAll struct{}

func (*ShortcutSelectAll) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (*ShortcutSelectAll) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (*ShortcutSelectAll) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutUndo)(nil)

type ShortcutUndo struct{}

func (*ShortcutUndo) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (*ShortcutUndo) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (*ShortcutUndo) ShortcutName() string { _ = "STUB: not implemented"; return "" }

var _ KeyboardShortcut = (*ShortcutRedo)(nil)

type ShortcutRedo struct{}

func (*ShortcutRedo) Key() KeyName { _ = "STUB: not implemented"; return *new(KeyName) }

func (*ShortcutRedo) Mod() KeyModifier { _ = "STUB: not implemented"; return *new(KeyModifier) }

func (*ShortcutRedo) ShortcutName() string { _ = "STUB: not implemented"; return "" }
