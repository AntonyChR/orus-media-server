package services

import (
	"log"
	"path/filepath"

	domain "github.com/AntonyChR/orus-media-server/internal/domain"
	repositories "github.com/AntonyChR/orus-media-server/internal/domain/repositories"
)

type InfoService struct {
	TitleInfoProvider   domain.TitleInfoProvider
	FileExplorerService domain.MediaFileExplorer
	TitleInfoRepository repositories.TitleInfoRepository
	FileInfoRepository  repositories.FileInfoRepository
}

func (i *InfoService) NewSerie(folderPath string) error {
	dirName := filepath.Base(folderPath)
	titleInfo, err := i.TitleInfoProvider.Search(dirName)
	if err != nil {
		log.Println(err)
		return err
	}

	titleInfo.Folder = folderPath

	localInfo, _ := i.TitleInfoRepository.GetOneBy("imdb_id", titleInfo.ImdbID)

	if localInfo.ID != 0 {
		return nil
	}

	err = i.TitleInfoRepository.Save(&titleInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
