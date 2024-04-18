package repositories

import "github.com/AntonyChR/orus-media-server/internal/domain/models"

type FileInfoRepository interface {
	Repository
	Save(fileInfo *models.FileInfo) error
	GetAll() ([]models.FileInfo, error)
	GetBy(field string, value interface{}) ([]models.FileInfo, error)
	GetOneBy(field string, value interface{}) (models.FileInfo, error)
	Delete(models.FileInfo) error
	DeleteBy(field string, value interface{}) error
}
