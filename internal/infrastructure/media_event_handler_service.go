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
	fileInfoService *services.FileInfoService,
	fileExplorerService domain.MediaFileExplorer,
	titleInfoProvider domain.TitleInfoProvider,
) *MediaEventHandlerService {
	return &MediaEventHandlerService{
		WatchedMediaDir:     watchedMediaDir,
		TitleInfoService:    titleInfoService,
		FileInfoService:     fileInfoService,
		FileExplorerService: fileExplorerService,
		TitleInfoProvider:   titleInfoProvider,
	}
}

type MediaEventHandlerService struct {
	WatchedMediaDir     string
	TitleInfoService    *services.TitleInfoService
	FileInfoService     *services.FileInfoService
	FileExplorerService domain.MediaFileExplorer
	TitleInfoProvider   domain.TitleInfoProvider
}

func (s *MediaEventHandlerService) HandleNewDir(event MediaChangeEvent) error {
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

func (s *MediaEventHandlerService) HandleRemoveDir(event MediaChangeEvent) error {
	titleInfo, err := s.TitleInfoService.GetByFolder(event.FilePath)

	if err != nil {
		return err
	}

	// delete all the files associated with the title information
	if err = s.FileInfoService.DeleteByTitleId(titleInfo.ID); err != nil {
		return err
	}

	err = s.TitleInfoService.DeleteById(titleInfo.ID)

	return err
}

func (s *MediaEventHandlerService) HandleNewFile(event MediaChangeEvent) error {
	dir := filepath.Dir(event.FilePath)

	// if the file is in the root directory of the watched media directory that means it is a movie
	// otherwise it is the chapter of a series
	if dir == filepath.Base(s.WatchedMediaDir) {
		fileInfo, err := s.FileExplorerService.GetFileInfo(event.FilePath)

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
		fileInfo.TitleId = titleInfo.ID

		err = s.FileInfoService.Save(&fileInfo)
		return err
	}

	fileInfo, err := s.FileExplorerService.GetFileInfo(event.FilePath)
	if err != nil {
		return err
	}
	titleInfo, err := s.TitleInfoService.GetByFolder(dir)
	if err != nil || titleInfo.ID == 0 {
		return err
	}

	fileInfo.TitleId = titleInfo.ID
	err = s.FileInfoService.Save(&fileInfo)
	return err

}

func (s *MediaEventHandlerService) HandleRemoveFile(event MediaChangeEvent) error {
	dir := filepath.Dir(event.FilePath)

	// if the file is in the root directory of the watched media directory that means it is a movie
	if dir == filepath.Base(s.WatchedMediaDir) {

		fileInfo, err := s.FileInfoService.GetByName(filepath.Base(event.FilePath))

		if err != nil {
			return err
		}

		titleInfo, err := s.TitleInfoService.GetById(fileInfo.TitleId)

		if err != nil {
			return err
		}

		s.FileInfoService.DeleteById(fileInfo.ID)
		err = s.TitleInfoService.DeleteById(titleInfo.ID)
		return err
	}

	fileInfo, err := s.FileInfoService.GetByName(filepath.Base(event.FilePath))

	if err != nil {
		return err
	}

	err = s.FileInfoService.DeleteById(fileInfo.ID)

	return err
}
