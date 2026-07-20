package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/widget"
)

var (
	_ fyne.Draggable    = (*Slider)(nil)
	_ fyne.Focusable    = (*Slider)(nil)
	_ desktop.Hoverable = (*Slider)(nil)
	_ fyne.Tappable     = (*Slider)(nil)
	_ fyne.Disableable  = (*Slider)(nil)
)

type Slider struct {
	BaseWidget

	Value float64
	Min   float64
	Max   float64
	Step  float64

	Orientation Orientation
	OnChanged   func(float64) `json:"-"`

	OnChangeEnded func(float64) `json:"-"`

	binder        basicBinder
	hovered       bool
	focused       bool
	disabled      bool
	pendingChange bool
}

func NewSlider(min, max float64) *Slider { _ = "STUB: not implemented"; return nil }

func NewSliderWithData(min, max float64, data binding.Float) *Slider {
	_ = "STUB: not implemented"
	return nil
}

func (s *Slider) Bind(data binding.Float) { _ = "STUB: not implemented"; return }

func (s *Slider) DragEnd() { _ = "STUB: not implemented"; return }

func (s *Slider) Dragged(e *fyne.DragEvent) { _ = "STUB: not implemented"; return }

func (s *Slider) Tapped(e *fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (s *Slider) positionChanged(lastValue, currentValue float64) {
	_ = "STUB: not implemented"
	return
}

func (s *Slider) fireChangeEnded() { _ = "STUB: not implemented"; return }

func (s *Slider) FocusGained() { _ = "STUB: not implemented"; return }

func (s *Slider) FocusLost() { _ = "STUB: not implemented"; return }

func (s *Slider) MouseIn(_ *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (s *Slider) MouseMoved(_ *desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (s *Slider) MouseOut() { _ = "STUB: not implemented"; return }

func (s *Slider) TypedKey(key *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (s *Slider) TypedRune(_ rune) { _ = "STUB: not implemented"; return }

func (s *Slider) buttonDiameter(inlineIconSize float32) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (s *Slider) endOffset(inlineIconSize, innerPadding float32) float32 {
	_ = "STUB: not implemented"
	//revive:disable-next-line:add-constant -- TODO: clarify what this 1.5 is about
	return 0
}

func (s *Slider) getRatio(e *fyne.PointEvent) float64 { _ = "STUB: not implemented"; return 0 }

func (s *Slider) clampValueToRange() { _ = "STUB: not implemented"; return }

func (s *Slider) updateValue(ratio float64) { _ = "STUB: not implemented"; return }

func (s *Slider) SetValue(value float64) { _ = "STUB: not implemented"; return }

func (s *Slider) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *Slider) Disable() { _ = "STUB: not implemented"; return }

func (s *Slider) Enable() { _ = "STUB: not implemented"; return }

func (s *Slider) Disabled() bool { _ = "STUB: not implemented"; return false }

func (s *Slider) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (s *Slider) almostEqual(a, b float64) bool { _ = "STUB: not implemented"; return false }

func (s *Slider) updateFromData(data binding.DataItem) { _ = "STUB: not implemented"; return }

func (s *Slider) writeData(data binding.DataItem) { _ = "STUB: not implemented"; return }

func (s *Slider) Unbind() { _ = "STUB: not implemented"; return }

const minLongSide = float32(34)

type sliderRenderer struct {
	widget.BaseRenderer
	track          *canvas.Rectangle
	active         *canvas.Rectangle
	thumb          *canvas.Circle
	focusIndicator *canvas.Circle
	slider         *Slider
}

func (s *sliderRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (s *sliderRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (s *sliderRenderer) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (s *sliderRenderer) getOffset(iconInlineSize, innerPadding float32) float32 {
	_ = "STUB: not implemented"
	return 0
}
