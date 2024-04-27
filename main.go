package main

import (
	"embed"
	"log"
	"net/http"

	config "github.com/AntonyChR/orus-media-server/config"
	controllers "github.com/AntonyChR/orus-media-server/internal/delivery/controllers"
	middlewares "github.com/AntonyChR/orus-media-server/internal/delivery/middlewares"
	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
	infrastructure "github.com/AntonyChR/orus-media-server/internal/infrastructure"
	repositoryImplementations "github.com/AntonyChR/orus-media-server/internal/infrastructure/repositories"
	cors "github.com/gin-contrib/cors"

	static "github.com/gin-contrib/static"
	gin "github.com/gin-gonic/gin"
	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

//directive that loads files into the binary at compile time

//go:embed gui/dist/*
var staticContent embed.FS

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	if err := infrastructure.CheckDirectories(config.MEDIA_DIR, config.SUBTITLE_DIR); err != nil {
		log.Fatal(err)
	}

	sqliteDb, err := gorm.Open(sqlite.Open(config.DB_PATH), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Instance repositories

	sqliteVideoRepo := repositoryImplementations.NewSqliteVideoRepo(sqliteDb)
	sqliteSubtitleRepo := repositoryImplementations.NewSqliteSubtitleRepo(sqliteDb)
	sqliteTitleInfoRepo := repositoryImplementations.NewSqliteTitleInfoRepo(sqliteDb)

	omdbInfoProvider := repositoryImplementations.NewOmdbProvider("http://www.omdbapi.com", config.API_KEY)

	// services

	videoService := services.NewVideoService(sqliteVideoRepo)
	subtitleService := services.NewSubtitleService(sqliteSubtitleRepo)
	titleInfoService := services.NewTitleInfoService(sqliteTitleInfoRepo)

	fileExporer := infrastructure.NewMediaFileExplorer()

	mediaInfoSyncService := services.NewMediaInfoSyncService(
		config.MEDIA_DIR,
		fileExporer,
		omdbInfoProvider,
		videoService,
		titleInfoService,
		subtitleService,
	)

	eventHanlder := infrastructure.NewMediaEventHandlerService(
		config.MEDIA_DIR,
		titleInfoService,
		videoService,
		fileExporer,
		omdbInfoProvider,
	)

	// watch media file directory events
	watcher := infrastructure.NewMediaDirWatcher(
		config.MEDIA_DIR,
		fileExporer,
		omdbInfoProvider,
		eventHanlder,
	)
	watcher.StartWatching()

	// API REST
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))
	router.Use(middlewares.RedirectToRoot())

	controller := controllers.NewMediaInfoController(
		videoService,
		titleInfoService,
		mediaInfoSyncService,
		subtitleService,
	)

	manageData := router.Group("/api/manage")
	{
		manageData.GET("/reset", controller.ResetDatabase)
	}

	infoRouter := router.Group("/api/media")

	{
		infoRouter.GET("/titles/all", controller.GetAllTitlesInfo)
		infoRouter.GET("/titles/series", controller.GetSeries)
		infoRouter.GET("/titles/movies", controller.GetMovies)

		infoRouter.GET("/video/all", controller.GetAllVideos)
		infoRouter.GET("/video/:titleId", controller.GetVideoByTitleId)

		infoRouter.GET("/all-subtitles", controller.GetAllSubtitles)

		infoRouter.GET("/stream/:videoId", controller.StreamVideo)
		infoRouter.StaticFS("/subtitles", http.Dir(config.SUBTITLE_DIR))
	}

	// serve subtitle files

	// serve embed web app
	//
	// Due to the way paths are treated, gin does not allow the use of the
	// root path "/", so the web application must be in a specific path and
	// is served as follows:
	//
	//	staticFiles, _ := fs.Sub(staticContent, "gui/dist")
	//	router.StaticFS("/static", http.FS(staticFiles))
	//
	// To solve this we can use a middleware:
	router.Use(static.Serve("/", static.EmbedFolder(staticContent, "gui/dist")))

	err = router.Run(config.PORT)

	if err != nil {
		log.Fatal(err)
	}
}
