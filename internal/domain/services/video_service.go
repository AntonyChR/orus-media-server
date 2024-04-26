package services

import (
	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	repositories "github.com/AntonyChR/orus-media-server/internal/domain/repositories"
)

func NewVideoService(videoRepository repositories.VideoRepository) *VideoService {
	return &VideoService{
		VideoRepository: videoRepository,
	}
}

type VideoService struct {
	VideoRepository repositories.VideoRepository
}

func (f *VideoService) GetAll() ([]models.Video, error) {
	data, err := f.VideoRepository.GetAll()
	return data, err
}

func (f *VideoService) GetByTitleId(titleId uint) ([]models.Video, error) {
	data, err := f.VideoRepository.GetBy("title_id", titleId)
	return data, err
}
func (f *VideoService) GetByVideoId(videoId uint) (models.Video, error) {
	data, err := f.VideoRepository.GetOneBy("id", videoId)
	return data, err
}

func (f *VideoService) GetByName(name string) (models.Video, error) {
	fileInfo, err := f.VideoRepository.GetOneBy("name", name)
	return fileInfo, err
}

func (f *VideoService) DeleteById(id uint) error {
	return f.VideoRepository.DeleteBy("id", id)
}

func (f *VideoService) DeleteByTitleId(titleId uint) error {
	return f.VideoRepository.DeleteBy("title_id", titleId)
}

func (f *VideoService) Save(fileInfo *models.Video) error {
	return f.VideoRepository.Save(fileInfo)
}

func (f *VideoService) GetVideosWithNoTitleId() ([]models.Video, error) {
	data, err := f.VideoRepository.GetBy("title_id", 0)
	return data, err
}

func (f *VideoService) Reset() error {
	err := f.VideoRepository.DropDatabase()
	if err != nil {
		return err
	}
	return f.VideoRepository.Migrate()
}
