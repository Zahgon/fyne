package repository

import (
	"regexp"

	"fyne.io/fyne/v2"
)

const domainLabelPattern = "[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?"

var rxHostName = regexp.MustCompile("^" + domainLabelPattern + `(?:\.` + domainLabelPattern + ")*$")

func NewFileURI(path string) fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

func ParseURI(s string) (fyne.URI, error) { _ = "STUB: not implemented"; return *new(fyne.URI), nil }
