package widget

type Importance int

const (
	MediumImportance Importance = iota

	HighImportance

	LowImportance

	DangerImportance

	WarningImportance

	SuccessImportance
)
