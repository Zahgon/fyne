package binding

func BoolToString(v Bool) String { _ = "STUB: not implemented"; return *new(String) }

func BoolToStringWithFormat(v Bool, format string) String {
	_ = "STUB: not implemented"
	return *new(String)
}

func FloatToString(v Float) String { _ = "STUB: not implemented"; return *new(String) }

func FloatToStringWithFormat(v Float, format string) String {
	_ = "STUB: not implemented"
	return *new(String)
}

func IntToFloat(val Int) Float { _ = "STUB: not implemented"; return *new(Float) }

func FloatToInt(v Float) Int { _ = "STUB: not implemented"; return *new(Int) }

func IntToString(v Int) String { _ = "STUB: not implemented"; return *new(String) }

func IntToStringWithFormat(v Int, format string) String {
	_ = "STUB: not implemented"
	return *new(String)
}

func URIToString(v URI) String { _ = "STUB: not implemented"; return *new(String) }

func ItemToString[T any](v Item[T], formatter func(T) (string, error), parser func(string) (T, error), comparator func(T, T) bool) String {
	_ = "STUB: not implemented"
	return *new(String)
}

func StringToBool(str String) Bool { _ = "STUB: not implemented"; return *new(Bool) }

func StringToBoolWithFormat(str String, format string) Bool {
	_ = "STUB: not implemented"
	return *new(Bool)
}

func StringToFloat(str String) Float { _ = "STUB: not implemented"; return *new(Float) }

func StringToFloatWithFormat(str String, format string) Float {
	_ = "STUB: not implemented"
	return *new(Float)
}

func StringToInt(str String) Int { _ = "STUB: not implemented"; return *new(Int) }

func StringToIntWithFormat(str String, format string) Int {
	_ = "STUB: not implemented"
	return *new(Int)
}

func StringToURI(str String) URI { _ = "STUB: not implemented"; return *new(URI) }

func toString[T any](v Item[T], formatter func(T) (string, error), comparator func(T, T) bool, parser func(string) (T, error)) *toStringFrom[T] {
	_ = "STUB: not implemented"
	return nil
}

func toStringComparable[T bool | float64 | int](v Item[T], formatter func(T) (string, error), parser func(string) (T, error)) *toStringFrom[T] {
	_ = "STUB: not implemented"
	return nil
}

func toStringWithFormat[T any](v Item[T], format, defaultFormat string, formatter func(T) (string, error), comparator func(T, T) bool, parser func(string) (T, error)) String {
	_ = "STUB: not implemented"
	return *new(String)
}

func toStringWithFormatComparable[T bool | float64 | int](v Item[T], format, defaultFormat string, formatter func(T) (string, error), parser func(string) (T, error)) String {
	_ = "STUB: not implemented"
	return *new(String)
}

type convertBaseItem struct {
	base
}

func (s *convertBaseItem) DataChanged() { _ = "STUB: not implemented"; return }

type toStringFrom[T any] struct {
	convertBaseItem

	format string

	formatter  func(T) (string, error)
	comparator func(T, T) bool
	parser     func(string) (T, error)

	from Item[T]
}

func (s *toStringFrom[T]) Get() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *toStringFrom[T]) Set(str string) error { _ = "STUB: not implemented"; return nil }

type fromStringTo[T any] struct {
	convertBaseItem

	format    string
	formatter func(string) (T, error)
	parser    func(T) (string, error)

	from String
}

func (s *fromStringTo[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (s *fromStringTo[T]) Set(val T) error { _ = "STUB: not implemented"; return nil }

type toInt[T float64] struct {
	convertBaseItem

	formatter func(int) (T, error)
	parser    func(T) (int, error)

	from Item[T]
}

func (s *toInt[T]) Get() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *toInt[T]) Set(v int) error { _ = "STUB: not implemented"; return nil }

type fromIntTo[T float64] struct {
	convertBaseItem

	formatter func(int) (T, error)
	parser    func(T) (int, error)
	from      Item[int]
}

func (s *fromIntTo[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (s *fromIntTo[T]) Set(val T) error { _ = "STUB: not implemented"; return nil }
