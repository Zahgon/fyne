package widget

import (
	"fyne.io/fyne/v2"
)

type weekday int

const (
	monday weekday = iota
	sunday
	saturday
)

type localeSetting struct {
	dateFormat   string
	weekStartDay weekday
}

const defaultDateFormat = "02/01/2006"

var localeSettings = map[string]localeSetting{
	"": {
		dateFormat:   defaultDateFormat,
		weekStartDay: monday,
	},
	"BR": {
		weekStartDay: sunday,
	},
	"BZ": {
		weekStartDay: sunday,
	},
	"CA": {
		weekStartDay: sunday,
	},
	"CO": {
		weekStartDay: sunday,
	},
	"DE": {
		dateFormat: "02.01.2006",
	},
	"DO": {
		weekStartDay: sunday,
	},
	"GT": {
		weekStartDay: sunday,
	},
	"JP": {
		weekStartDay: sunday,
	},
	"MX": {
		weekStartDay: sunday,
	},
	"NI": {
		weekStartDay: sunday,
	},
	"PE": {
		weekStartDay: sunday,
	},
	"PA": {
		weekStartDay: sunday,
	},
	"PY": {
		weekStartDay: sunday,
	},
	"SE": {
		dateFormat: "2006-01-02",
	},
	"US": {
		dateFormat:   "01/02/2006",
		weekStartDay: sunday,
	},
	"VE": {
		weekStartDay: sunday,
	},
	"ZA": {
		weekStartDay: sunday,
	},
}

func getLocaleDateFormat() string { _ = "STUB: not implemented"; return "" }

func getLocaleWeekStart() weekday { _ = "STUB: not implemented"; return *new(weekday) }

func lookupLocaleSetting(l fyne.Locale) localeSetting {
	_ = "STUB: not implemented"
	return *new(localeSetting)
}
