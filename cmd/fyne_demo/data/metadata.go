//go:generate fyne bundle -o metadata_bundled.go -package data -name resourceAuthors ../../../AUTHORS

package data

var Authors = resourceAuthors
