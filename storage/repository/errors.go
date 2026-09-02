package repository

import (
	"errors"
)

var (
	ErrOperationNotSupported = errors.New("operation not supported for this URI")

	ErrURIRoot = errors.New("cannot take the parent of the root element in a URI")
)
