package lang

import (
	"embed"
	"encoding/json"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"fyne.io/fyne/v2"

	"golang.org/x/text/language"
)

var (
	L = Localize

	N = LocalizePlural

	X = LocalizeKey

	XN = LocalizePluralKey

	bundle    *i18n.Bundle
	localizer *i18n.Localizer
	setupOnce sync.Once

	//go:embed translations
	translations embed.FS
	translated   []language.Tag
)

func Localize(in string, data ...any) string { _ = "STUB: not implemented"; return "" }

func LocalizeKey(key, fallback string, data ...any) string { _ = "STUB: not implemented"; return "" }

func LocalizePlural(in string, count int, data ...any) string { _ = "STUB: not implemented"; return "" }

func LocalizePluralKey(key, fallback string, count int, data ...any) string {
	_ = "STUB: not implemented"
	return ""
}

func AddTranslations(r fyne.Resource) error { _ = "STUB: not implemented"; return nil }

func AddTranslationsForLocale(data []byte, l fyne.Locale) error {
	_ = "STUB: not implemented"
	return nil
}

func AddTranslationsFS(fs embed.FS, dir string) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

func addLanguage(data []byte, name string) error { _ = "STUB: not implemented"; return nil }

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	translated = []language.Tag{language.Make("en")}
	err := AddTranslationsFS(translations, "translations")
	if err != nil {
		fyne.LogError("Error occurred loading built-in translations", err)
	}
}

func fallbackWithData(key, fallback string, data any) string { _ = "STUB: not implemented"; return "" }

func setupLang(lang string) { _ = "STUB: not implemented"; return }

func updateLocalizer() { _ = "STUB: not implemented"; return }
