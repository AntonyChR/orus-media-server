package repositories

import "github.com/AntonyChR/orus-media-server/internal/domain/models"

type VideoRepository interface {
	Repository
	Save(video *models.Video) error
	GetAll() ([]models.Video, error)
	GetBy(field string, value interface{}) ([]models.Video, error)
	GetOneBy(field string, value interface{}) (models.Video, error)
	Delete(models.Video) error
	DeleteBy(field string, value interface{}) error
}
