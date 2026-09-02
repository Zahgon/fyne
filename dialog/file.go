package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type ViewLayout int

const (
	defaultView ViewLayout = iota
	ListView
	GridView
)

const (
	folderDesktop   = "Desktop"
	folderDocuments = "Documents"
	folderDownloads = "Downloads"
	folderHome      = "Home"
	folderMusic     = "Music"
	folderPictures  = "Pictures"

	lastFolderKey = "fyne:fileDialogLastFolder"
	viewLayoutKey = "fyne:fileDialogViewLayout"
)

type textWidget interface {
	fyne.Widget
	SetText(string)
}

type favoriteItem struct {
	locName string
	locIcon fyne.Resource
	loc     fyne.URI
}

type fileDialogPanel interface {
	fyne.Widget

	Unselect(int)
}

type fileDialog struct {
	file             *FileDialog
	fileName         textWidget
	title            *widget.Label
	dismiss          *widget.Button
	open             *widget.Button
	breadcrumb       *fyne.Container
	breadcrumbScroll *container.Scroll
	files            fileDialogPanel
	filesScroll      *container.Scroll
	favorites        []favoriteItem
	favoritesList    *widget.List
	showHidden       bool

	view ViewLayout

	data []fyne.URI

	win        *widget.PopUp
	selected   fyne.URI
	selectedID int
	dir        fyne.ListableURI

	initialFileName string

	toggleViewButton *widget.Button
}

type FileDialog struct {
	callback         any
	onClosedCallback func(bool)
	parent           fyne.Window
	dialog           *fileDialog

	titleText                string
	confirmText, dismissText string
	desiredSize              fyne.Size
	filter                   storage.FileFilter
	save                     bool

	startingLocation fyne.ListableURI

	initialFileName string

	initialView ViewLayout
}

var _ Dialog = (*FileDialog)(nil)

func (f *fileDialog) makeUI() fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func (f *fileDialog) makeOpenButton(label string) *widget.Button {
	_ = "STUB: not implemented"
	return nil
}

func (f *fileDialog) makeDismissButton(label string) *widget.Button {
	_ = "STUB: not implemented"
	return nil
}

func (f *fileDialog) optionsMenu(position fyne.Position, buttonSize fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func getFavoriteLocations() (map[string]fyne.ListableURI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *fileDialog) loadFavorites() { _ = "STUB: not implemented"; return }

func (f *fileDialog) refreshDir(dir fyne.ListableURI) { _ = "STUB: not implemented"; return }

func (f *fileDialog) setLocation(dir fyne.URI) { _ = "STUB: not implemented"; return }

func (f *fileDialog) setSelected(file fyne.URI, id int) { _ = "STUB: not implemented"; return }

func (f *fileDialog) setView(view ViewLayout) { _ = "STUB: not implemented"; return }

func (f *fileDialog) getDataItem(id int) (fyne.URI, bool) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), false
}

func (f *FileDialog) effectiveStartingDir() fyne.ListableURI {
	_ = "STUB: not implemented"
	return *new(fyne.ListableURI)
}

func showFile(file *FileDialog) *fileDialog { _ = "STUB: not implemented"; return nil }

func (f *FileDialog) Dismiss() { _ = "STUB: not implemented"; return }

func (f *FileDialog) MinSize() fyne.Size { _ = "STUB: not implemented"; return *new(fyne.Size) }

func (f *FileDialog) Show() { _ = "STUB: not implemented"; return }

func (f *FileDialog) Refresh() { _ = "STUB: not implemented"; return }

func (f *FileDialog) Resize(size fyne.Size) { _ = "STUB: not implemented"; return }

func (f *FileDialog) Hide() { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetConfirmText(label string) { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetDismissText(label string) { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetTitleText(label string) { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetLocation(u fyne.ListableURI) { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetOnClosed(closed func()) { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetFilter(filter storage.FileFilter) { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetFileName(fileName string) { _ = "STUB: not implemented"; return }

func (f *FileDialog) SetView(v ViewLayout) { _ = "STUB: not implemented"; return }

func NewFileOpen(callback func(reader fyne.URIReadCloser, err error), parent fyne.Window) *FileDialog {
	_ = "STUB: not implemented"
	return nil
}

func NewFileSave(callback func(writer fyne.URIWriteCloser, err error), parent fyne.Window) *FileDialog {
	_ = "STUB: not implemented"
	return nil
}

func ShowFileOpen(callback func(reader fyne.URIReadCloser, err error), parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}

func ShowFileSave(callback func(writer fyne.URIWriteCloser, err error), parent fyne.Window) {
	_ = "STUB: not implemented"
	return
}

func getFavoritesIcon(location string) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func getFavoritesOrder() []string { _ = "STUB: not implemented"; return nil }

func hasAppFiles(a fyne.App) bool { _ = "STUB: not implemented"; return false }

func storageURI(a fyne.App) fyne.URI { _ = "STUB: not implemented"; return *new(fyne.URI) }

type iconPaddingLayout struct{}

func (*iconPaddingLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	_ = "STUB: not implemented"
	return
}

func (*iconPaddingLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	_ = "STUB: not implemented"
	return *new(fyne.Size)
}
