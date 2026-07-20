//go:build !android

package driver

func RunNative(fn func(any) error) error { _ = "STUB: not implemented"; return nil }
