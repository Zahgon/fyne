package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
)

const (
	IconNameCancel fyne.ThemeIconName = "cancel"

	IconNameConfirm fyne.ThemeIconName = "confirm"

	IconNameDelete fyne.ThemeIconName = "delete"

	IconNameSearch fyne.ThemeIconName = "search"

	IconNameSearchReplace fyne.ThemeIconName = "searchReplace"

	IconNameMenu fyne.ThemeIconName = "menu"

	IconNameMenuExpand fyne.ThemeIconName = "menuExpand"

	IconNameCheckButton fyne.ThemeIconName = "unchecked"

	IconNameCheckButtonChecked fyne.ThemeIconName = "checked"

	IconNameCheckButtonFill fyne.ThemeIconName = "iconNameCheckButtonFill"

	IconNameCheckButtonPartial fyne.ThemeIconName = "partial"

	IconNameRadioButton fyne.ThemeIconName = "radioButton"

	IconNameRadioButtonChecked fyne.ThemeIconName = "radioButtonChecked"

	IconNameRadioButtonFill fyne.ThemeIconName = "iconNameRadioButtonFill"

	IconNameColorAchromatic fyne.ThemeIconName = "colorAchromatic"

	IconNameColorChromatic fyne.ThemeIconName = "colorChromatic"

	IconNameColorPalette fyne.ThemeIconName = "colorPalette"

	IconNameContentAdd fyne.ThemeIconName = "contentAdd"

	IconNameContentRemove fyne.ThemeIconName = "contentRemove"

	IconNameContentCut fyne.ThemeIconName = "contentCut"

	IconNameContentCopy fyne.ThemeIconName = "contentCopy"

	IconNameContentPaste fyne.ThemeIconName = "contentPaste"

	IconNameContentClear fyne.ThemeIconName = "contentClear"

	IconNameContentRedo fyne.ThemeIconName = "contentRedo"

	IconNameContentUndo fyne.ThemeIconName = "contentUndo"

	IconNameInfo fyne.ThemeIconName = "info"

	IconNameQuestion fyne.ThemeIconName = "question"

	IconNameWarning fyne.ThemeIconName = "warning"

	IconNameError fyne.ThemeIconName = "error"

	IconNameBrokenImage fyne.ThemeIconName = "broken-image"

	IconNameDocument fyne.ThemeIconName = "document"

	IconNameDocumentCreate fyne.ThemeIconName = "documentCreate"

	IconNameDocumentPrint fyne.ThemeIconName = "documentPrint"

	IconNameDocumentSave fyne.ThemeIconName = "documentSave"

	IconNameDragCornerIndicator fyne.ThemeIconName = "dragCornerIndicator"

	IconNameMoreHorizontal fyne.ThemeIconName = "moreHorizontal"

	IconNameMoreVertical fyne.ThemeIconName = "moreVertical"

	IconNameMailAttachment fyne.ThemeIconName = "mailAttachment"

	IconNameMailCompose fyne.ThemeIconName = "mailCompose"

	IconNameMailForward fyne.ThemeIconName = "mailForward"

	IconNameMailReply fyne.ThemeIconName = "mailReply"

	IconNameMailReplyAll fyne.ThemeIconName = "mailReplyAll"

	IconNameMailSend fyne.ThemeIconName = "mailSend"

	IconNameMediaMusic fyne.ThemeIconName = "mediaMusic"

	IconNameMediaPhoto fyne.ThemeIconName = "mediaPhoto"

	IconNameMediaVideo fyne.ThemeIconName = "mediaVideo"

	IconNameMediaFastForward fyne.ThemeIconName = "mediaFastForward"

	IconNameMediaFastRewind fyne.ThemeIconName = "mediaFastRewind"

	IconNameMediaPause fyne.ThemeIconName = "mediaPause"

	IconNameMediaPlay fyne.ThemeIconName = "mediaPlay"

	IconNameMediaRecord fyne.ThemeIconName = "mediaRecord"

	IconNameMediaReplay fyne.ThemeIconName = "mediaReplay"

	IconNameMediaSkipNext fyne.ThemeIconName = "mediaSkipNext"

	IconNameMediaSkipPrevious fyne.ThemeIconName = "mediaSkipPrevious"

	IconNameMediaStop fyne.ThemeIconName = "mediaStop"

	IconNameMoveDown fyne.ThemeIconName = "arrowDown"

	IconNameMoveUp fyne.ThemeIconName = "arrowUp"

	IconNameNavigateBack fyne.ThemeIconName = "arrowBack"

	IconNameNavigateNext fyne.ThemeIconName = "arrowForward"

	IconNameArrowDropDown fyne.ThemeIconName = "arrowDropDown"

	IconNameArrowDropUp fyne.ThemeIconName = "arrowDropUp"

	IconNameFile fyne.ThemeIconName = "file"

	IconNameFileApplication fyne.ThemeIconName = "fileApplication"

	IconNameFileAudio fyne.ThemeIconName = "fileAudio"

	IconNameFileImage fyne.ThemeIconName = "fileImage"

	IconNameFileText fyne.ThemeIconName = "fileText"

	IconNameFileVideo fyne.ThemeIconName = "fileVideo"

	IconNameFolder fyne.ThemeIconName = "folder"

	IconNameFolderNew fyne.ThemeIconName = "folderNew"

	IconNameFolderOpen fyne.ThemeIconName = "folderOpen"

	IconNameHelp fyne.ThemeIconName = "help"

	IconNameHistory fyne.ThemeIconName = "history"

	IconNameHome fyne.ThemeIconName = "home"

	IconNameSettings fyne.ThemeIconName = "settings"

	IconNameStorage fyne.ThemeIconName = "storage"

	IconNameUpload fyne.ThemeIconName = "upload"

	IconNameViewFullScreen fyne.ThemeIconName = "viewFullScreen"

	IconNameViewRefresh fyne.ThemeIconName = "viewRefresh"

	IconNameViewZoomFit fyne.ThemeIconName = "viewZoomFit"

	IconNameViewZoomIn fyne.ThemeIconName = "viewZoomIn"

	IconNameViewZoomOut fyne.ThemeIconName = "viewZoomOut"

	IconNameViewRestore fyne.ThemeIconName = "viewRestore"

	IconNameVisibility fyne.ThemeIconName = "visibility"

	IconNameVisibilityOff fyne.ThemeIconName = "visibilityOff"

	IconNameVolumeDown fyne.ThemeIconName = "volumeDown"

	IconNameVolumeMute fyne.ThemeIconName = "volumeMute"

	IconNameVolumeUp fyne.ThemeIconName = "volumeUp"

	IconNameDownload fyne.ThemeIconName = "download"

	IconNameComputer fyne.ThemeIconName = "computer"

	IconNameDesktop fyne.ThemeIconName = "desktop"

	IconNameAccount fyne.ThemeIconName = "account"

	IconNameCalendar fyne.ThemeIconName = "calendar"

	IconNameLogin fyne.ThemeIconName = "login"

	IconNameLogout fyne.ThemeIconName = "logout"

	IconNameList fyne.ThemeIconName = "list"

	IconNameGrid fyne.ThemeIconName = "grid"

	IconNameWindowClose fyne.ThemeIconName = "windowClose"

	IconNameWindowMaximize fyne.ThemeIconName = "windowMaximize"

	IconNameWindowMinimize fyne.ThemeIconName = "windowMinimize"
)

var icons = map[fyne.ThemeIconName]fyne.Resource{
	IconNameCancel:        NewThemedResource(cancelIconRes),
	IconNameConfirm:       NewThemedResource(checkIconRes),
	IconNameDelete:        NewThemedResource(deleteIconRes),
	IconNameSearch:        NewThemedResource(searchIconRes),
	IconNameSearchReplace: NewThemedResource(searchreplaceIconRes),
	IconNameMenu:          NewThemedResource(menuIconRes),
	IconNameMenuExpand:    NewThemedResource(menuexpandIconRes),

	IconNameCheckButton:        NewThemedResource(checkboxIconRes),
	IconNameCheckButtonChecked: NewThemedResource(checkboxcheckedIconRes),
	IconNameCheckButtonFill:    NewThemedResource(checkboxfillIconRes),
	IconNameCheckButtonPartial: NewThemedResource(checkboxpartialIconRes),
	IconNameRadioButton:        NewThemedResource(radiobuttonIconRes),
	IconNameRadioButtonChecked: NewThemedResource(radiobuttoncheckedIconRes),
	IconNameRadioButtonFill:    NewThemedResource(radiobuttonfillIconRes),

	IconNameContentAdd:    NewThemedResource(contentaddIconRes),
	IconNameContentClear:  NewThemedResource(cancelIconRes),
	IconNameContentRemove: NewThemedResource(contentremoveIconRes),
	IconNameContentCut:    NewThemedResource(contentcutIconRes),
	IconNameContentCopy:   NewThemedResource(contentcopyIconRes),
	IconNameContentPaste:  NewThemedResource(contentpasteIconRes),
	IconNameContentRedo:   NewThemedResource(contentredoIconRes),
	IconNameContentUndo:   NewThemedResource(contentundoIconRes),

	IconNameColorAchromatic: NewThemedResource(colorachromaticIconRes),
	IconNameColorChromatic:  NewThemedResource(colorchromaticIconRes),
	IconNameColorPalette:    NewThemedResource(colorpaletteIconRes),

	IconNameDocument:       NewThemedResource(documentIconRes),
	IconNameDocumentCreate: NewThemedResource(documentcreateIconRes),
	IconNameDocumentPrint:  NewThemedResource(documentprintIconRes),
	IconNameDocumentSave:   NewThemedResource(documentsaveIconRes),

	IconNameDragCornerIndicator: NewThemedResource(dragcornerindicatorIconRes),

	IconNameMoreHorizontal: NewThemedResource(morehorizontalIconRes),
	IconNameMoreVertical:   NewThemedResource(moreverticalIconRes),

	IconNameInfo:        NewThemedResource(infoIconRes),
	IconNameQuestion:    NewThemedResource(questionIconRes),
	IconNameWarning:     NewThemedResource(warningIconRes),
	IconNameError:       NewThemedResource(errorIconRes),
	IconNameBrokenImage: NewThemedResource(brokenimageIconRes),

	IconNameMailAttachment: NewThemedResource(mailattachmentIconRes),
	IconNameMailCompose:    NewThemedResource(mailcomposeIconRes),
	IconNameMailForward:    NewThemedResource(mailforwardIconRes),
	IconNameMailReply:      NewThemedResource(mailreplyIconRes),
	IconNameMailReplyAll:   NewThemedResource(mailreplyallIconRes),
	IconNameMailSend:       NewThemedResource(mailsendIconRes),

	IconNameMediaMusic:        NewThemedResource(mediamusicIconRes),
	IconNameMediaPhoto:        NewThemedResource(mediaphotoIconRes),
	IconNameMediaVideo:        NewThemedResource(mediavideoIconRes),
	IconNameMediaFastForward:  NewThemedResource(mediafastforwardIconRes),
	IconNameMediaFastRewind:   NewThemedResource(mediafastrewindIconRes),
	IconNameMediaPause:        NewThemedResource(mediapauseIconRes),
	IconNameMediaPlay:         NewThemedResource(mediaplayIconRes),
	IconNameMediaRecord:       NewThemedResource(mediarecordIconRes),
	IconNameMediaReplay:       NewThemedResource(mediareplayIconRes),
	IconNameMediaSkipNext:     NewThemedResource(mediaskipnextIconRes),
	IconNameMediaSkipPrevious: NewThemedResource(mediaskippreviousIconRes),
	IconNameMediaStop:         NewThemedResource(mediastopIconRes),

	IconNameNavigateBack:  NewThemedResource(arrowbackIconRes),
	IconNameMoveDown:      NewThemedResource(arrowdownIconRes),
	IconNameNavigateNext:  NewThemedResource(arrowforwardIconRes),
	IconNameMoveUp:        NewThemedResource(arrowupIconRes),
	IconNameArrowDropDown: NewThemedResource(arrowdropdownIconRes),
	IconNameArrowDropUp:   NewThemedResource(arrowdropupIconRes),

	IconNameFile:            NewThemedResource(fileIconRes),
	IconNameFileApplication: NewThemedResource(fileapplicationIconRes),
	IconNameFileAudio:       NewThemedResource(fileaudioIconRes),
	IconNameFileImage:       NewThemedResource(fileimageIconRes),
	IconNameFileText:        NewThemedResource(filetextIconRes),
	IconNameFileVideo:       NewThemedResource(filevideoIconRes),
	IconNameFolder:          NewThemedResource(folderIconRes),
	IconNameFolderNew:       NewThemedResource(foldernewIconRes),
	IconNameFolderOpen:      NewThemedResource(folderopenIconRes),
	IconNameHelp:            NewThemedResource(helpIconRes),
	IconNameHistory:         NewThemedResource(historyIconRes),
	IconNameHome:            NewThemedResource(homeIconRes),
	IconNameSettings:        NewThemedResource(settingsIconRes),

	IconNameViewFullScreen: NewThemedResource(viewfullscreenIconRes),
	IconNameViewRefresh:    NewThemedResource(viewrefreshIconRes),
	IconNameViewRestore:    NewThemedResource(viewzoomfitIconRes),
	IconNameViewZoomFit:    NewThemedResource(viewzoomfitIconRes),
	IconNameViewZoomIn:     NewThemedResource(viewzoominIconRes),
	IconNameViewZoomOut:    NewThemedResource(viewzoomoutIconRes),

	IconNameVisibility:    NewThemedResource(visibilityIconRes),
	IconNameVisibilityOff: NewThemedResource(visibilityoffIconRes),

	IconNameVolumeDown: NewThemedResource(volumedownIconRes),
	IconNameVolumeMute: NewThemedResource(volumemuteIconRes),
	IconNameVolumeUp:   NewThemedResource(volumeupIconRes),

	IconNameDownload: NewThemedResource(downloadIconRes),
	IconNameComputer: NewThemedResource(computerIconRes),
	IconNameDesktop:  NewThemedResource(desktopIconRes),
	IconNameStorage:  NewThemedResource(storageIconRes),
	IconNameUpload:   NewThemedResource(uploadIconRes),

	IconNameAccount:  NewThemedResource(accountIconRes),
	IconNameCalendar: NewThemedResource(calendarIconRes),
	IconNameLogin:    NewThemedResource(loginIconRes),
	IconNameLogout:   NewThemedResource(logoutIconRes),

	IconNameList: NewThemedResource(listIconRes),
	IconNameGrid: NewThemedResource(gridIconRes),

	IconNameWindowClose:    NewThemedResource(cancelIconRes),
	IconNameWindowMaximize: NewThemedResource(maximizeIconRes),
	IconNameWindowMinimize: NewThemedResource(minimizeIconRes),
}

func Icon(name fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func IconForWidget(name fyne.ThemeIconName, w fyne.Widget) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (t *builtinTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

var _ fyne.ThemedResource = (*ThemedResource)(nil)

type ThemedResource struct {
	source fyne.Resource

	ColorName fyne.ThemeColorName
}

func NewColoredResource(src fyne.Resource, name fyne.ThemeColorName) *ThemedResource {
	_ = "STUB: not implemented"
	return nil
}

func NewSuccessThemedResource(src fyne.Resource) *ThemedResource {
	_ = "STUB: not implemented"
	return nil
}

func NewThemedResource(src fyne.Resource) *ThemedResource { _ = "STUB: not implemented"; return nil }

func NewWarningThemedResource(src fyne.Resource) *ThemedResource {
	_ = "STUB: not implemented"
	return nil
}

func (res *ThemedResource) Name() string { _ = "STUB: not implemented"; return "" }

func (res *ThemedResource) ThemeColorName() fyne.ThemeColorName {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeColorName)
}

func (res *ThemedResource) Content() []byte { _ = "STUB: not implemented"; return nil }

func (res *ThemedResource) Error() *ErrorThemedResource { _ = "STUB: not implemented"; return nil }

var _ fyne.ThemedResource = (*InvertedThemedResource)(nil)

type InvertedThemedResource struct {
	source fyne.Resource
}

func NewInvertedThemedResource(orig fyne.Resource) *InvertedThemedResource {
	_ = "STUB: not implemented"
	return nil
}

func (res *InvertedThemedResource) Name() string { _ = "STUB: not implemented"; return "" }

func (res *InvertedThemedResource) Content() []byte { _ = "STUB: not implemented"; return nil }

func (res *InvertedThemedResource) ThemeColorName() fyne.ThemeColorName {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeColorName)
}

func (res *InvertedThemedResource) Original() fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

var _ fyne.ThemedResource = (*ErrorThemedResource)(nil)

type ErrorThemedResource struct {
	source fyne.Resource
}

func NewErrorThemedResource(orig fyne.Resource) *ErrorThemedResource {
	_ = "STUB: not implemented"
	return nil
}

func (res *ErrorThemedResource) Name() string { _ = "STUB: not implemented"; return "" }

func (res *ErrorThemedResource) Content() []byte { _ = "STUB: not implemented"; return nil }

func (res *ErrorThemedResource) Original() fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (res *ErrorThemedResource) ThemeColorName() fyne.ThemeColorName {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeColorName)
}

var _ fyne.ThemedResource = (*PrimaryThemedResource)(nil)

type PrimaryThemedResource struct {
	source fyne.Resource
}

func NewPrimaryThemedResource(orig fyne.Resource) *PrimaryThemedResource {
	_ = "STUB: not implemented"
	return nil
}

func (res *PrimaryThemedResource) Name() string { _ = "STUB: not implemented"; return "" }

func (res *PrimaryThemedResource) Content() []byte { _ = "STUB: not implemented"; return nil }

func (res *PrimaryThemedResource) Original() fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func (res *PrimaryThemedResource) ThemeColorName() fyne.ThemeColorName {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeColorName)
}

var _ fyne.ThemedResource = (*DisabledResource)(nil)

type DisabledResource struct {
	source fyne.Resource
}

func (res *DisabledResource) Name() string { _ = "STUB: not implemented"; return "" }

func (res *DisabledResource) Content() []byte { _ = "STUB: not implemented"; return nil }

func (res *DisabledResource) ThemeColorName() fyne.ThemeColorName {
	_ = "STUB: not implemented"
	return *new(fyne.ThemeColorName)
}

func NewDisabledResource(res fyne.Resource) *DisabledResource {
	_ = "STUB: not implemented"
	return nil
}

func FyneLogo() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func CancelIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ConfirmIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func DeleteIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func SearchIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func SearchReplaceIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MenuIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MenuExpandIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func CheckButtonIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func CheckButtonCheckedIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func CheckButtonFillIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func RadioButtonIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func RadioButtonCheckedIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func RadioButtonFillIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentAddIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentRemoveIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentClearIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentCutIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentCopyIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentPasteIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentRedoIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ContentUndoIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ColorAchromaticIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ColorChromaticIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ColorPaletteIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func DocumentIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func DocumentCreateIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func DocumentPrintIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func DocumentSaveIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MoreHorizontalIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MoreVerticalIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func InfoIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func QuestionIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func WarningIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ErrorIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func BrokenImageIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FileIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FileApplicationIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FileAudioIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FileImageIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FileTextIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FileVideoIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FolderIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FolderNewIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func FolderOpenIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func HelpIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func HistoryIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func HomeIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func SettingsIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MailAttachmentIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MailComposeIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MailForwardIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MailReplyIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MailReplyAllIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MailSendIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaMusicIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaPhotoIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaVideoIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaFastForwardIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaFastRewindIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaPauseIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaPlayIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaRecordIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaReplayIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaSkipNextIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaSkipPreviousIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MediaStopIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MoveDownIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MoveUpIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func NavigateBackIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func NavigateNextIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MenuDropDownIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func MenuDropUpIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ViewFullScreenIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ViewRestoreIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ViewRefreshIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ZoomFitIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ZoomInIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ZoomOutIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func VisibilityIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func VisibilityOffIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func VolumeDownIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func VolumeMuteIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func VolumeUpIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ComputerIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func DesktopIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func DownloadIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func StorageIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func UploadIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func AccountIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func CalendarIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func LoginIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func LogoutIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func ListIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func GridIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func WindowCloseIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func WindowMaximizeIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func WindowMinimizeIcon() fyne.Resource { _ = "STUB: not implemented"; return *new(fyne.Resource) }

func safeIconLookup(n fyne.ThemeIconName) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func unwrapResource(res fyne.Resource) fyne.Resource {
	_ = "STUB: not implemented"
	return *new(fyne.Resource)
}

func colorizeLogError(src []byte, clr color.Color) []byte { _ = "STUB: not implemented"; return nil }
