package infrastructure

import (
	"fmt"
	"log"
	"path/filepath"

	domain "github.com/AntonyChR/orus-media-server/internal/domain"
	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
)

func NewMediaEventHandlerService(
	watchedMediaDir string,
	titleInfoService *services.TitleInfoService,
	videoService *services.VideoService,
	fileExplorerService domain.MediaFileExplorer,
	titleInfoProvider domain.TitleInfoProvider,
) *MediaEventHandlerService {
	return &MediaEventHandlerService{
		WatchedMediaDir:     watchedMediaDir,
		TitleInfoService:    titleInfoService,
		videoService:        videoService,
		FileExplorerService: fileExplorerService,
		TitleInfoProvider:   titleInfoProvider,
	}
}

type MediaEventHandlerService struct {
	WatchedMediaDir     string
	TitleInfoService    *services.TitleInfoService
	videoService        *services.VideoService
	FileExplorerService domain.MediaFileExplorer
	TitleInfoProvider   domain.TitleInfoProvider
}

func (s *MediaEventHandlerService) HandleNewDir(event MediaDirEvent) error {
	fileName := filepath.Base(event.FilePath)
	titleInfo, err := s.TitleInfoProvider.Search(fileName)
	if err != nil {
		return err
	}

	titleInfo.Folder = event.FilePath

	localInfo, _ := s.TitleInfoService.GetByImdbId(titleInfo.ImdbID)

	if localInfo.ID != 0 {
		return fmt.Errorf("the title information:\"%s\" already exists", localInfo.Title)
	}

	err = s.TitleInfoService.Save(&titleInfo)

	return err
}

func (s *MediaEventHandlerService) HandleRemoveDir(event MediaDirEvent) error {
	titleInfo, err := s.TitleInfoService.GetByFolder(event.FilePath)

	if err != nil {
		return err
	}

	// delete all the files associated with the title information
	if err = s.videoService.DeleteByTitleId(titleInfo.ID); err != nil {
		return err
	}

	err = s.TitleInfoService.DeleteById(titleInfo.ID)

	return err
}

func (s *MediaEventHandlerService) HandleNewFile(event MediaDirEvent) error {
	dir := filepath.Dir(event.FilePath)

	// if the file is in the root directory of the watched media directory that means it is a movie
	// otherwise it is the chapter of a series
	if dir == filepath.Base(s.WatchedMediaDir) {
		videoInfo, err := s.FileExplorerService.GetVideoInfo(event.FilePath)

		if err != nil {
			return err
		}

		titleInfo, err := s.TitleInfoProvider.Search(filepath.Base(event.FilePath))

		if err == nil {
			if err := s.TitleInfoService.Save(&titleInfo); err != nil {
				return err
			}
		} else {
			log.Println(err)
		}

		// associate the file with the title information
		videoInfo.TitleId = titleInfo.ID

		err = s.videoService.Save(&videoInfo)
		return err

	} else {

		video, err := s.FileExplorerService.GetVideoInfo(event.FilePath)
		if err != nil {
			return err
		}
		titleInfo, err := s.TitleInfoService.GetByFolder(dir)
		if err != nil || titleInfo.ID == 0 {
			return err
		}

		video.TitleId = titleInfo.ID
		err = s.videoService.Save(&video)
		return err
	}

}

func (s *MediaEventHandlerService) HandleRemoveFile(event MediaDirEvent) error {
	dir := filepath.Dir(event.FilePath)

	// if the file is in the root directory of the watched media directory that means it is a movie
	if dir == filepath.Base(s.WatchedMediaDir) {

		video, err := s.videoService.GetByName(filepath.Base(event.FilePath))

		if err != nil {
			return err
		}

		if video.TitleId == 0 {
			return s.videoService.DeleteById(video.ID)
		}

		titleInfo, err := s.TitleInfoService.GetById(video.TitleId)

		if err != nil {
			return err
		}

		s.videoService.DeleteById(video.ID)
		err = s.TitleInfoService.DeleteById(titleInfo.ID)
		return err
	} else {

		video, err := s.videoService.GetByName(filepath.Base(event.FilePath))

		if err != nil {
			return err
		}

		err = s.videoService.DeleteById(video.ID)

		return err
	}
}
