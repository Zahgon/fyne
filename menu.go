package fyne

type systemTrayDriver interface {
	Driver
	SetSystemTrayMenu(*Menu)
	SystemTrayMenu() *Menu
}

type Menu struct {
	Label string
	Items []*MenuItem
}

func NewMenu(label string, items ...*MenuItem) *Menu { _ = "STUB: not implemented"; return nil }

func (m *Menu) Refresh() { _ = "STUB: not implemented"; return }

type MenuItem struct {
	ChildMenu *Menu

	IsQuit      bool
	IsSeparator bool
	Label       string
	Action      func() `json:"-"`

	Disabled bool

	Checked bool

	Icon Resource

	Shortcut Shortcut
}

func NewMenuItem(label string, action func()) *MenuItem { _ = "STUB: not implemented"; return nil }

func NewMenuItemWithIcon(label string, icon Resource, action func()) *MenuItem {
	_ = "STUB: not implemented"
	return nil
}

func NewMenuItemSeparator() *MenuItem { _ = "STUB: not implemented"; return nil }

type MainMenu struct {
	Items []*Menu
}

func NewMainMenu(items ...*Menu) *MainMenu { _ = "STUB: not implemented"; return nil }

func (m *MainMenu) Refresh() { _ = "STUB: not implemented"; return }
