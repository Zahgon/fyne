//go:build !mobile

package test

func (*device) IsMobile() bool { _ = "STUB: not implemented"; return false }
