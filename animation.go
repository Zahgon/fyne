package fyne

import "time"

type AnimationCurve func(float32) float32

const AnimationRepeatForever = -1

var (
	AnimationEaseInOut = animationEaseInOut

	AnimationEaseIn = animationEaseIn

	AnimationEaseOut = animationEaseOut

	AnimationLinear = animationLinear
)

type Animation struct {
	AutoReverse bool
	Curve       AnimationCurve
	Duration    time.Duration
	RepeatCount int
	Tick        func(float32)
}

func NewAnimation(d time.Duration, fn func(float32)) *Animation {
	_ = "STUB: not implemented"
	return nil
}

func (a *Animation) Start() { _ = "STUB: not implemented"; return }

func (a *Animation) Stop() { _ = "STUB: not implemented"; return }

func animationEaseIn(val float32) float32 { _ = "STUB: not implemented"; return 0 }

func animationEaseInOut(val float32) float32 { _ = "STUB: not implemented"; return 0 }

func animationEaseOut(val float32) float32 { _ = "STUB: not implemented"; return 0 }

func animationLinear(val float32) float32 { _ = "STUB: not implemented"; return 0 }
