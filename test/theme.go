package test

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var defaultTheme fyne.Theme

var knownColorNames = [...]fyne.ThemeColorName{
	theme.ColorNameBackground,
	theme.ColorNameButton,
	theme.ColorNameDisabled,
	theme.ColorNameDisabledButton,
	theme.ColorNameError,
	theme.ColorNameFocus,
	theme.ColorNameForeground,
	theme.ColorNameForegroundOnError,
	theme.ColorNameForegroundOnPrimary,
	theme.ColorNameForegroundOnSuccess,
	theme.ColorNameForegroundOnWarning,
	theme.ColorNameHeaderBackground,
	theme.ColorNameHover,
	theme.ColorNameHyperlink,
	theme.ColorNameInputBackground,
	theme.ColorNameInputBorder,
	theme.ColorNameMenuBackground,
	theme.ColorNameOverlayBackground,
	theme.ColorNamePlaceHolder,
	theme.ColorNamePressed,
	theme.ColorNamePrimary,
	theme.ColorNameScrollBar,
	theme.ColorNameScrollBarBackground,
	theme.ColorNameSelection,
	theme.ColorNameSeparator,
	theme.ColorNameShadow,
	theme.ColorNameSuccess,
	theme.ColorNameWarning,
}

func KnownThemeVariants() map[string]fyne.ThemeVariant { _ = "STUB: not implemented"; return nil }

func NewTheme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func Theme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

type configurableTheme struct {
	colors map[fyne.ThemeColorName]color.Color
	fonts  map[fyne.TextStyle]fyne.Resource
	name   string
	sizes  map[fyne.ThemeSizeName]float32
}

var _ fyne.Theme = (*configurableTheme)(nil)

func (t *configurableTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (t *configurableTheme) Font(style fyne.TextStyle) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (*configurableTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (t *configurableTheme) Size(s fyne.ThemeSizeName) float32 { _ = "STUB: not implemented"; return 0 }
