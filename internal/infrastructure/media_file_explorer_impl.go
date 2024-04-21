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

// Gets information about all files within the "path" and returns an array of models.FileInfo
// If the file is inside the subdirectory, the name must be in a format like: "s1e1.mp4".
func (f *MediaFileExplorerImpl) ScanDir(path string) ([]*models.FileInfo, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return []*models.FileInfo{}, err
	}

	info := []*models.FileInfo{}
	for _, f := range files {
		season, episode := getSeasonAndEpisode(f.Name())
		tmp := models.FileInfo{
			Name:    f.Name(),
			Path:    filepath.Join(path, f.Name()),
			IsDir:   f.IsDir(),
			Season:  season,
			Episode: episode,
		}
		info = append(info, &tmp)
	}
	return info, nil
}

func (f *MediaFileExplorerImpl) GetFileInfo(path string) (*models.FileInfo, error) {
	fileData, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	season, episode := getSeasonAndEpisode(fileData.Name())

	return &models.FileInfo{
		Name:    fileData.Name(),
		Path:    path,
		IsDir:   fileData.IsDir(),
		Season:  season,
		Episode: episode,
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

func CreateDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}
