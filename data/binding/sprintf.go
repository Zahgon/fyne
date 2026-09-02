package binding

type sprintfString struct {
	String

	format string
	source []DataItem
	err    error
}

func NewSprintf(format string, b ...DataItem) String {
	_ = "STUB: not implemented"
	return *new(String)
}

func (s *sprintfString) DataChanged() { _ = "STUB: not implemented"; return }

func (s *sprintfString) Get() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *sprintfString) Set(str string) error { _ = "STUB: not implemented"; return nil }

func StringToStringWithFormat(str String, format string) String {
	_ = "STUB: not implemented"
	return *new(String)
}
