package theme

import (
	"image/color"
	"io"

	"fyne.io/fyne/v2"
)

func FromJSON(data string) (fyne.Theme, error) {
	_ = "STUB: not implemented"
	return *new(fyne.Theme), nil
}

func FromJSONWithFallback(data string, fallback fyne.Theme) (fyne.Theme, error) {
	_ = "STUB: not implemented"
	return *new(fyne.Theme), nil
}

func FromJSONReader(r io.Reader) (fyne.Theme, error) {
	_ = "STUB: not implemented"
	return *new(fyne.Theme), nil
}

func FromJSONReaderWithFallback(r io.Reader, fallback fyne.Theme) (fyne.Theme, error) {
	_ = "STUB: not implemented"
	return *new(fyne.Theme), nil
}

func fromJSONWithFallback(r io.Reader, fallback fyne.Theme) (fyne.Theme, error) {
	_ = "STUB: not implemented"
	return *new(fyne.Theme), nil
}

type jsonColor struct {
	color color.Color
}

func (h *jsonColor) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (h *jsonColor) parseColor(str string) error { _ = "STUB: not implemented"; return nil }

type uriString string

func (u uriString) resource() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

type schema struct {
	Colors      map[string]jsonColor `json:"Colors,omitempty"`
	DarkColors  map[string]jsonColor `json:"Colors-dark,omitempty"`
	LightColors map[string]jsonColor `json:"Colors-light,omitempty"`
	Sizes       map[string]float32   `json:"Sizes,omitempty"`

	Fonts map[string]uriString `json:"Fonts,omitempty"`
	Icons map[string]uriString `json:"Icons,omitempty"`
}

type jsonTheme struct {
	data     *schema
	fallback fyne.Theme
}

func (t *jsonTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (t *jsonTheme) Font(style fyne.TextStyle) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (t *jsonTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (t *jsonTheme) Size(name fyne.ThemeSizeName) float32 { _ = "STUB: not implemented"; return 0 }

func styleString(s fyne.TextStyle) string { _ = "STUB: not implemented"; return "" }
