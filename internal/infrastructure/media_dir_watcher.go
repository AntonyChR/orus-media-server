package infrastructure

import (
	"errors"
	"log"
	"os"
	"strings"

	domain "github.com/AntonyChR/orus-media-server/internal/domain"
	fsnotify "github.com/fsnotify/fsnotify"
)

func NewMediaDirWatcher(
	mediaDir string,
	fileExplorerService domain.MediaFileExplorer,
	titleInfoProvider domain.TitleInfoProvider,
	eventHandler EventHandlerService,

) *WatchMediafileEvents {
	return &WatchMediafileEvents{
		WatchedMediaDir:     mediaDir,
		EventChannel:        make(chan MediaChangeEvent),
		FileExplorerService: fileExplorerService,
		TitleInfoProvider:   titleInfoProvider,
		EventHandlerService: eventHandler,
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

type EventHandlerService interface {
	HandleNewDir(event MediaChangeEvent) error
	HandleRemoveDir(event MediaChangeEvent) error
	HandleNewFile(event MediaChangeEvent) error
	HandleRemoveFile(event MediaChangeEvent) error
}

type WatchMediafileEvents struct {
	WatchedMediaDir     string
	EventChannel        chan MediaChangeEvent
	EventHandlerService EventHandlerService
	FileExplorerService domain.MediaFileExplorer

	TitleInfoProvider domain.TitleInfoProvider
}

func (w *WatchMediafileEvents) WatchDirectoryEvents() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Watchig: ", w.WatchedMediaDir)
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
					//	|____movie1 (1992).mp4
					//	|____movie2 (2007).mp4
					//	|____tv show (1994)/  -> subdir: added to event watch list
					//	| |____s2e23.mp4
					//	| |____s1e1.mp4
					//	| |____subdir/        -> inSubdir (invalid): this directory is ignored
					//
					if isDir(event.Name) {
						if !isInSubdir(w.WatchedMediaDir, event.Name) {
							log.Println("Watchig: ", event.Name)
							watcher.Add(event.Name)
							mediaEvent.Type = NEW_DIR
						}
					} else {
						mediaEvent.Type = NEW_FILE
					}

				case fsnotify.Remove == event.Op || fsnotify.Rename == event.Op:
					if isDir(event.Name) {
						if !isInSubdir(w.WatchedMediaDir, event.Name) {
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
	err = watcher.AddWith(w.WatchedMediaDir)
	if err != nil {
		log.Fatal(err)
	}

	// add subdiretctories
	files, _ := w.FileExplorerService.ScanDir(w.WatchedMediaDir)

	for _, f := range files {
		if f.IsDir {
			watcher.AddWith(f.Path)
		}
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func (w *WatchMediafileEvents) ListenMediaEvents() []string {
	for {
		event := <-w.EventChannel

		if event.Error != nil {
			log.Println(event.Error)
			continue
		}

		switch event.Type {

		case NEW_DIR:

			if err := w.EventHandlerService.HandleNewDir(event); err != nil {
				log.Println(err)
			}

		case REMOVE_DIR:

			if err := w.EventHandlerService.HandleRemoveDir(event); err != nil {
				log.Println(err)
			}

		case NEW_FILE:

			if err := w.EventHandlerService.HandleNewFile(event); err != nil {
				log.Println(err)
			}

		case REMOVE_FILE:

			if err := w.EventHandlerService.HandleRemoveFile(event); err != nil {
				log.Println(err)
			}

		}
	}

}

// StartWatching starts watching for media file events.
// It spawns two goroutines: one for watching directory events and another for listening to media events.
func (w *WatchMediafileEvents) StartWatching() {
	go w.WatchDirectoryEvents()
	go w.ListenMediaEvents()
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
