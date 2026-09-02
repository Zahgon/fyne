package test

import (
	"fyne.io/fyne/v2"
)

type mockCloud struct {
	configured bool
}

func (c *mockCloud) Cleanup(_ fyne.App) { _ = "STUB: not implemented"; return }

func (*mockCloud) ProviderDescription() string { _ = "STUB: not implemented"; return "" }

func (*mockCloud) ProviderIcon() fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (*mockCloud) ProviderName() string { _ = "STUB: not implemented"; return "" }

func (c *mockCloud) Setup(_ fyne.App) error { _ = "STUB: not implemented"; return nil }
