package fyne

type Preferences interface {
	Bool(key string) bool

	BoolWithFallback(key string, fallback bool) bool

	SetBool(key string, value bool)

	BoolList(key string) []bool

	BoolListWithFallback(key string, fallback []bool) []bool

	SetBoolList(key string, value []bool)

	Float(key string) float64

	FloatWithFallback(key string, fallback float64) float64

	SetFloat(key string, value float64)

	FloatList(key string) []float64

	FloatListWithFallback(key string, fallback []float64) []float64

	SetFloatList(key string, value []float64)

	Int(key string) int

	IntWithFallback(key string, fallback int) int

	SetInt(key string, value int)

	IntList(key string) []int

	IntListWithFallback(key string, fallback []int) []int

	SetIntList(key string, value []int)

	String(key string) string

	StringWithFallback(key, fallback string) string

	SetString(key string, value string)

	StringList(key string) []string

	StringListWithFallback(key string, fallback []string) []string

	SetStringList(key string, value []string)

	RemoveValue(key string)

	AddChangeListener(func())

	ChangeListeners() []func()
}
