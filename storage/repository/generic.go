package repository

import (
	"net/url"

	"fyne.io/fyne/v2"
)

func getUserHost(authority string) (*url.Userinfo, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func GenericParent(u fyne.URI) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}

func GenericChild(u fyne.URI, component string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}

func GenericCopy(source fyne.URI, destination fyne.URI) error {
	_ = "STUB: not implemented"
	return nil
}

func GenericDeleteAll(u fyne.URI) error { _ = "STUB: not implemented"; return nil }

func GenericMove(source fyne.URI, destination fyne.URI) error {
	_ = "STUB: not implemented"
	return nil
}

func genericCopyMoveListable(source, destination fyne.URI, repo Repository, deleteSource bool) error {
	_ = "STUB: not implemented"
	return nil
}

func genericDeleteAll(u fyne.URI, wrepo WritableRepository, lrepo ListableRepository) error {
	_ = "STUB: not implemented"
	return nil
}
