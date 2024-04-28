package services

import (
	"log"
	"os"
	"path/filepath"

	domain "github.com/AntonyChR/orus-media-server/internal/domain"
	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
)

func NewMediaInfoSyncService(
	mediaDir string,
	subtitlesDir string,
	fileExplorerService domain.MediaFileExplorer,
	titleInfoProvider domain.TitleInfoProvider,
	videoService *VideoService,
	titleInfoService *TitleInfoService,
	subtitleService *SubtitleService,
) *MediaInfoSyncService {
	return &MediaInfoSyncService{
		MediaDir:            mediaDir,
		SubtitlesDir:        subtitlesDir,
		FileExplorerService: fileExplorerService,
		TitleInfoProvider:   titleInfoProvider,
		VideoService:        videoService,
		TitleInfoService:    titleInfoService,
		SubtitleService:     subtitleService,
	}
}

type MediaInfoSyncService struct {
	MediaDir     string
	SubtitlesDir string

	FileExplorerService domain.MediaFileExplorer
	TitleInfoProvider   domain.TitleInfoProvider

	VideoService     *VideoService
	TitleInfoService *TitleInfoService
	SubtitleService  *SubtitleService
}

func (m *MediaInfoSyncService) GetTitleInfoAboutAllMediaFiles() error {

	log.Println("Scanning media files")

	mediaFiles, err := m.FileExplorerService.ScanDir(m.MediaDir)

	if err != nil {
		return err
	}

	for _, file := range mediaFiles {
		if file.IsDir {
			titleInfo, searchErr := m.TitleInfoProvider.Search(file.Name)
			titleInfo.Folder = file.Path
			seriesFileInfo, scanErr := m.FileExplorerService.ScanDir(file.Path)

			if scanErr != nil {
				log.Printf("Error getting info about: \"%s\", %s\n", file.Name, scanErr.Error())
			}

			if searchErr == nil {
				if err := m.TitleInfoService.Save(&titleInfo); err != nil {
					log.Println(err)
				}
			}

			if scanErr == nil {
				for _, fileInfo := range seriesFileInfo {
					video := fileInfo.Video
					if searchErr == nil {
						video.TitleId = titleInfo.ID
					}

					m.VideoService.Save(&video)
				}
			}

		} else {
			video := file.Video
			titleInfo, err := m.TitleInfoProvider.Search(video.Name)
			titleInfo.Folder = filepath.Dir(video.Path)
			if err != nil {
				log.Println(err)
			} else {
				if err = m.TitleInfoService.Save(&titleInfo); err != nil { // Create record and asign ID value
					log.Println(err)
				} else {
					video.TitleId = titleInfo.ID
				}
			}
			if err := m.VideoService.Save(&video); err != nil {
				return err
			}
		}
	}

	return nil

}

func (m *MediaInfoSyncService) ScanSubtitles() error {

	log.Println("Scanning subtitles")

	var subtitles []models.Subtitle
	files, err := os.ReadDir(m.SubtitlesDir)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		subtitles = append(subtitles, models.Subtitle{
			Path: filepath.Join(m.SubtitlesDir, file.Name()),
			Name: file.Name(),
			Lang: m.SubtitleService.GetLang(file.Name()),
		})
	}

	if len(subtitles) == 0 {
		return nil
	}

	return m.SubtitleService.SaveAll(&subtitles)

}
