package services

import (
	"github.com/AntonyChR/orus-media-server/internal/domain/models"
	"github.com/AntonyChR/orus-media-server/internal/domain/repositories"
)

func NewFileInfoService(fileInfoRepository repositories.FileInfoRepository) *FileInfoService {
	return &FileInfoService{
		FileInfoRepository: fileInfoRepository,
	}
}

type FileInfoService struct {
	FileInfoRepository repositories.FileInfoRepository
}

func (f *FileInfoService) GetAll() ([]models.FileInfo, error) {
	data, err := f.FileInfoRepository.GetAll()
	return data, err
}

func (f *FileInfoService) GetByTitleId(titleId uint) ([]models.FileInfo, error) {
	data, err := f.FileInfoRepository.GetBy("title_id", titleId)
	return data, err
}
func (f *FileInfoService) GetByVideoId(videoId uint) (models.FileInfo, error) {
	data, err := f.FileInfoRepository.GetOneBy("id", videoId)
	return data, err
}

func (f *FileInfoService) GetByName(name string) (models.FileInfo, error) {
	fileInfo, err := f.FileInfoRepository.GetOneBy("name", name)
	return fileInfo, err
}

func (f *FileInfoService) DeleteByTitleId(titleId string) error {
	return f.FileInfoRepository.DeleteBy("title_id", titleId)
}

func (f *FileInfoService) Save(fileInfo *models.FileInfo) error {
	return f.FileInfoRepository.Save(fileInfo)
}

func (f *FileInfoService) Reset() error {
	err := f.FileInfoRepository.DropDatabase()
	if err != nil {
		return err
	}
	return f.FileInfoRepository.Migrate()
}
