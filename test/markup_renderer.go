package test

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	fynecanvas "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
)

type markupRenderer struct {
	indentation int
	w           strings.Builder
}

func snapshot(c fyne.Canvas) string { _ = "STUB: not implemented"; return "" }

func (r *markupRenderer) setAlignmentAttr(attrs map[string]*string, name string, a fyne.TextAlign) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setBoolAttr(attrs map[string]*string, name string, b bool) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setColorAttr(attrs map[string]*string, name string, c color.Color) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setColorAttrWithDefault(attrs map[string]*string, name string, c color.Color, d color.Color) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setFillModeAttr(attrs map[string]*string, name string, m fynecanvas.ImageFill) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setFloatAttr(attrs map[string]*string, name string, f float64) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setFloatAttrWithDefault(attrs map[string]*string, name string, f float64, d float64) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setFloatPosAttr(attrs map[string]*string, name string, x, y float64) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setSizeAttrWithDefault(attrs map[string]*string, name string, i float32, d float32) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setPosAttr(attrs map[string]*string, name string, pos fyne.Position) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setResourceAttr(attrs map[string]*string, name string, rsc fyne.Resource) {
	_ = "STUB: not implemented"
	return
}

//gosec:disable G103

func (r *markupRenderer) setScaleModeAttr(attrs map[string]*string, name string, m fynecanvas.ImageScale) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setSizeAttr(attrs map[string]*string, name string, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setStringAttr(attrs map[string]*string, name string, s string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeCanvas(c fyne.Canvas) { _ = "STUB: not implemented"; return }

func (r *markupRenderer) writeCanvasObject(obj fyne.CanvasObject, _, _ fyne.Position, _ fyne.Size) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *markupRenderer) writeBlur(b *fynecanvas.Blur, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeArc(a *fynecanvas.Arc, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeCircle(c *fynecanvas.Circle, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeCloseCanvasObject(o fyne.CanvasObject, _ fyne.Position, _ fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeCloseTag(name string) { _ = "STUB: not implemented"; return }

func (r *markupRenderer) writeContainer(_ *fyne.Container, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeIndent() { _ = "STUB: not implemented"; return }

func (r *markupRenderer) writeImage(i *fynecanvas.Image, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeLine(l *fynecanvas.Line, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeBezierCurve(bc *fynecanvas.BezierCurve, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeLinearGradient(g *fynecanvas.LinearGradient, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeRadialGradient(g *fynecanvas.RadialGradient, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeRaster(rst *fynecanvas.Raster, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writePolygon(rct *fynecanvas.RegularPolygon, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeArbitraryPolygon(p *fynecanvas.ArbitraryPolygon, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeRectangle(rct *fynecanvas.Rectangle, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeEllipse(e *fynecanvas.Ellipse, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) setShadowAttrs(attrs map[string]*string, s fynecanvas.Shadow) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeSpacer(_ *layout.Spacer, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeTag(name string, isEmpty bool, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeText(t *fynecanvas.Text, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func (r *markupRenderer) writeWidget(w fyne.Widget, attrs map[string]*string) {
	_ = "STUB: not implemented"
	return
}

func nrgbaColor(c color.Color) color.NRGBA { _ = "STUB: not implemented"; return *new(color.NRGBA) }

//gocyclo:ignore
func knownColor(c color.Color) string { _ = "STUB: not implemented"; return "" }

//gocyclo:ignore
func knownResource(rsc fyne.Resource) string { _ = "STUB: not implemented"; return "" }

//lint:ignore SA1019 This needs to stay until the API is removed.

func sortedKeys(m map[string]*string) []string { _ = "STUB: not implemented"; return nil }
