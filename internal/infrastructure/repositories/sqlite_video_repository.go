package repositoryImplementations

import (
	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	gorm "gorm.io/gorm"
)

func NewSqliteVideoRepo(db *gorm.DB) *SqliteVideoRepository {
	return &SqliteVideoRepository{db}
}

type SqliteVideoRepository struct {
	Db *gorm.DB
}

func (s *SqliteVideoRepository) Delete(video models.Video) error {
	// By default Gorm does not delete the record, it only adds the "deleted_at" field
	// which will make the record not available in queries. To permanently delete the record,
	// use the "Unscoped()" method.
	result := s.Db.Unscoped().Delete(&video)
	return result.Error
}

func (s *SqliteVideoRepository) DeleteBy(field string, value interface{}) error {
	query := field + " = ?"
	result := s.Db.Unscoped().Where(query, value).Delete(&models.Video{})
	return result.Error
}

func (s *SqliteVideoRepository) Save(video *models.Video) error {
	res := s.Db.Create(video)
	return res.Error
}

func (s *SqliteVideoRepository) GetOneBy(field string, value interface{}) (models.Video, error) {
	var video models.Video
	query := field + " = ? "
	result := s.Db.Where(query, value).First(&video)
	if result.Error != nil {
		return video, result.Error
	}
	return video, nil

}

func (s *SqliteVideoRepository) GetBy(field string, value interface{}) ([]models.Video, error) {
	var data []models.Video
	query := field + " = ? "

	result := s.Db.Where(query, value).Find(&data)
	if result.Error != nil {
		return []models.Video{}, result.Error
	}
	return data, nil
}

func (s *SqliteVideoRepository) GetAll() ([]models.Video, error) {
	var data []models.Video
	res := s.Db.Find(&data)
	if res.Error != nil {
		return []models.Video{}, res.Error
	}
	return data, nil
}

func (s *SqliteVideoRepository) DropDatabase() error {
	err := s.Db.Migrator().DropTable(&models.Video{})
	return err
}

func (s *SqliteVideoRepository) Migrate() error {
	return s.Db.Migrator().CreateTable(&models.Video{})
}
