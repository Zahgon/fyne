package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
)

func FromLegacy(t fyne.LegacyTheme) fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

var _ fyne.Theme = (*legacyWrapper)(nil)

type legacyWrapper struct {
	old fyne.LegacyTheme
}

func (l *legacyWrapper) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	//revive:disable-line:identical-switch-branches
	return *new(color.Color)
}

func (l *legacyWrapper) Font(s fyne.TextStyle) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (*legacyWrapper) Icon(n fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (l *legacyWrapper) Size(n fyne.ThemeSizeName) float32 { _ = "STUB: not implemented"; return 0 }
