//go:generate go run gen.go

package binding

import (
	"errors"
	"reflect"
	"sync"
)

var (
	errKeyNotFound = errors.New("key not found")
	errOutOfBounds = errors.New("index out of bounds")
	errParseFailed = errors.New("format did not match 1 value")

	prefBinds = newPreferencesMap()
)

type DataItem interface {
	AddListener(DataListener)

	RemoveListener(DataListener)
}

type DataListener interface {
	DataChanged()
}

func NewDataListener(fn func()) DataListener { _ = "STUB: not implemented"; return *new(DataListener) }

type listener struct {
	callback func()
}

func (l *listener) DataChanged() { _ = "STUB: not implemented"; return }

type base struct {
	listeners []DataListener

	lock sync.RWMutex
}

func (b *base) AddListener(l DataListener) { _ = "STUB: not implemented"; return }

func (b *base) RemoveListener(l DataListener) { _ = "STUB: not implemented"; return }

func (b *base) trigger() { _ = "STUB: not implemented"; return }

func (b *base) triggerFromMain() { _ = "STUB: not implemented"; return }

type Untyped = Item[any]

func NewUntyped() Untyped { _ = "STUB: not implemented"; return *new(Untyped) }

type ExternalUntyped = ExternalItem[any]

func BindUntyped(v any) ExternalUntyped { _ = "STUB: not implemented"; return *new(ExternalUntyped) }

type boundExternalUntyped struct {
	base

	val reflect.Value
	old any
}

func (b *boundExternalUntyped) Get() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (b *boundExternalUntyped) Set(val any) error { _ = "STUB: not implemented"; return nil }

func (b *boundExternalUntyped) Reload() error { _ = "STUB: not implemented"; return nil }
