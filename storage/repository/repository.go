package repository

import (
	"fyne.io/fyne/v2"
)

var repositoryTable = map[string]Repository{}

type Repository interface {
	Exists(u fyne.URI) (bool, error)

	Reader(u fyne.URI) (fyne.URIReadCloser, error)

	CanRead(u fyne.URI) (bool, error)

	Destroy(string)
}

type CustomURIRepository interface {
	Repository

	ParseURI(string) (fyne.URI, error)
}

type WritableRepository interface {
	Repository

	Writer(u fyne.URI) (fyne.URIWriteCloser, error)

	CanWrite(u fyne.URI) (bool, error)

	Delete(u fyne.URI) error
}

type AppendableRepository interface {
	WritableRepository

	Appender(u fyne.URI) (fyne.URIWriteCloser, error)
}

type ListableRepository interface {
	Repository

	CanList(u fyne.URI) (bool, error)

	List(u fyne.URI) ([]fyne.URI, error)

	CreateListable(u fyne.URI) error
}

type HierarchicalRepository interface {
	Repository

	Parent(fyne.URI) (fyne.URI, error)

	Child(fyne.URI, string) (fyne.URI, error)
}

type DeleteAllRepository interface {
	WritableRepository

	DeleteAll(fyne.URI) error
}

type CopyableRepository interface {
	Repository

	Copy(fyne.URI, fyne.URI) error
}

type MovableRepository interface {
	Repository

	Move(fyne.URI, fyne.URI) error
}

func Register(scheme string, repository Repository) { _ = "STUB: not implemented"; return }

func ForURI(u fyne.URI) (Repository, error) {
	_ = "STUB: not implemented"
	return *new(Repository), nil
}

func ForScheme(scheme string) (Repository, error) {
	_ = "STUB: not implemented"
	return *new(Repository), nil
}
