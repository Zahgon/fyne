package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	internaltheme "fyne.io/fyne/v2/internal/theme"
)

const (
	VariantDark = internaltheme.VariantDark

	VariantLight = internaltheme.VariantLight
)

const (
	fontVariantRegular = "Regular"
)

var defaultTheme, systemTheme fyne.Theme

func DarkTheme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func DefaultTheme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func LightTheme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

type builtinTheme struct {
	variant fyne.ThemeVariant

	regular, bold, italic, boldItalic, monospace, symbol fyne.Resource
}

func (t *builtinTheme) initFonts() { _ = "STUB: not implemented"; return }

func (t *builtinTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (t *builtinTheme) Font(style fyne.TextStyle) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func Current() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }

func CurrentForWidget(w fyne.CanvasObject) fyne.Theme {
	_ = "STUB: not implemented"
	return *new(fyne.Theme)
}

func currentVariant() fyne.ThemeVariant { _ = "STUB: not implemented"; return *new(fyne.ThemeVariant) }

func darkPaletteColorNamed(name fyne.ThemeColorName) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func focusColorNamed(name string) color.NRGBA { _ = "STUB: not implemented"; return *new(color.NRGBA) }

func lightPaletteColorNamed(name fyne.ThemeColorName) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func loadCustomFont(env, variant string, fallback fyne.Resource) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func selectionColorNamed(name string) color.NRGBA {
	_ = "STUB: not implemented"
	return *new(color.NRGBA)
}

func setupDefaultTheme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }
