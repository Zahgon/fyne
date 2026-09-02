package container

import (
	"fyne.io/fyne/v2"
	intTheme "fyne.io/fyne/v2/internal/theme"
	"fyne.io/fyne/v2/widget"
)

type ThemeOverride struct {
	widget.BaseWidget

	Content fyne.CanvasObject
	Theme   fyne.Theme

	holder *fyne.Container

	mobile bool
}

func NewThemeOverride(obj fyne.CanvasObject, th fyne.Theme) *ThemeOverride {
	_ = "STUB: not implemented"
	return nil
}

func (t *ThemeOverride) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *ThemeOverride) Refresh() { _ = "STUB: not implemented"; return }

func (t *ThemeOverride) SetDeviceIsMobile(on bool) { _ = "STUB: not implemented"; return }

type featureTheme struct {
	fyne.Theme

	over *ThemeOverride
}

func addFeatures(th fyne.Theme, o *ThemeOverride) fyne.Theme {
	_ = "STUB: not implemented"
	return *new(fyne.Theme)
}

func (f *featureTheme) Feature(n intTheme.FeatureName) any {
	_ = "STUB: not implemented"
	return *new(any)
}

type overrideRenderer struct {
	parent *ThemeOverride

	objs []fyne.CanvasObject
}

func (*overrideRenderer) Destroy() { _ = "STUB: not implemented"; return }

func (r *overrideRenderer) Layout(s fyne.Size) { _ = "STUB: not implemented"; return }

func (r *overrideRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *overrideRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (*overrideRenderer) Refresh() { _ = "STUB: not implemented"; return }
