//go:build ignore

package main

import (
	"os"

	"fyne.io/fyne/v2"
)

func main() {
	symbolFont := "InterSymbols-Regular.ttf"
	err := createFontByStripping(symbolFont, "Inter-Regular.ttf", []rune{
		'←',
		'↑',
		'→',
		'↓',
		'↖',
		'↘',
		'↩',
		'↪',
		'↳',
		'↵',
		'⇞',
		'⇟',
		'⇤',
		'⇥',
		'⇧',
		'⌃',
		'⌘',
		'⌥',
		'⌦',
		'⌫',
		'⎋',
		'␣',
		'❖',
	})
	if err != nil {
		fyne.LogError("symbol font creation failed", err)
		os.Exit(1)
	}
}

func createFontByStripping(newFontFile, fontFile string, runes []rune) error {
	_ = "STUB: not implemented"
	return nil
}

func fontPath(filename string) string { _ = "STUB: not implemented"; return "" }
