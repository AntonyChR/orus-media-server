package repositoryImplementations

import (
	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	gorm "gorm.io/gorm"
)

func NewSqliteFileInfoRepo(db *gorm.DB) *SqliteFileInfoRepository {
	return &SqliteFileInfoRepository{db}
}

type SqliteFileInfoRepository struct {
	Db *gorm.DB
}

func (s *SqliteFileInfoRepository) Delete(fileInfo models.FileInfo) error {
	// By default Gorm does not delete the record, it only adds the "deleted_at" field
	// which will make the record not available in queries. To permanently delete the record,
	// use the "Unscoped()" method.
	result := s.Db.Unscoped().Delete(&fileInfo)
	return result.Error
}

func (s *SqliteFileInfoRepository) DeleteBy(field string, value interface{}) error {
	query := field + " = ?"
	result := s.Db.Unscoped().Where(query, value).Delete(&models.FileInfo{})
	return result.Error
}

func (s *SqliteFileInfoRepository) Save(fileInfo *models.FileInfo) error {
	res := s.Db.Create(fileInfo)
	return res.Error
}

func (s *SqliteFileInfoRepository) GetOneBy(field string, value interface{}) (models.FileInfo, error) {
	var fileInfo models.FileInfo
	query := field + " = ? "
	result := s.Db.Where(query, value).First(&fileInfo)
	if result.Error != nil {
		return fileInfo, result.Error
	}
	return fileInfo, nil

}

func (s *SqliteFileInfoRepository) GetBy(field string, value interface{}) ([]models.FileInfo, error) {
	var data []models.FileInfo
	query := field + " = ? "

	result := s.Db.Where(query, value).Find(&data)
	if result.Error != nil {
		return []models.FileInfo{}, result.Error
	}
	return data, nil
}

func (s *SqliteFileInfoRepository) GetAll() ([]models.FileInfo, error) {
	var data []models.FileInfo
	res := s.Db.Find(&data)
	if res.Error != nil {
		return []models.FileInfo{}, res.Error
	}
	return data, nil
}

func (s *SqliteFileInfoRepository) DropDatabase() error {
	err := s.Db.Migrator().DropTable(&models.FileInfo{})
	return err
}

func (s *SqliteFileInfoRepository) Migrate() error {
	return s.Db.Migrator().CreateTable(&models.FileInfo{})
}
