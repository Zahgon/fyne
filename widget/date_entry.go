package widget

import (
	"time"

	"fyne.io/fyne/v2"
)

var (
	_ fyne.Widget      = (*DateEntry)(nil)
	_ fyne.Tappable    = (*DateEntry)(nil)
	_ fyne.Disableable = (*DateEntry)(nil)
)

type DateEntry struct {
	Entry
	Date      *time.Time
	OnChanged func(*time.Time) `json:"-"`

	dropDown *Calendar
	popUp    *PopUp
}

func NewDateEntry() *DateEntry { _ = "STUB: not implemented"; return nil }

func (e *DateEntry) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (e *DateEntry) Enable() { _ = "STUB: not implemented"; return }

func (e *DateEntry) Disable() { _ = "STUB: not implemented"; return }

func (e *DateEntry) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (e *DateEntry) Move(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (e *DateEntry) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (e *DateEntry) SetDate(d *time.Time) { _ = "STUB: not implemented"; return }

func (e *DateEntry) popUpPos() fyne.Position { _ = "STUB: not implemented"; return *new(fyne.Position) }

func (e *DateEntry) setDate(d time.Time) { _ = "STUB: not implemented"; return }

func (e *DateEntry) setupDropDown() *Button { _ = "STUB: not implemented"; return nil }
