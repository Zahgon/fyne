//go:build !ios

package dialog

import (
	"fyne.io/fyne/v2"
)

const folderVideos = "Movies"

func getFavoriteLocation(homeURI fyne.URI, name string) (fyne.URI, error) {
	_ = "STUB: not implemented"
	return *new(fyne.URI), nil
}
