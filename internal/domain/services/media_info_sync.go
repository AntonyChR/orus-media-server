package services

import (
	"log"
	"path/filepath"

	domain "github.com/AntonyChR/orus-media-server/internal/domain"
)

func NewMediaInfoSyncService(
	path string,
	fileExplorerService domain.MediaFileExplorer,
	titleInfoProvider domain.TitleInfoProvider,
	fileInfoService *FileInfoService,
	titleInfoService *TitleInfoService,
) *MediaInfoSyncService {
	return &MediaInfoSyncService{
		Path:                path,
		FileExplorerService: fileExplorerService,
		TitleInfoProvider:   titleInfoProvider,
		FileInfoService:     fileInfoService,
		TitleInfoService:    titleInfoService,
	}
}

type MediaInfoSyncService struct {
	Path string

	FileExplorerService domain.MediaFileExplorer
	TitleInfoProvider   domain.TitleInfoProvider

	FileInfoService  *FileInfoService
	TitleInfoService *TitleInfoService
}

func (m *MediaInfoSyncService) GetTitleInfoAboutAllMediaFiles() error {
	mediaFileInf, err := m.FileExplorerService.ScanDir(m.Path)

	if err != nil {
		return err
	}

	for _, f := range mediaFileInf {
		if f.IsDir {
			titleInfo, titleInfofErr := m.TitleInfoProvider.Search(f.Name)
			titleInfo.Folder = f.Path
			seriesFileInfo, fileInfoErr := m.FileExplorerService.ScanDir(f.Path)

			if fileInfoErr != nil {
				log.Printf("Error getting info about: \"%s\", %s\n", f.Name, fileInfoErr.Error())
			}

			if titleInfofErr == nil {
				err := m.TitleInfoService.Save(&titleInfo)
				if err != nil {
					log.Println(err)
				}
			}

			if fileInfoErr == nil {
				for _, fileInfo := range seriesFileInfo {
					if titleInfofErr == nil {
						fileInfo.TitleId = titleInfo.ID
					}

					m.FileInfoService.Save(&fileInfo)
				}
			}

		} else {
			titleInfo, err := m.TitleInfoProvider.Search(f.Name)
			titleInfo.Folder = filepath.Dir(f.Path)
			if err != nil {
				log.Println(err)
			} else {
				if err = m.TitleInfoService.Save(&titleInfo); err != nil { // Create record and asign ID value
					log.Println(err)
				} else {
					f.TitleId = titleInfo.ID
				}
			}
			if err := m.FileInfoService.Save(&f); err != nil {
				return err
			}
		}
	}

	return nil

}
