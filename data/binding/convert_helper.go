package binding

import (
	"fyne.io/fyne/v2"
)

func stripFormatPrecision(in string) string { _ = "STUB: not implemented"; return "" }

func uriFromString(in string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}

func uriToString(in fyne.URI) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseBool(in string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func parseFloat(in string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseInt(in string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func formatBool(in bool) (string, error) { _ = "STUB: not implemented"; return "", nil }

func formatFloat(in float64) (string, error) { _ = "STUB: not implemented"; return "", nil }

func formatInt(in int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func internalFloatToInt(val float64) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func internalIntToFloat(val int) (float64, error) { _ = "STUB: not implemented"; return 0, nil }
