package lang

import (
	"golang.org/x/text/language"

	"fyne.io/fyne/v2"
)

func SystemLocale() fyne.Locale { _ = "STUB: not implemented"; return *new(fyne.Locale) }

func closestSupportedLocale(locs []string) fyne.Locale {
	_ = "STUB: not implemented"
	return *new(fyne.Locale)
}

func localeFromTag(in language.Tag) fyne.Locale {
	_ = "STUB: not implemented"
	return *new(fyne.Locale)
}
