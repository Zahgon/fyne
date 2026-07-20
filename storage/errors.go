package storage

import "errors"

var (
	ErrAlreadyExists = errors.New("document already exists")

	ErrNotExists = errors.New("document does not exist")
)
