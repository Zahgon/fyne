package widget

import (
	"sync"
	"sync/atomic"

	"fyne.io/fyne/v2/data/binding"
)

type basicBinder struct {
	callback atomic.Pointer[func(binding.DataItem)]

	dataListenerPairLock sync.RWMutex
	dataListenerPair     annotatedListener
}

func (binder *basicBinder) Bind(data binding.DataItem) { _ = "STUB: not implemented"; return }

func (binder *basicBinder) CallWithData(f func(data binding.DataItem)) {
	_ = "STUB: not implemented"
	return
}

func (binder *basicBinder) SetCallback(f func(data binding.DataItem)) {
	_ = "STUB: not implemented"
	return
}

func (binder *basicBinder) Unbind() { _ = "STUB: not implemented"; return }

func (binder *basicBinder) unbindLocked() { _ = "STUB: not implemented"; return }

type annotatedListener struct {
	data     binding.DataItem
	listener binding.DataListener
}
