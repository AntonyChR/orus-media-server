package services

import (
	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	repositories "github.com/AntonyChR/orus-media-server/internal/domain/repositories"
)

func NewSubtitleService(subtitleRepository repositories.SubtitleRepository) *SubtitleService {
	return &SubtitleService{
		SubtitleRepository: subtitleRepository,
	}
}

type SubtitleService struct {
	SubtitleRepository repositories.SubtitleRepository
}

func (s *SubtitleService) Save(subtitle *models.Subtitle) error {
	return s.SubtitleRepository.Save(subtitle)
}

func (s *SubtitleService) SaveAll(subtitles *[]models.Subtitle) error {
	return s.SubtitleRepository.SaveAll(subtitles)
}

func (s *SubtitleService) GetAll() ([]models.Subtitle, error) {
	data, err := s.SubtitleRepository.GetAll()
	return data, err
}

func (s *SubtitleService) GetByVideoId(videoId uint) ([]models.Subtitle, error) {
	data, err := s.SubtitleRepository.GetBy("video_id", videoId)
	return data, err
}

func (s *SubtitleService) SetNullVideoFileId(subtitleId uint) error {
	subtitle, err := s.SubtitleRepository.GetOneBy("id", subtitleId)
	if err != nil {
		return err
	}
	subtitle.VideoId = 0
	err = s.SubtitleRepository.Update(&subtitle)
	return err
}

func (s *SubtitleService) SetVideoId(subtitleId, videoId uint) error {
	subtitle, err := s.SubtitleRepository.GetOneBy("id", subtitleId)
	if err != nil {
		return err
	}
	subtitle.VideoId = videoId
	err = s.SubtitleRepository.Update(&subtitle)
	return err
}

func (s *SubtitleService) Reset() error {
	err := s.SubtitleRepository.DropDatabase()
	if err != nil {
		return err
	}
	return s.SubtitleRepository.Migrate()
}
