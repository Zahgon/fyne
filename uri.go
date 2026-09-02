package fyne

import (
	"fmt"
	"io"
)

const (
	URIAuthorityPrefix = "//"
	URIPathSeparator   = "/"
	URISchemeFile      = "file"
	URISchemeSeparator = ":"
)

type URIReadCloser interface {
	io.ReadCloser

	URI() URI
}

type URIWriteCloser interface {
	io.WriteCloser

	URI() URI
}

type URI interface {
	fmt.Stringer

	Extension() string

	Name() string

	MimeType() string

	Scheme() string

	Authority() string

	Path() string

	Query() string

	Fragment() string
}

type ListableURI interface {
	URI

	List() ([]URI, error)
}

type URIWithIcon interface {
	URI

	Icon() Resource
}
