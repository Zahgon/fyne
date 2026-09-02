package binding

import (
	"fyne.io/fyne/v2"
)

type Item[T any] interface {
	DataItem
	Get() (T, error)
	Set(T) error
}

type ExternalItem[T any] interface {
	Item[T]
	Reload() error
}

func NewItem[T any](comparator func(T, T) bool) Item[T] { _ = "STUB: not implemented"; return nil }

func BindItem[T any](val *T, comparator func(T, T) bool) ExternalItem[T] {
	_ = "STUB: not implemented"
	return nil
}

type Bool = Item[bool]

type ExternalBool = ExternalItem[bool]

func NewBool() Bool { _ = "STUB: not implemented"; return *new(Bool) }

func BindBool(v *bool) ExternalBool { _ = "STUB: not implemented"; return *new(ExternalBool) }

type Bytes = Item[[]byte]

type ExternalBytes = ExternalItem[[]byte]

func NewBytes() Bytes { _ = "STUB: not implemented"; return *new(Bytes) }

func BindBytes(v *[]byte) ExternalBytes { _ = "STUB: not implemented"; return *new(ExternalBytes) }

type Float = Item[float64]

type ExternalFloat = ExternalItem[float64]

func NewFloat() Float { _ = "STUB: not implemented"; return *new(Float) }

func BindFloat(v *float64) ExternalFloat { _ = "STUB: not implemented"; return *new(ExternalFloat) }

type Int = Item[int]

type ExternalInt = ExternalItem[int]

func NewInt() Int { _ = "STUB: not implemented"; return *new(Int) }

func BindInt(v *int) ExternalInt { _ = "STUB: not implemented"; return *new(ExternalInt) }

type Rune = Item[rune]

type ExternalRune = ExternalItem[rune]

func NewRune() Rune { _ = "STUB: not implemented"; return *new(Rune) }

func BindRune(v *rune) ExternalRune { _ = "STUB: not implemented"; return *new(ExternalRune) }

type String = Item[string]

type ExternalString = ExternalItem[string]

func NewString() String { _ = "STUB: not implemented"; return *new(String) }

func BindString(v *string) ExternalString { _ = "STUB: not implemented"; return *new(ExternalString) }

type URI = Item[fyne.URI]

type ExternalURI = ExternalItem[fyne.URI]

func NewURI() URI { _ = "STUB: not implemented"; return *new(URI) }

func BindURI(v *fyne.URI) ExternalURI { _ = "STUB: not implemented"; return *new(ExternalURI) }

func newItemComparable[T bool | float64 | int | rune | string]() Item[T] {
	_ = "STUB: not implemented"
	return nil
}

type item[T any] struct {
	base

	comparator func(T, T) bool
	val        *T
}

func (b *item[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (b *item[T]) Set(val T) error { _ = "STUB: not implemented"; return nil }

func bindExternalComparable[T bool | float64 | int | rune | string](val *T) ExternalItem[T] {
	_ = "STUB: not implemented"
	return nil
}

type externalItem[T any] struct {
	item[T]

	old T
}

func (b *externalItem[T]) Set(val T) error { _ = "STUB: not implemented"; return nil }

func (b *externalItem[T]) Reload() error { _ = "STUB: not implemented"; return nil }
