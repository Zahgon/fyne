package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/async"
)

type SettingsSchema struct {
	ThemeName         string  `json:"theme"`
	Scale             float32 `json:"scale"`
	PrimaryColor      string  `json:"primary_color"`
	CloudName         string  `json:"cloud_name"`
	CloudConfig       string  `json:"cloud_config"`
	DisableAnimations bool    `json:"no_animations"`
}

func (sc *SettingsSchema) StoragePath() string { _ = "STUB: not implemented"; return "" }

var _ fyne.Settings = (*settings)(nil)

const (
	themeVariantNameDark  = "dark"
	themeVariantNameLight = "light"
)

type settings struct {
	theme          fyne.Theme
	themeSpecified bool
	variant        fyne.ThemeVariant

	listeners       []func(fyne.Settings)
	changeListeners async.Map[chan fyne.Settings, bool]
	watcher         any

	schema SettingsSchema
}

func (s *settings) BuildType() fyne.BuildType {
	_ = "STUB: not implemented"
	return *new(fyne.BuildType)
}

func (s *settings) PrimaryColor() string { _ = "STUB: not implemented"; return "" }

func (s *settings) OverrideTheme(theme fyne.Theme, name string) { _ = "STUB: not implemented"; return }

func (s *settings) Theme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func (s *settings) SetTheme(theme fyne.Theme) { _ = "STUB: not implemented"; return }

func (s *settings) ShowAnimations() bool { _ = "STUB: not implemented"; return false }

func (s *settings) ThemeVariant() fyne.ThemeVariant {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeVariant)
}

func (s *settings) applyTheme(theme fyne.Theme, variant fyne.ThemeVariant) {
	_ = "STUB: not implemented"
	return
}

func (s *settings) Scale() float32 { _ = "STUB: not implemented"; return 0 }

func (s *settings) AddChangeListener(listener chan fyne.Settings) {
	_ = "STUB: not implemented"
	return
}

func (s *settings) AddListener(listener func(fyne.Settings)) { _ = "STUB: not implemented"; return }

func (s *settings) apply() { _ = "STUB: not implemented"; return }

func (s *settings) fileChanged() { _ = "STUB: not implemented"; return }

func (s *settings) setupTheme() { _ = "STUB: not implemented"; return }

func (s *settings) explicitThemeVariantName() string { _ = "STUB: not implemented"; return "" }

func loadSettings() *settings { _ = "STUB: not implemented"; return nil }
