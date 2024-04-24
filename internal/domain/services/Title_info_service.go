package services

import (
	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	repositories "github.com/AntonyChR/orus-media-server/internal/domain/repositories"
)

func NewTitleInfoService(titleInfoRepository repositories.TitleInfoRepository) *TitleInfoService {
	return &TitleInfoService{TitleInfoRepository: titleInfoRepository}
}

type TitleInfoService struct {
	TitleInfoRepository repositories.TitleInfoRepository
}

func (t *TitleInfoService) GetByFolder(folder string) (models.TitleInfo, error) {
	titleInfo, err := t.TitleInfoRepository.GetOneBy("folder", folder)
	return titleInfo, err
}

func (t *TitleInfoService) GetAll() ([]models.TitleInfo, error) {
	titles, err := t.TitleInfoRepository.GetAll()
	return titles, err
}
func (t *TitleInfoService) GetSeries() ([]models.TitleInfo, error) {
	titleInfo, err := t.TitleInfoRepository.GetBy("type", "series")
	return titleInfo, err
}

func (t *TitleInfoService) GetMovies() ([]models.TitleInfo, error) {
	titleInfo, err := t.TitleInfoRepository.GetBy("type", "movie")
	return titleInfo, err
}

func (t *TitleInfoService) GetById(id uint) (models.TitleInfo, error) {
	titleInfo, err := t.TitleInfoRepository.GetOneBy("id", id)
	return titleInfo, err
}

func (t *TitleInfoService) GetByImdbId(imdbId string) (models.TitleInfo, error) {
	titleInfo, err := t.TitleInfoRepository.GetOneBy("imdb_id", imdbId)
	return titleInfo, err
}

func (t *TitleInfoService) Save(titleInfo *models.TitleInfo) error {
	return t.TitleInfoRepository.Save(titleInfo)
}

func (t *TitleInfoService) DeleteById(id uint) error {
	return t.TitleInfoRepository.DeleteBy("id", id)
}

func (t *TitleInfoService) Reset() error {
	err := t.TitleInfoRepository.DropDatabase()
	if err != nil {
		return err
	}
	return t.TitleInfoRepository.Migrate()
}
