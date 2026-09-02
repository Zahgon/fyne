package binding

import (
	"sync/atomic"

	"fyne.io/fyne/v2"
)

type preferenceLookupSetter[T any] func(fyne.Preferences) (func(string) T, func(string, T))

const keyTypeMismatchError = "A previous preference binding exists with different type for key: "

func BindPreferenceBool(key string, p fyne.Preferences) Bool {
	_ = "STUB: not implemented"
	return *new(Bool)
}

func BindPreferenceBoolList(key string, p fyne.Preferences) BoolList {
	_ = "STUB: not implemented"
	return *new(BoolList)
}

func BindPreferenceFloat(key string, p fyne.Preferences) Float {
	_ = "STUB: not implemented"
	return *new(Float)
}

func BindPreferenceFloatList(key string, p fyne.Preferences) FloatList {
	_ = "STUB: not implemented"
	return *new(FloatList)
}

func BindPreferenceInt(key string, p fyne.Preferences) Int {
	_ = "STUB: not implemented"
	return *new(Int)
}

func BindPreferenceIntList(key string, p fyne.Preferences) IntList {
	_ = "STUB: not implemented"
	return *new(IntList)
}

func BindPreferenceString(key string, p fyne.Preferences) String {
	_ = "STUB: not implemented"
	return *new(String)
}

func BindPreferenceStringList(key string, p fyne.Preferences) StringList {
	_ = "STUB: not implemented"
	return *new(StringList)
}

func bindPreferenceItem[T bool | float64 | int | string](key string, p fyne.Preferences, setLookup preferenceLookupSetter[T]) Item[T] {
	_ = "STUB: not implemented"
	return nil
}

func lookupExistingBinding[T any](key string, p fyne.Preferences) (Item[T], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func lookupExistingListBinding[T bool | float64 | int | string](key string, p fyne.Preferences) (*prefBoundList[T], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type prefBoundBase[T bool | float64 | int | string] struct {
	base
	key string

	get       func(string) T
	set       func(string, T)
	setLookup preferenceLookupSetter[T]
	cache     atomic.Pointer[T]
}

func (b *prefBoundBase[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (b *prefBoundBase[T]) Set(v T) error { _ = "STUB: not implemented"; return nil }

func (b *prefBoundBase[T]) checkForChange() { _ = "STUB: not implemented"; return }

func (b *prefBoundBase[T]) replaceProvider(p fyne.Preferences) { _ = "STUB: not implemented"; return }

type prefBoundList[T bool | float64 | int | string] struct {
	boundList[T]
	key string

	get       func(string) []T
	set       func(string, []T)
	setLookup preferenceLookupSetter[[]T]
}

func (b *prefBoundList[T]) checkForChange() { _ = "STUB: not implemented"; return }

func (b *prefBoundList[T]) replaceProvider(p fyne.Preferences) { _ = "STUB: not implemented"; return }

type internalPrefs = interface{ WriteValues(func(map[string]any)) }

func bindPreferenceListComparable[T bool | float64 | int | string](key string, p fyne.Preferences,
	setLookup preferenceLookupSetter[[]T],
) *prefBoundList[T] {
	_ = "STUB: not implemented"
	return nil
}
