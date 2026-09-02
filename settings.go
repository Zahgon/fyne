package fyne

type BuildType int

const (
	BuildStandard BuildType = iota

	BuildDebug

	BuildRelease
)

type Settings interface {
	Theme() Theme
	SetTheme(Theme)

	ThemeVariant() ThemeVariant
	Scale() float32

	PrimaryColor() string

	AddChangeListener(chan Settings)

	AddListener(func(Settings))

	BuildType() BuildType

	ShowAnimations() bool
}
