package canvas

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
)

const (
	DurationStandard = time.Millisecond * 300

	DurationShort = time.Millisecond * 150
)

const shaderMaxFrameDelta = 100 * time.Millisecond

func NewColorRGBAAnimation(start, stop color.Color, d time.Duration, fn func(color.Color)) *fyne.Animation {
	_ = "STUB: not implemented"
	return nil
}

func NewPositionAnimation(start, stop fyne.Position, d time.Duration, fn func(fyne.Position)) *fyne.Animation {
	_ = "STUB: not implemented"
	return nil
}

func NewSizeAnimation(start, stop fyne.Size, d time.Duration, fn func(fyne.Size)) *fyne.Animation {
	_ = "STUB: not implemented"
	return nil
}

func NewShaderAnimation(s *Shader) *fyne.Animation { _ = "STUB: not implemented"; return nil }

func advanceShaderTime(elapsed time.Duration, lastTick, now time.Time) (time.Duration, time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time)
}

func scaleChannel(start uint8, diff int, done float32) uint8 { _ = "STUB: not implemented"; return 0 }

func scaleVal(start float32, delta, done float32) float32 { _ = "STUB: not implemented"; return 0 }
