package fyne

type DeviceOrientation int

const (
	OrientationVertical DeviceOrientation = iota

	OrientationVerticalUpsideDown

	OrientationHorizontalLeft

	OrientationHorizontalRight
)

func IsVertical(orient DeviceOrientation) bool { _ = "STUB: not implemented"; return false }

func IsHorizontal(orient DeviceOrientation) bool { _ = "STUB: not implemented"; return false }

type Device interface {
	Orientation() DeviceOrientation
	IsMobile() bool
	IsBrowser() bool
	HasKeyboard() bool
	SystemScaleForWindow(Window) float32

	Locale() Locale
}

func CurrentDevice() Device { _ = "STUB: not implemented"; return *new(Device) }
