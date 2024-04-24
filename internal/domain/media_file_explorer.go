package domain

import "github.com/AntonyChR/orus-media-server/internal/domain/models"

type MediaFileExplorer interface {
	ScanDir(path string) ([]models.FileInfo, error)
	GetFileInfo(path string) (models.FileInfo, error)
}
