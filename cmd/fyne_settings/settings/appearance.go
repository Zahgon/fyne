package settings

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	themeNameDark        = "dark"
	themeNameLight       = "light"
	themeNameSystem      = ""
	themeNameSystemLabel = "system default"
)

type Settings struct {
	fyneSettings app.SettingsSchema

	preview *fyne.Container
	colors  []fyne.CanvasObject

	userTheme fyne.Theme
}

func NewSettings() *Settings { _ = "STUB: not implemented"; return nil }

func (*Settings) AppearanceIcon() fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (s *Settings) LoadAppearanceScreen(w fyne.Window) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (s *Settings) chooseTheme(name string) { _ = "STUB: not implemented"; return }

func (s *Settings) createPreview() *fyne.Container { _ = "STUB: not implemented"; return nil }

func (s *Settings) load() { _ = "STUB: not implemented"; return }

func (s *Settings) loadFromFile(path string) error { _ = "STUB: not implemented"; return nil }

func (s *Settings) refreshPreview() { _ = "STUB: not implemented"; return }

func (s *Settings) save() error { _ = "STUB: not implemented"; return nil }

func (s *Settings) saveToFile(path string) error { _ = "STUB: not implemented"; return nil }

type primaryColorButton struct {
	widget.BaseWidget
	name string

	s *Settings
}

func newPrimaryColorButton(name string, s *Settings) *primaryColorButton {
	_ = "STUB: not implemented"
	return nil
}

func (c *primaryColorButton) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *primaryColorButton) Tapped(_ *fyne.PointEvent) { _ = "STUB: not implemented"; return }

type primaryColorButtonRenderer struct {
	c    *primaryColorButton
	rect *canvas.Rectangle
	objs []fyne.CanvasObject
}

func (c *primaryColorButtonRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (*primaryColorButtonRenderer) MinSize() fyne.Size {
	_ = "STUB: not implemented"
	return *
	//revive:disable-line:add-constant
	new(fyne.Size)
}

func (c *primaryColorButtonRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (c *primaryColorButtonRenderer) Objects() []fyne.CanvasObject {
	_ = "STUB: not implemented"
	return nil
}

func (*primaryColorButtonRenderer) Destroy() { _ = "STUB: not implemented"; return }

type previewTheme struct {
	s *Settings
	t fyne.Theme
}

func (p *previewTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (p *previewTheme) Font(s fyne.TextStyle) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (p *previewTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (p *previewTheme) Size(n fyne.ThemeSizeName) float32 { _ = "STUB: not implemented"; return 0 }

func createPreviewWidget() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}
