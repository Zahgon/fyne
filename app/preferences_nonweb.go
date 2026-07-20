//go:build !wasm

package app

import (
	"io"
)

func (p *preferences) storageWriter() (writeSyncCloser, error) {
	_ = "STUB: not implemented"
	return *new(writeSyncCloser), nil
}

func (p *preferences) storageReader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (p *preferences) storageWriterForPath(path string) (writeSyncCloser, error) {
	_ = "STUB: not implemented"
	return *new(writeSyncCloser), nil
}

func (p *preferences) storageReaderForPath(path string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (p *preferences) saveToFile(path string) error { _ = "STUB: not implemented"; return nil }

func (p *preferences) loadFromFile(path string) error { _ = "STUB: not implemented"; return nil }
