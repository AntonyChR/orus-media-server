package repositories

import "github.com/AntonyChR/orus-media-server/internal/domain/models"

type SubtitleRepository interface {
	Repository
	Save(subtitle *models.Subtitle) error
	GetAll() ([]models.Subtitle, error)
	GetBy(field string, value interface{}) ([]models.Subtitle, error)
	GetOneBy(field string, value interface{}) (models.Subtitle, error)
	Update(subtitle *models.Subtitle) error
	UpdateSingleColumn(subtitle *models.Subtitle, field string, value interface{}) error
	SaveAll(subtitles *[]models.Subtitle) error
	//Delete(models.SubtitleFile) error
	//DeleteBy(field string, value interface{}) error
}
