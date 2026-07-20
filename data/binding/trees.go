package binding

import (
	"fyne.io/fyne/v2"
)

const DataTreeRootID = ""

type Tree[T any] interface {
	DataTree

	Append(parent, id string, value T) error
	Get() (map[string][]string, map[string]T, error)
	GetValue(id string) (T, error)
	Prepend(parent, id string, value T) error
	Remove(id string) error
	Set(ids map[string][]string, values map[string]T) error
	SetValue(id string, value T) error
}

type ExternalTree[T any] interface {
	Tree[T]

	Reload() error
}

func NewTree[T any](comparator func(T, T) bool) Tree[T] { _ = "STUB: not implemented"; return nil }

func BindTree[T any](ids *map[string][]string, v *map[string]T, comparator func(T, T) bool) ExternalTree[T] {
	_ = "STUB: not implemented"
	return nil
}

type DataTree interface {
	DataItem
	GetItem(id string) (DataItem, error)
	ChildIDs(string) []string
}

type BoolTree = Tree[bool]

type ExternalBoolTree = ExternalTree[bool]

func NewBoolTree() Tree[bool] { _ = "STUB: not implemented"; return nil }

func BindBoolTree(ids *map[string][]string, v *map[string]bool) ExternalTree[bool] {
	_ = "STUB: not implemented"
	return nil
}

type BytesTree = Tree[[]byte]

type ExternalBytesTree = ExternalTree[[]byte]

func NewBytesTree() Tree[[]byte] { _ = "STUB: not implemented"; return nil }

func BindBytesTree(ids *map[string][]string, v *map[string][]byte) ExternalTree[[]byte] {
	_ = "STUB: not implemented"
	return nil
}

type FloatTree = Tree[float64]

type ExternalFloatTree = ExternalTree[float64]

func NewFloatTree() Tree[float64] { _ = "STUB: not implemented"; return nil }

func BindFloatTree(ids *map[string][]string, v *map[string]float64) ExternalTree[float64] {
	_ = "STUB: not implemented"
	return nil
}

type IntTree = Tree[int]

type ExternalIntTree = ExternalTree[int]

func NewIntTree() Tree[int] { _ = "STUB: not implemented"; return nil }

func BindIntTree(ids *map[string][]string, v *map[string]int) ExternalTree[int] {
	_ = "STUB: not implemented"
	return nil
}

type RuneTree = Tree[rune]

type ExternalRuneTree = ExternalTree[rune]

func NewRuneTree() Tree[rune] { _ = "STUB: not implemented"; return nil }

func BindRuneTree(ids *map[string][]string, v *map[string]rune) ExternalTree[rune] {
	_ = "STUB: not implemented"
	return nil
}

type StringTree = Tree[string]

type ExternalStringTree = ExternalTree[string]

func NewStringTree() Tree[string] { _ = "STUB: not implemented"; return nil }

func BindStringTree(ids *map[string][]string, v *map[string]string) ExternalTree[string] {
	_ = "STUB: not implemented"
	return nil
}

type UntypedTree = Tree[any]

type ExternalUntypedTree = ExternalTree[any]

func NewUntypedTree() Tree[any] { _ = "STUB: not implemented"; return nil }

func BindUntypedTree(ids *map[string][]string, v *map[string]any) ExternalTree[any] {
	_ = "STUB: not implemented"
	return nil
}

type URITree = Tree[fyne.URI]

type ExternalURITree = ExternalTree[fyne.URI]

func NewURITree() Tree[fyne.URI] { _ = "STUB: not implemented"; return nil }

func BindURITree(ids *map[string][]string, v *map[string]fyne.URI) ExternalTree[fyne.URI] {
	_ = "STUB: not implemented"
	return nil
}

type treeBase struct {
	base

	ids   map[string][]string
	items map[string]DataItem
}

func (t *treeBase) GetItem(id string) (DataItem, error) {
	_ = "STUB: not implemented"
	return *new(DataItem), nil
}

func (t *treeBase) ChildIDs(id string) []string { _ = "STUB: not implemented"; return nil }

func (t *treeBase) appendItem(i DataItem, id, parent string) { _ = "STUB: not implemented"; return }

func (t *treeBase) deleteItem(id, parent string) { _ = "STUB: not implemented"; return }

func parentIDFor(id string, ids map[string][]string) string { _ = "STUB: not implemented"; return "" }

func newTree[T any](comparator func(T, T) bool) *boundTree[T] {
	_ = "STUB: not implemented"
	return nil
}

func newTreeComparable[T comparable]() *boundTree[T] { _ = "STUB: not implemented"; return nil }

func bindTree[T any](ids *map[string][]string, v *map[string]T, comparator func(T, T) bool) *boundTree[T] {
	_ = "STUB: not implemented"
	return nil
}

func bindTreeComparable[T comparable](ids *map[string][]string, v *map[string]T) *boundTree[T] {
	_ = "STUB: not implemented"
	return nil
}

type boundTree[T any] struct {
	treeBase

	comparator     func(T, T) bool
	val            *map[string]T
	updateExternal bool
}

func (t *boundTree[T]) Append(parent, id string, val T) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *boundTree[T]) Get() (map[string][]string, map[string]T, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (t *boundTree[T]) GetValue(id string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (t *boundTree[T]) Prepend(parent, id string, val T) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *boundTree[T]) Remove(id string) error { _ = "STUB: not implemented"; return nil }

func (t *boundTree[T]) removeChildren(id string) { _ = "STUB: not implemented"; return }

func (t *boundTree[T]) Reload() error { _ = "STUB: not implemented"; return nil }

func (t *boundTree[T]) Set(ids map[string][]string, v map[string]T) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *boundTree[T]) doReload() (fire bool, retErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *boundTree[T]) SetValue(id string, v T) error { _ = "STUB: not implemented"; return nil }

func bindTreeItem[T any](v *map[string]T, id string, external bool, comparator func(T, T) bool) Item[T] {
	_ = "STUB: not implemented"
	return nil
}

type boundTreeItem[T any] struct {
	base

	val *map[string]T
	id  string
}

func (t *boundTreeItem[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (t *boundTreeItem[T]) Set(val T) error { _ = "STUB: not implemented"; return nil }

func (t *boundTreeItem[T]) doSet(val T) error { _ = "STUB: not implemented"; return nil }

type boundExternalTreeItem[T any] struct {
	boundTreeItem[T]

	comparator func(T, T) bool
	old        T
}

func (t *boundExternalTreeItem[T]) setIfChanged(val T) error { _ = "STUB: not implemented"; return nil }
