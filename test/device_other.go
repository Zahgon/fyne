//go:build !mobile

package test

func (d *device) IsMobile() bool { _ = "STUB: not implemented"; return false }
