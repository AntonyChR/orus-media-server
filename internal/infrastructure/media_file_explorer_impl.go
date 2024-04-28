package infrastructure

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
)

func NewMediaFileExplorer() *MediaFileExplorerImpl {
	return &MediaFileExplorerImpl{}
}

// Implements diferent methods to get file information
type MediaFileExplorerImpl struct{}

func (f *MediaFileExplorerImpl) ScanDir(path string) ([]models.FileInfo, error) {
	files, err := os.ReadDir(path)

	info := []models.FileInfo{}

	if err != nil {
		return info, err
	}

	for _, f := range files {
		var tmp models.FileInfo

		if f.IsDir() {
			tmp = models.FileInfo{
				Name:  f.Name(),
				Path:  filepath.Join(path, f.Name()),
				IsDir: true,
			}
			info = append(info, tmp)
			continue
		}

		season, episode := getSeasonAndEpisode(f.Name())

		tmp = models.FileInfo{
			Video: models.Video{
				Name:    f.Name(),
				Path:    filepath.Join(path, f.Name()),
				Season:  season,
				Episode: episode,
				Ext:     filepath.Ext(f.Name())[1:],
			},
			IsDir: false,
		}

		info = append(info, tmp)
	}
	return info, nil
}

func (f *MediaFileExplorerImpl) GetVideoInfo(path string) (models.Video, error) {
	fileData, err := os.Stat(path)
	if err != nil {
		return models.Video{}, err
	}
	season, episode := getSeasonAndEpisode(fileData.Name())

	return models.Video{
		Name:    fileData.Name(),
		Path:    path,
		Season:  season,
		Episode: episode,
		Ext:     filepath.Ext(fileData.Name())[1:],
	}, nil
}

func getSeasonAndEpisode(fileName string) (uint, uint) {
	regEx := regexp.MustCompile(`s(\d+)e(\d+)`)
	seStr := regEx.FindString(fileName)
	if seStr == "" {
		return 0, 0
	}
	splited := strings.Split(seStr, "e")
	seasonStr := splited[0][1:]
	episodeStr := splited[1]
	season, _ := strconv.ParseUint(seasonStr, 10, 32)
	episode, _ := strconv.ParseUint(episodeStr, 10, 32)
	return uint(season), uint(episode)
}

func CheckDirectories(paths ...string) error {
	for _, path := range paths {
		err := MkDirIfNotExist(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func MkDirIfNotExist(path string) error {
	if _, err := os.Stat(path); err != nil {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
