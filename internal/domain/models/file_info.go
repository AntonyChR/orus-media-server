package models

type FileInfo struct {
	Video Video
	IsDir bool
	Path  string
	Name  string
}
