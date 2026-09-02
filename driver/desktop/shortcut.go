package desktop

import (
	"strings"

	"fyne.io/fyne/v2"
)

var (
	_ fyne.Shortcut         = (*CustomShortcut)(nil)
	_ fyne.KeyboardShortcut = (*CustomShortcut)(nil)
)

type CustomShortcut struct {
	fyne.KeyName
	Modifier fyne.KeyModifier
}

func (cs *CustomShortcut) Key() fyne.KeyName { _ = "STUB: not implemented"; return *new(fyne.KeyName) }

func (cs *CustomShortcut) Mod() fyne.KeyModifier {
	_ = "STUB: not implemented"
	return *new(fyne.KeyModifier)
}

func (cs *CustomShortcut) ShortcutName() string { _ = "STUB: not implemented"; return "" }

func writeModifiers(w *strings.Builder, mods fyne.KeyModifier) { _ = "STUB: not implemented"; return }
