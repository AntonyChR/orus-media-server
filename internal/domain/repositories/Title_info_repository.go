package repositories

import "github.com/AntonyChR/orus-media-server/internal/domain/models"

type TitleInfoRepository interface {
	Repository
	Save(info *models.TitleInfo) error
	GetAll() ([]models.TitleInfo, error)
	GetBy(field string, value interface{}) ([]models.TitleInfo, error)
	GetOneBy(field string, value interface{}) (models.TitleInfo, error)
	Delete(models.TitleInfo) error
	DeleteBy(field string, value interface{}) error
}
