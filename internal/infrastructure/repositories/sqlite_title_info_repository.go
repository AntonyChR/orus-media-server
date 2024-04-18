package repositoryImplementations

import (
	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	gorm "gorm.io/gorm"
)

func NewSqliteTitleInfoRepo(db *gorm.DB) *SqliteTitleInfoRepository {
	return &SqliteTitleInfoRepository{Db: db}
}

type SqliteTitleInfoRepository struct {
	Db *gorm.DB
}

func (s *SqliteTitleInfoRepository) Delete(titleInfo models.TitleInfo) error {
	result := s.Db.Unscoped().Delete(&titleInfo)
	return result.Error
}

func (s *SqliteTitleInfoRepository) DeleteBy(field string, value interface{}) error {
	query := field + " = ?"
	result := s.Db.Unscoped().Where(query, value).Delete(&models.TitleInfo{})
	return result.Error
}

func (s *SqliteTitleInfoRepository) GetOneBy(field string, value interface{}) (models.TitleInfo, error) {
	var titleInfo models.TitleInfo
	query := field + " = ? "
	result := s.Db.Where(query, value).First(&titleInfo)
	if result.Error != nil {
		return titleInfo, result.Error
	}
	return titleInfo, nil

}

func (s *SqliteTitleInfoRepository) GetBy(field string, value interface{}) ([]models.TitleInfo, error) {
	var data []models.TitleInfo
	query := field + " = ? "

	result := s.Db.Where(query, value).Find(&data)
	if result.Error != nil {
		return []models.TitleInfo{}, result.Error
	}
	return data, nil
}

func (s *SqliteTitleInfoRepository) GetAll() ([]models.TitleInfo, error) {
	var data []models.TitleInfo
	result := s.Db.Find(&data)
	if result.Error != nil {
		return []models.TitleInfo{}, nil
	}
	return data, nil
}

func (s *SqliteTitleInfoRepository) Save(mediaInfo *models.TitleInfo) error {
	result := s.Db.Create(mediaInfo)
	return result.Error
}

func (s *SqliteTitleInfoRepository) DropDatabase() error {
	err := s.Db.Migrator().DropTable(&models.TitleInfo{})
	return err
}
func (s *SqliteTitleInfoRepository) Migrate() error {
	err := s.Db.Migrator().CreateTable(&models.TitleInfo{})
	return err
}
