package repository

import (
	"net/url"

	"fyne.io/fyne/v2"
)

func EqualURI(t1, t2 fyne.URI) bool { _ = "STUB: not implemented"; return false }

var _ fyne.URI = &uri{}

type uri struct {
	url.URL
}

func (u *uri) Extension() string { _ = "STUB: not implemented"; return "" }

func (u *uri) Name() string { _ = "STUB: not implemented"; return "" }

func (u *uri) MimeType() string { _ = "STUB: not implemented"; return "" }

func (u *uri) Scheme() string { _ = "STUB: not implemented"; return "" }

func (u *uri) Authority() string { _ = "STUB: not implemented"; return "" }

func (u *uri) Path() string { _ = "STUB: not implemented"; return "" }

func (u *uri) Query() string { _ = "STUB: not implemented"; return "" }

func (u *uri) Fragment() string { _ = "STUB: not implemented"; return "" }
