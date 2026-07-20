package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	internaltheme "fyne.io/fyne/v2/internal/theme"
)

const (
	ColorRed = internaltheme.ColorRed

	ColorOrange = internaltheme.ColorOrange

	ColorYellow = internaltheme.ColorYellow

	ColorGreen = internaltheme.ColorGreen

	ColorBlue = internaltheme.ColorBlue

	ColorPurple = internaltheme.ColorPurple

	ColorBrown = internaltheme.ColorBrown

	ColorGray = internaltheme.ColorGray

	ColorNameBackground fyne.ThemeColorName = "background"

	ColorNameButton fyne.ThemeColorName = "button"

	ColorNameDisabledButton fyne.ThemeColorName = "disabledButton"

	ColorNameDisabled fyne.ThemeColorName = "disabled"

	ColorNameError fyne.ThemeColorName = "error"

	ColorNameFocus fyne.ThemeColorName = "focus"

	ColorNameForeground fyne.ThemeColorName = "foreground"

	ColorNameForegroundOnError fyne.ThemeColorName = "foregroundOnError"

	ColorNameForegroundOnPrimary fyne.ThemeColorName = "foregroundOnPrimary"

	ColorNameForegroundOnSuccess fyne.ThemeColorName = "foregroundOnSuccess"

	ColorNameForegroundOnWarning fyne.ThemeColorName = "foregroundOnWarning"

	ColorNameHeaderBackground fyne.ThemeColorName = "headerBackground"

	ColorNameHover fyne.ThemeColorName = "hover"

	ColorNameHyperlink fyne.ThemeColorName = "hyperlink"

	ColorNameInnerWindowBorder fyne.ThemeColorName = "innerWindowBorder"

	ColorNameInnerWindowBorderInactive fyne.ThemeColorName = "innerWindowBorderInactive"

	ColorNameInputBackground fyne.ThemeColorName = "inputBackground"

	ColorNameInputBorder fyne.ThemeColorName = "inputBorder"

	ColorNameMenuBackground fyne.ThemeColorName = "menuBackground"

	ColorNameOverlayBackground fyne.ThemeColorName = "overlayBackground"

	ColorNamePlaceHolder fyne.ThemeColorName = "placeholder"

	ColorNamePressed fyne.ThemeColorName = "pressed"

	ColorNamePrimary fyne.ThemeColorName = "primary"

	ColorNameScrollBar fyne.ThemeColorName = "scrollBar"

	ColorNameScrollBarBackground fyne.ThemeColorName = "scrollBarBackground"

	ColorNameSelection fyne.ThemeColorName = "selection"

	ColorNameSeparator fyne.ThemeColorName = "separator"

	ColorNameShadow fyne.ThemeColorName = "shadow"

	ColorNameSuccess fyne.ThemeColorName = "success"

	ColorNameWarning fyne.ThemeColorName = "warning"
)

var (
	colorDarkBackground                = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkButton                    = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkDisabled                  = color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	colorDarkDisabledButton            = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkError                     = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	colorDarkForeground                = color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	colorDarkForegroundOnError         = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkForegroundOnSuccess       = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkForegroundOnWarning       = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkHeaderBackground          = color.NRGBA{R: 0x1b, G: 0x1b, B: 0x1b, A: 0xff}
	colorDarkHover                     = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x0f}
	colorDarkInnerWindowBorder         = color.NRGBA{R: 0x18, G: 0x1d, B: 0x25, A: 0xff}
	colorDarkInnerWindowBorderInactive = color.NRGBA{R: 0x1b, G: 0x1b, B: 0x1b, A: 0xff}
	colorDarkInputBackground           = color.NRGBA{R: 0x20, G: 0x20, B: 0x23, A: 0xff}
	colorDarkInputBorder               = color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	colorDarkMenuBackground            = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkOverlayBackground         = color.NRGBA{R: 0x18, G: 0x1d, B: 0x25, A: 0xff}
	colorDarkPlaceholder               = color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	colorDarkPressed                   = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	colorDarkScrollBar                 = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x99}
	colorDarkScrollBarBackground       = color.NRGBA{R: 0x20, G: 0x20, B: 0x23, A: 0xff}
	colorDarkSeparator                 = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	colorDarkShadow                    = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x77}
	colorDarkSuccess                   = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	colorDarkWarning                   = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}

	colorLightBackground                = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightButton                    = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightDisabled                  = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightDisabledButton            = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightError                     = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	colorLightFocusBlue                 = color.NRGBA{R: 0x00, G: 0x6c, B: 0xff, A: 0x2a}
	colorLightFocusBrown                = color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0x7f}
	colorLightFocusGray                 = color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0x7f}
	colorLightFocusGreen                = color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x7f}
	colorLightFocusOrange               = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0x7f}
	colorLightFocusPurple               = color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0x7f}
	colorLightFocusRed                  = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0x7f}
	colorLightFocusYellow               = color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0x7f}
	colorLightForeground                = color.NRGBA{R: 0x56, G: 0x56, B: 0x56, A: 0xff}
	colorLightForegroundOnError         = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightForegroundOnSuccess       = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightForegroundOnWarning       = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightHeaderBackground          = color.NRGBA{R: 0xf9, G: 0xf9, B: 0xf9, A: 0xff}
	colorLightHover                     = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x0f}
	colorLightInnerWindowBorder         = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightInnerWindowBorderInactive = color.NRGBA{R: 0xf9, G: 0xf9, B: 0xf9, A: 0xff}
	colorLightInputBackground           = color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	colorLightInputBorder               = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightMenuBackground            = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightOverlayBackground         = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightPlaceholder               = color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	colorLightPressed                   = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x19}
	colorLightScrollBar                 = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x99}
	colorLightScrollBarBackground       = color.NRGBA{R: 0xdb, G: 0xdb, B: 0xdb, A: 0xff}
	colorLightSelectionBlue             = color.NRGBA{R: 0x00, G: 0x6c, B: 0xff, A: 0x40}
	colorLightSelectionBrown            = color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0x3f}
	colorLightSelectionGray             = color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0x3f}
	colorLightSelectionGreen            = color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x3f}
	colorLightSelectionOrange           = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0x3f}
	colorLightSelectionPurple           = color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0x3f}
	colorLightSelectionRed              = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0x3f}
	colorLightSelectionYellow           = color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0x3f}
	colorLightSeparator                 = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightShadow                    = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x44}
	colorLightSuccess                   = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	colorLightWarning                   = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
)

func BackgroundColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func ButtonColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func Color(name fyne.ThemeColorName) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func ColorForWidget(name fyne.ThemeColorName, w fyne.Widget) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func DisabledButtonColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func DisabledColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func DisabledTextColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func ErrorColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func FocusColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func ForegroundColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func HeaderBackgroundColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func HoverColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func HyperlinkColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func InputBackgroundColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func InputBorderColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func MenuBackgroundColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func OverlayBackgroundColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func PlaceHolderColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func PressedColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func PrimaryColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func PrimaryColorNamed(name string) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func PrimaryColorNames() []string { _ = "STUB: not implemented"; return nil }

func ScrollBarColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func SelectionColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func SeparatorColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func ShadowColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func SuccessColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func WarningColor() color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func safeColorLookup(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}
