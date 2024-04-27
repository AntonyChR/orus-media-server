package repositoryImplementations

import (
	"errors"

	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	gorm "gorm.io/gorm"
)

func NewSqliteSubtitleRepo(db *gorm.DB) *SqliteSubtitleRepository {
	return &SqliteSubtitleRepository{Db: db}
}

type SqliteSubtitleRepository struct {
	Db *gorm.DB
}

func (r *SqliteSubtitleRepository) Save(subtitle *models.Subtitle) error {
	return r.Db.Create(subtitle).Error
}

func (r *SqliteSubtitleRepository) SaveAll(subtitles *[]models.Subtitle) error {
	return r.Db.Create(subtitles).Error
}

func (r *SqliteSubtitleRepository) GetAll() ([]models.Subtitle, error) {
	var subtitles []models.Subtitle
	err := r.Db.Find(&subtitles).Error
	return subtitles, err
}

func (r *SqliteSubtitleRepository) GetBy(field string, value interface{}) ([]models.Subtitle, error) {
	var subtitles []models.Subtitle
	err := r.Db.Where(field+" = ?", value).Find(&subtitles).Error
	return subtitles, err
}

func (r *SqliteSubtitleRepository) GetOneBy(field string, value interface{}) (models.Subtitle, error) {
	var subtitle models.Subtitle
	err := r.Db.Where(field+" = ?", value).First(&subtitle).Error
	return subtitle, err
}

func (r *SqliteSubtitleRepository) Update(subtitle *models.Subtitle) error {
	if subtitle.ID == 0 {
		return errors.New("ID is required")
	}
	return r.Db.Save(subtitle).Error
}

func (r *SqliteSubtitleRepository) UpdateSingleColumn(subtitle *models.Subtitle, field string, value interface{}) error {
	return r.Db.Model(subtitle).Update(field, value).Error
}

func (r *SqliteSubtitleRepository) DropDatabase() error {
	return r.Db.Migrator().DropTable(&models.Subtitle{})
}

func (r *SqliteSubtitleRepository) Migrate() error {
	return r.Db.AutoMigrate(&models.Subtitle{})
}
