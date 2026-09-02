package commands

import "fyne.io/fyne/v2/cmd/fyne/internal/commands"

type Command interface {
	AddFlags()
	PrintHelp(string)
	Run(args []string)
}

type Getter = commands.Getter

func NewGetter() *Getter { _ = "STUB: not implemented"; return nil }

func NewBundler() Command { _ = "STUB: not implemented"; return *new(Command) }

func NewInstaller() Command { _ = "STUB: not implemented"; return *new(Command) }

func NewPackager() Command { _ = "STUB: not implemented"; return *new(Command) }

func NewReleaser() Command { _ = "STUB: not implemented"; return *new(Command) }
