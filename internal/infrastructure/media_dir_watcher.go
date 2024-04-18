package infrastructure

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	domain "github.com/AntonyChR/orus-media-server/internal/domain"
	repositories "github.com/AntonyChR/orus-media-server/internal/domain/repositories"
	fsnotify "github.com/fsnotify/fsnotify"
)

func NewMediaDirWatcher(
	mediaDir string,
	fileExplorerService domain.MediaFileExplorer,
	titleInfoProvider domain.TitleInfoProvider,
	titleInfoRepository repositories.TitleInfoRepository,
	fileInfoRepository repositories.FileInfoRepository,
) *WatchMediafileEvents {
	return &WatchMediafileEvents{
		WatchedDirectoryPath: mediaDir,
		EventChannel:         make(chan MediaChangeEvent),
		FileExplorerService:  fileExplorerService,
		TitleInfoProvider:    titleInfoProvider,
		TitleInfoRepository:  titleInfoRepository,
		FileInfoRepository:   fileInfoRepository,
	}
}

const (
	NEW_FILE = iota
	NEW_DIR
	REMOVE_FILE
	REMOVE_DIR
)

type MediaChangeEvent struct {
	Type     int
	FilePath string
	Error    error
}

type WatchMediafileEvents struct {
	WatchedDirectoryPath string
	EventChannel         chan MediaChangeEvent
	TitleInfoProvider    domain.TitleInfoProvider
	FileExplorerService  domain.MediaFileExplorer
	TitleInfoRepository  repositories.TitleInfoRepository
	FileInfoRepository   repositories.FileInfoRepository
}

func (w *WatchMediafileEvents) WatchDirectoryEvents() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Watchig: ", w.WatchedDirectoryPath)
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					w.EventChannel <- MediaChangeEvent{Error: errors.New("error getting channel value")}
				}
				if event.Has(fsnotify.Chmod) || event.Has(fsnotify.Write) {
					continue
				}

				mediaEvent := MediaChangeEvent{FilePath: event.Name}

				switch {
				case fsnotify.Create == event.Op:
					// If the creation of a folder within a subdirectory is detected, it is ignored
					//
					//	media directory/
					//	|____movie1-1992.mp4
					//	|____movie2-2007.mp4
					//	|____tv show - 1994/  -> subdir: added to event watch list
					//	| |____s2e23.mp4
					//	| |____s1e1.mp4
					//	| |____subdir/        -> inSubdir (invalid): this directory is ignored
					//
					if isDir(event.Name) {
						if !isInSubdir(w.WatchedDirectoryPath, event.Name) {
							log.Println("Watchig: ", event.Name)
							watcher.Add(event.Name)
							mediaEvent.Type = NEW_DIR
						}
					} else {
						mediaEvent.Type = NEW_FILE
					}

				case fsnotify.Remove == event.Op || fsnotify.Rename == event.Op:
					if isDir(event.Name) {
						if !isInSubdir(w.WatchedDirectoryPath, event.Name) {
							log.Println("Remove directory from the watch list: ", event.Name)
							watcher.Remove(event.Name)
							mediaEvent.Type = REMOVE_DIR

						}
					} else {
						mediaEvent.Type = REMOVE_FILE
					}
				}

				w.EventChannel <- mediaEvent

			case err := <-watcher.Errors:
				w.EventChannel <- MediaChangeEvent{Error: err}
			}
		}
	}()

	// Add main path.
	err = watcher.AddWith(w.WatchedDirectoryPath)
	if err != nil {
		log.Fatal(err)
	}

	// add subdiretctories
	fileInfo, _ := w.FileExplorerService.GetInfoAboutMediaFiles(w.WatchedDirectoryPath)

	for _, f := range fileInfo {
		if f.IsDir {
			watcher.AddWith(f.Path)
		}
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

//FIXME:The value of the 'path' is very ambiguous, in several cases it requires
//      the use of 'filepath.Base()' or 'filepath.Dir()', which makes the code
//      even less readable."

//TODO: create service to encapsulate creation/remove data from database and not use repositories instances directly
//REFACTOR: create service to encapsulate creation/remove data from database and not use repositories instances directly

func (w *WatchMediafileEvents) ListenMediaEvents() []string {
	for {
		event := <-w.EventChannel
		if event.Error != nil {
			log.Println("Error: ", event.Error)
			continue
		}
		switch event.Type {

		case NEW_DIR:
			fileName := filepath.Base(event.FilePath)
			titleInfo, err := w.TitleInfoProvider.Search(fileName)
			if err != nil {
				log.Println(err)
				continue
			}

			titleInfo.Folder = event.FilePath

			localInfo, _ := w.TitleInfoRepository.GetOneBy("imdb_id", titleInfo.ImdbID)

			if localInfo.ID != 0 {
				log.Printf("the title information:\"%s\" already exists\n", localInfo.Title)
				continue
			}

			err = w.TitleInfoRepository.Save(&titleInfo)
			if err != nil {
				log.Println(err)
			}
			//TODO: check if the folder contains files, if so, iterate over them and get the information
		case REMOVE_DIR:
			titleInfo, err := w.TitleInfoRepository.GetOneBy("folder", event.FilePath)
			if err != nil {
				log.Println(err)
				continue
			}

			err = w.FileInfoRepository.DeleteBy("title_id", titleInfo.ID)
			if err != nil {
				log.Println(err)
				continue
			}

			err = w.TitleInfoRepository.Delete(titleInfo)
			if err != nil {
				log.Println(err)
				continue
			}
		case NEW_FILE:
			dir := filepath.Dir(event.FilePath)
			if dir == filepath.Base(w.WatchedDirectoryPath) {
				titleInfo, err := w.TitleInfoProvider.Search(filepath.Base(event.FilePath))
				if err != nil {
					log.Println(err)
					continue
				}
				err = w.TitleInfoRepository.Save(&titleInfo)
				if err != nil {
					log.Println(err)
					continue
				}
				fileInfo, _ := w.FileExplorerService.GetFileInfo(event.FilePath)
				fileInfo.TitleId = titleInfo.ID

				err = w.FileInfoRepository.Save(fileInfo)
				if err != nil {
					log.Println(err)
				}
				continue
			}

			fileInfo, err := w.FileExplorerService.GetFileInfo(event.FilePath)
			if err != nil {
				log.Println(err)
				continue
			}
			titleInfo, err := w.TitleInfoRepository.GetOneBy("folder", dir)
			if err != nil || titleInfo.ID == 0 {
				continue
			}

			fileInfo.TitleId = titleInfo.ID
			err = w.FileInfoRepository.Save(fileInfo)
			if err != nil {
				log.Println(err)
			}

		case REMOVE_FILE:
			dir := filepath.Dir(event.FilePath)
			if dir == filepath.Base(w.WatchedDirectoryPath) {
				fileInfo, err := w.FileInfoRepository.GetOneBy("name", filepath.Base(event.FilePath))
				if err != nil {
					log.Println(err)
					continue
				}
				titleInfo, err := w.TitleInfoRepository.GetOneBy("id", fileInfo.TitleId)

				if err != nil {
					log.Println(err)
					continue
				}

				w.FileInfoRepository.Delete(fileInfo)
				w.TitleInfoRepository.Delete(titleInfo)
				continue
			}
			fileInfo, err := w.FileInfoRepository.GetOneBy("name", filepath.Base(event.FilePath))
			if err != nil {
				log.Println(err)
				continue
			}
			w.FileInfoRepository.Delete(fileInfo)
		}
	}

}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		splittedPath := strings.Split(path, ".")
		return len(splittedPath) == 1
	}
	return info.IsDir()
}

// Check if "subdirPath" is inside subdirectory of "mainPath"
//
//	mainPath/
//		|____subdir/     -> false
//		| |____subdir2/  -> true
func isInSubdir(mainPath, subdirPath string) bool {
	splMainPath := strings.Split(mainPath, "/")
	splSubdir := strings.Split(subdirPath, "/")
	return len(splSubdir)-len(splMainPath) == 1

}
