package widget

import (
	"time"

	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*calendarLayout)(nil)

const (
	daysPerWeek      = 7
	maxWeeksPerMonth = 6
)

var minCellContent = NewLabel("22")

type Calendar struct {
	BaseWidget
	currentTime time.Time

	monthPrevious *Button
	monthNext     *Button
	monthLabel    *Label

	dates *fyne.Container

	OnChanged func(time.Time) `json:"-"`
}

func NewCalendar(cT time.Time, changed func(time.Time)) *Calendar {
	_ = "STUB: not implemented"
	return nil
}

func (c *Calendar) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *Calendar) calendarObjects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (c *Calendar) dateForButton(dayNum int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *Calendar) daysOfMonth() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (c *Calendar) monthYear() string { _ = "STUB: not implemented"; return "" }

type calendarLayout struct {
	cellSize fyne.Size
}

func newCalendarLayout() fyne.Layout { _ = "STUB: not implemented"; return *new(fyne.Layout) }

func (g *calendarLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (g *calendarLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (g *calendarLayout) getLeading(row, col int) fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func (g *calendarLayout) getTrailing(row, col int) fyne.Position {
	_ = "STUB: not implemented"
	return *new(fyne.Position)
}

func shortDayName(in string) string { _ = "STUB: not implemented"; return "" }
