package binding

import (
	"fyne.io/fyne/v2"
)

type List[T any] interface {
	DataList

	Append(value T) error
	Get() ([]T, error)
	GetValue(index int) (T, error)
	Prepend(value T) error
	Remove(value T) error
	Set(list []T) error
	SetValue(index int, value T) error
}

type ExternalList[T any] interface {
	List[T]

	Reload() error
}

func NewList[T any](comparator func(T, T) bool) List[T] { _ = "STUB: not implemented"; return nil }

func BindList[T any](v *[]T, comparator func(T, T) bool) ExternalList[T] {
	_ = "STUB: not implemented"
	return nil
}

type DataList interface {
	DataItem
	GetItem(index int) (DataItem, error)
	Length() int
}

type BoolList = List[bool]

type ExternalBoolList = ExternalList[bool]

func NewBoolList() List[bool] { _ = "STUB: not implemented"; return nil }

func BindBoolList(v *[]bool) ExternalList[bool] { _ = "STUB: not implemented"; return nil }

type BytesList = List[[]byte]

type ExternalBytesList = ExternalList[[]byte]

func NewBytesList() List[[]byte] { _ = "STUB: not implemented"; return nil }

func BindBytesList(v *[][]byte) ExternalList[[]byte] { _ = "STUB: not implemented"; return nil }

type FloatList = List[float64]

type ExternalFloatList = ExternalList[float64]

func NewFloatList() List[float64] { _ = "STUB: not implemented"; return nil }

func BindFloatList(v *[]float64) ExternalList[float64] { _ = "STUB: not implemented"; return nil }

type IntList = List[int]

type ExternalIntList = ExternalList[int]

func NewIntList() List[int] { _ = "STUB: not implemented"; return nil }

func BindIntList(v *[]int) ExternalList[int] { _ = "STUB: not implemented"; return nil }

type RuneList = List[rune]

type ExternalRuneList = ExternalList[rune]

func NewRuneList() List[rune] { _ = "STUB: not implemented"; return nil }

func BindRuneList(v *[]rune) ExternalList[rune] { _ = "STUB: not implemented"; return nil }

type StringList = List[string]

type ExternalStringList = ExternalList[string]

func NewStringList() List[string] { _ = "STUB: not implemented"; return nil }

func BindStringList(v *[]string) ExternalList[string] { _ = "STUB: not implemented"; return nil }

type UntypedList = List[any]

type ExternalUntypedList = ExternalList[any]

func NewUntypedList() List[any] { _ = "STUB: not implemented"; return nil }

func BindUntypedList(v *[]any) ExternalList[any] { _ = "STUB: not implemented"; return nil }

type URIList = List[fyne.URI]

type ExternalURIList = ExternalList[fyne.URI]

func NewURIList() List[fyne.URI] { _ = "STUB: not implemented"; return nil }

func BindURIList(v *[]fyne.URI) ExternalList[fyne.URI] { _ = "STUB: not implemented"; return nil }

type listBase struct {
	base
	items []DataItem
}

func (b *listBase) GetItem(i int) (DataItem, error) {
	_ = "STUB: not implemented"
	return *new(DataItem), nil
}

func (b *listBase) Length() int { _ = "STUB: not implemented"; return 0 }

func (b *listBase) appendItem(i DataItem) { _ = "STUB: not implemented"; return }

func (b *listBase) deleteItem(i int) { _ = "STUB: not implemented"; return }

func newList[T any](comparator func(T, T) bool) *boundList[T] {
	_ = "STUB: not implemented"
	return nil
}

func newListComparable[T comparable]() *boundList[T] { _ = "STUB: not implemented"; return nil }

func newExternalList[T any](v *[]T, comparator func(T, T) bool) *boundList[T] {
	_ = "STUB: not implemented"
	return nil
}

func bindList[T any](v *[]T, comparator func(T, T) bool) *boundList[T] {
	_ = "STUB: not implemented"
	return nil
}

func bindListComparable[T comparable](v *[]T) *boundList[T] { _ = "STUB: not implemented"; return nil }

type boundList[T any] struct {
	listBase

	comparator     func(T, T) bool
	updateExternal bool
	val            *[]T

	parentListener func(int)
}

func (l *boundList[T]) Append(val T) error { _ = "STUB: not implemented"; return nil }

func (l *boundList[T]) Get() ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *boundList[T]) GetValue(i int) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (l *boundList[T]) Prepend(val T) error { _ = "STUB: not implemented"; return nil }

func (l *boundList[T]) Reload() error { _ = "STUB: not implemented"; return nil }

func (l *boundList[T]) Remove(val T) error { _ = "STUB: not implemented"; return nil }

func (l *boundList[T]) Set(v []T) error { _ = "STUB: not implemented"; return nil }

func (l *boundList[T]) doReload() (trigger bool, retErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *boundList[T]) SetValue(i int, v T) error { _ = "STUB: not implemented"; return nil }

func bindListItem[T any](v *[]T, i int, external bool, comparator func(T, T) bool) Item[T] {
	_ = "STUB: not implemented"
	return nil
}

type boundListItem[T any] struct {
	base

	comparator func(T, T) bool
	val        *[]T
	index      int
}

func (b *boundListItem[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (b *boundListItem[T]) Set(val T) error { _ = "STUB: not implemented"; return nil }

func (b *boundListItem[T]) doSet(val T) error { _ = "STUB: not implemented"; return nil }

type boundExternalListItem[T any] struct {
	boundListItem[T]

	old T
}

func (b *boundExternalListItem[T]) setIfChanged(val T) error { _ = "STUB: not implemented"; return nil }
