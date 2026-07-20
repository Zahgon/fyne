package widget

import (
	"fyne.io/fyne/v2"
)

var (
	_ fyne.Widget      = (*SelectEntry)(nil)
	_ fyne.Disableable = (*SelectEntry)(nil)
)

type SelectEntry struct {
	Entry
	dropDown *fyne.Menu
	popUp    *PopUpMenu
	options  []string
}

func NewSelectEntry(options []string) *SelectEntry { _ = "STUB: not implemented"; return nil }

func (e *SelectEntry) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (e *SelectEntry) Enable() { _ = "STUB: not implemented"; return }

func (e *SelectEntry) Disable() { _ = "STUB: not implemented"; return }

func (e *SelectEntry) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (e *SelectEntry) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (e *SelectEntry) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (e *SelectEntry) SetOptions(options []string) { _ = "STUB: not implemented"; return }

func (e *SelectEntry) popUpPos() fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func (e *SelectEntry) setupDropDown() *Button { _ = "STUB: not implemented"; return nil }
