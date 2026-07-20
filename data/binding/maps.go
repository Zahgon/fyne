package binding

import (
	"reflect"
)

type DataMap interface {
	DataItem
	GetItem(string) (DataItem, error)
	Keys() []string
}

type ExternalUntypedMap interface {
	UntypedMap
	Reload() error
}

type UntypedMap interface {
	DataMap
	Delete(string)
	Get() (map[string]any, error)
	GetValue(string) (any, error)
	Set(map[string]any) error
	SetValue(string, any) error
}

func NewUntypedMap() UntypedMap { _ = "STUB: not implemented"; return *new(UntypedMap) }

func BindUntypedMap(d *map[string]any) ExternalUntypedMap {
	_ = "STUB: not implemented"
	return *new(ExternalUntypedMap)
}

type Struct interface {
	DataMap
	GetValue(string) (any, error)
	SetValue(string, any) error
	Reload() error
}

func BindStruct(i any) Struct { _ = "STUB: not implemented"; return *new(Struct) }

type reflectUntyped interface {
	DataItem
	get() (any, error)
	set(any) error
}

type mapBase struct {
	base

	updateExternal bool
	items          map[string]reflectUntyped
	val            *map[string]any
}

func (b *mapBase) GetItem(key string) (DataItem, error) {
	_ = "STUB: not implemented"
	return *new(DataItem), nil
}

func (b *mapBase) Keys() []string { _ = "STUB: not implemented"; return nil }

func (b *mapBase) Delete(key string) { _ = "STUB: not implemented"; return }

func (b *mapBase) Get() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *mapBase) GetValue(key string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (b *mapBase) Reload() error { _ = "STUB: not implemented"; return nil }

func (b *mapBase) Set(v map[string]any) error { _ = "STUB: not implemented"; return nil }

func (b *mapBase) SetValue(key string, d any) error { _ = "STUB: not implemented"; return nil }

func (b *mapBase) doReload() (retErr error) { _ = "STUB: not implemented"; return nil }

func (b *mapBase) setItem(key string, d reflectUntyped) { _ = "STUB: not implemented"; return }

type boundStruct struct {
	mapBase

	orig any
}

func (b *boundStruct) Reload() (retErr error) { _ = "STUB: not implemented"; return nil }

func bindUntypedMapValue(m *map[string]any, k string, external bool) reflectUntyped {
	_ = "STUB: not implemented"
	return *new(reflectUntyped)
}

type boundMapValue struct {
	base

	val *map[string]any
	key string
}

func (b *boundMapValue) get() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (b *boundMapValue) set(val any) error { _ = "STUB: not implemented"; return nil }

type boundExternalMapValue struct {
	boundMapValue

	old any
}

func (b *boundExternalMapValue) setIfChanged(val any) error { _ = "STUB: not implemented"; return nil }

type boundReflect[T any] struct {
	base

	val reflect.Value
}

func (b *boundReflect[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (b *boundReflect[T]) Set(val T) error { _ = "STUB: not implemented"; return nil }

func (b *boundReflect[T]) get() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (b *boundReflect[T]) set(val any) error { _ = "STUB: not implemented"; return nil }

func bindReflect(field reflect.Value) reflectUntyped {
	_ = "STUB: not implemented"
	return *new(reflectUntyped)
}
