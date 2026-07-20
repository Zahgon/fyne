package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/internal/async"
	"fyne.io/fyne/v2/internal/widget"
)

type TreeNodeID = string

const (
	allTreeNodesID     TreeNodeID = "_ALLNODES"
	onlyNewTreeNodesID TreeNodeID = "_ONLYNEWNODES"
)

var (
	_ fyne.Focusable = (*Tree)(nil)
	_ fyne.Widget    = (*Tree)(nil)
)

type Tree struct {
	BaseWidget
	Root TreeNodeID

	HideSeparators bool

	ChildUIDs      func(uid TreeNodeID) (c []TreeNodeID)                     `json:"-"`
	CreateNode     func(branch bool) (o fyne.CanvasObject)                   `json:"-"`
	IsBranch       func(uid TreeNodeID) (ok bool)                            `json:"-"`
	OnBranchClosed func(uid TreeNodeID)                                      `json:"-"`
	OnBranchOpened func(uid TreeNodeID)                                      `json:"-"`
	OnSelected     func(uid TreeNodeID)                                      `json:"-"`
	OnUnselected   func(uid TreeNodeID)                                      `json:"-"`
	UpdateNode     func(uid TreeNodeID, branch bool, node fyne.CanvasObject) `json:"-"`

	OnHighlighted func(id TreeNodeID) `json:"-"`

	branchMinSize    fyne.Size
	currentHighlight TreeNodeID
	focused          bool
	leafMinSize      fyne.Size
	offset           fyne.Position
	open             map[TreeNodeID]bool
	scroller         *widget.Scroll
	selected         []TreeNodeID
}

func NewTree(childUIDs func(TreeNodeID) []TreeNodeID, isBranch func(TreeNodeID) bool, create func(bool) fyne.CanvasObject, update func(TreeNodeID, bool, fyne.CanvasObject)) *Tree {
	_ = "STUB: not implemented"
	return nil
}

func NewTreeWithData(data binding.DataTree, createItem func(bool) fyne.CanvasObject, updateItem func(binding.DataItem, bool, fyne.CanvasObject)) *Tree {
	_ = "STUB: not implemented"
	return nil
}

func NewTreeWithStrings(data map[string][]string) (t *Tree) { _ = "STUB: not implemented"; return nil }

func (t *Tree) CloseAllBranches() { _ = "STUB: not implemented"; return }

func (t *Tree) CloseBranch(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (t *Tree) IsBranchOpen(uid TreeNodeID) bool { _ = "STUB: not implemented"; return false }

func (t *Tree) FocusGained() { _ = "STUB: not implemented"; return }

func (t *Tree) FocusLost() { _ = "STUB: not implemented"; return }

func (t *Tree) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (t *Tree) RefreshItem(id TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) OpenAllBranches() { _ = "STUB: not implemented"; return }

func (t *Tree) OpenBranch(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) openBranches(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (t *Tree) Highlight(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) ScrollToBottom() { _ = "STUB: not implemented"; return }

func (t *Tree) ScrollTo(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) ScrollToOffset(offset float32) { _ = "STUB: not implemented"; return }

func (t *Tree) ScrollToTop() { _ = "STUB: not implemented"; return }

func (t *Tree) Select(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) setItemFocus(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) ToggleBranch(uid string) { _ = "STUB: not implemented"; return }

func (t *Tree) TypedKey(event *fyne.KeyEvent) { _ = "STUB: not implemented"; return }

func (t *Tree) TypedRune(_ rune) { _ = "STUB: not implemented"; return }

func (t *Tree) Unselect(uid TreeNodeID) { _ = "STUB: not implemented"; return }

func (t *Tree) UnselectAll() { _ = "STUB: not implemented"; return }

func (t *Tree) findPath(current, target TreeNodeID) (bool, []TreeNodeID) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *Tree) ensureOpenMap() { _ = "STUB: not implemented"; return }

func (t *Tree) offsetAndSize(uid TreeNodeID) (y float32, size fyne.Size, found bool) {
	_ = "STUB: not implemented"
	return 0, *new(fyne.Size), false
}

func (t *Tree) offsetUpdated(pos fyne.Position) { _ = "STUB: not implemented"; return }

func (t *Tree) walk(uid, parent TreeNodeID, depth int, walkClosedBranch bool, onNode func(TreeNodeID, TreeNodeID, bool, int)) {
	_ = "STUB: not implemented"
	return
}

func (t *Tree) walkAll(onNode func(TreeNodeID, TreeNodeID, bool, int)) {
	_ = "STUB: not implemented"
	return
}

var _ fyne.WidgetRenderer = (*treeRenderer)(nil)

type treeRenderer struct {
	widget.BaseRenderer
	tree     *Tree
	content  *treeContent
	scroller *widget.Scroll
}

func (r *treeRenderer) MinSize() (min fyne.Size) { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (r *treeRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *treeRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *treeRenderer) updateMinSizes() { _ = "STUB: not implemented"; return }

var _ fyne.Widget = (*treeContent)(nil)

type treeContent struct {
	BaseWidget
	tree     *Tree
	viewport fyne.Size

	nextRefreshID TreeNodeID
}

func newTreeContent(tree *Tree) (c *treeContent) { _ = "STUB: not implemented"; return nil }

func (c *treeContent) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (c *treeContent) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (c *treeContent) refreshForID(id TreeNodeID) { _ = "STUB: not implemented"; return }

func (c *treeContent) Refresh() { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*treeContentRenderer)(nil)

type treeContentRenderer struct {
	widget.BaseRenderer
	treeContent *treeContent
	separators  []fyne.CanvasObject
	objects     []fyne.CanvasObject
	branches    map[string]*branch
	leaves      map[string]*leaf
	branchPool  async.Pool[fyne.CanvasObject]
	leafPool    async.Pool[fyne.CanvasObject]

	wasVisible   []TreeNodeID
	visible      []TreeNodeID
	minSizeCache fyne.Size
}

func (r *treeContentRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *treeContentRenderer) MinSize() (min fyne.Size) {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *treeContentRenderer) Objects() []fyne.CanvasObject { _ = "STUB: not implemented"; return nil }

func (r *treeContentRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *treeContentRenderer) refreshForID(toDraw TreeNodeID) { _ = "STUB: not implemented"; return }

func (r *treeContentRenderer) getBranch() (b *branch) { _ = "STUB: not implemented"; return nil }

func (r *treeContentRenderer) getLeaf() (l *leaf) { _ = "STUB: not implemented"; return nil }

var (
	_ desktop.Hoverable = (*treeNode)(nil)
	_ fyne.CanvasObject = (*treeNode)(nil)
	_ fyne.Tappable     = (*treeNode)(nil)
)

type treeNode struct {
	BaseWidget
	tree     *Tree
	uid      string
	depth    int
	hovered  bool
	icon     fyne.CanvasObject
	isBranch bool
	content  fyne.CanvasObject
}

func (n *treeNode) Content() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (n *treeNode) CreateRenderer() fyne.WidgetRenderer {
	_ = "STUB: not implemented"
	return *new(fyne.WidgetRenderer)
}

func (n *treeNode) Indent() float32 { _ = "STUB: not implemented"; return 0 }

func (n *treeNode) MouseIn(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (n *treeNode) MouseMoved(*desktop.MouseEvent) { _ = "STUB: not implemented"; return }

func (n *treeNode) MouseOut() { _ = "STUB: not implemented"; return }

func (n *treeNode) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (n *treeNode) partialRefresh() { _ = "STUB: not implemented"; return }

func (n *treeNode) update(uid string, depth int) { _ = "STUB: not implemented"; return }

var _ fyne.WidgetRenderer = (*treeNodeRenderer)(nil)

type treeNodeRenderer struct {
	widget.BaseRenderer
	treeNode   *treeNode
	background *canvas.Rectangle
}

func (r *treeNodeRenderer) Layout(size fyne.Size) { _ = "STUB: not implemented"; return }

func (r *treeNodeRenderer) MinSize() (min fyne.Size) {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}

func (r *treeNodeRenderer) Objects() (objects []fyne.CanvasObject) {
	_ = "STUB: not implemented"
	return nil
}

func (r *treeNodeRenderer) Refresh() { _ = "STUB: not implemented"; return }

func (r *treeNodeRenderer) partialRefresh() { _ = "STUB: not implemented"; return }

var _ fyne.Widget = (*branch)(nil)

type branch struct {
	*treeNode
}

func newBranch(tree *Tree, content fyne.CanvasObject) (b *branch) {
	_ = "STUB: not implemented"
	return nil
}

func (b *branch) update(uid string, depth int) { _ = "STUB: not implemented"; return }

var _ fyne.Tappable = (*branchIcon)(nil)

type branchIcon struct {
	Icon
	tree *Tree
	uid  string
}

func newBranchIcon(tree *Tree) (i *branchIcon) { _ = "STUB: not implemented"; return nil }

func (i *branchIcon) Refresh() { _ = "STUB: not implemented"; return }

func (i *branchIcon) Tapped(*fyne.PointEvent) { _ = "STUB: not implemented"; return }

func (i *branchIcon) update(uid string) { _ = "STUB: not implemented"; return }

var _ fyne.Widget = (*leaf)(nil)

type leaf struct {
	*treeNode
}

func newLeaf(tree *Tree, content fyne.CanvasObject) (l *leaf) {
	_ = "STUB: not implemented"
	return nil
}

func contains(slice []string, item string) bool { _ = "STUB: not implemented"; return false }
