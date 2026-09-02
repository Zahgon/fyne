package widget

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var timeNow = time.Now

const (
	cursorInterruptDuration = 300 * time.Millisecond
	cursorFadeAlpha         = uint8(0x16)
	cursorFadeRatio         = float32(0.2)

	fadeStart = 0.5 - cursorFadeRatio/2
	fadeStop  = 0.5 + cursorFadeRatio/2
)

type entryCursorAnimation struct {
	cursor            *canvas.Rectangle
	anim              *fyne.Animation
	lastInterruptTime time.Time
}

func newEntryCursorAnimation(cursor *canvas.Rectangle) *entryCursorAnimation {
	_ = "STUB: not implemented"
	return nil
}

func (a *entryCursorAnimation) createAnim() *fyne.Animation { _ = "STUB: not implemented"; return nil }

func (a *entryCursorAnimation) start() { _ = "STUB: not implemented"; return }

func (a *entryCursorAnimation) interrupt() { _ = "STUB: not implemented"; return }

func (a *entryCursorAnimation) stop() { _ = "STUB: not implemented"; return }
