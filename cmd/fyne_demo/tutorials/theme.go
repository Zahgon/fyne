package tutorials

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var (
	purple = &color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	orange = &color.NRGBA{R: 198, G: 123, B: 0, A: 255}
	grey   = &color.Gray{Y: 123}
)

type customTheme struct{}

func (customTheme) Color(c fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (customTheme) Font(style fyne.TextStyle) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (customTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (customTheme) Size(s fyne.ThemeSizeName) float32 { _ = "STUB: not implemented"; return 0 }

func newCustomTheme() fyne.Theme { _ = "STUB: not implemented"; return *new(fyne.Theme) }
