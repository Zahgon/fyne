package fyne

type Validatable interface {
	Validate() error

	SetOnValidationChanged(func(error))
}

type Requireable interface {
	HasValue() bool

	SetOnRequiredChanged(func(bool))
}

type StringValidator func(string) error
