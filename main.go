package main

import (
	"embed"
	"log"
	"net/http"
	"time"

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

	omdbInfoProvider := repositoryImplementations.NewOmdbProvider("http://www.omdbapi.com", &config.API_KEY)

	// services

	videoService := services.NewVideoService(sqliteVideoRepo)
	subtitleService := services.NewSubtitleService(sqliteSubtitleRepo)
	titleInfoService := services.NewTitleInfoService(sqliteTitleInfoRepo)

	fileExporer := infrastructure.NewMediaFileExplorer()

	mediaInfoSyncService := services.NewMediaInfoSyncService(
		config.MEDIA_DIR,
		config.SUBTITLE_DIR,
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

	// controllers

	configController := controllers.NewConfigController(
		&config,
		omdbInfoProvider,
		titleInfoService,
		videoService,
		mediaInfoSyncService,
		subtitleService,
	)
	videoController := controllers.NewVideoController(videoService)
	subtitleController := controllers.NewSubtitleController(subtitleService)
	titleInfoController := controllers.NewTitlInfoController(titleInfoService)
	serverLogsController := controllers.NewServerLogsController()

	// API REST
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true

	serverEventChan := make(chan string)
	router.Use(cors.New(corsConfig))
	router.Use(middlewares.RedirectToRoot())
	router.Use(middlewares.HandleReq(serverEventChan))

	go Ticker(serverEventChan)

	manageData := router.Group("/api/manage")
	{
		manageData.GET("/reset", configController.ResetDatabase)
		manageData.POST("/api-key", configController.SetApiKey)

		manageData.GET("/events", middlewares.SSEHeader(), func(ctx *gin.Context) {
			serverLogsController.ServerEvents(ctx, serverEventChan)
		})
	}

	infoRouter := router.Group("/api/media")

	{
		infoRouter.GET("/titles/all", titleInfoController.GetAllTitlesInfo)
		infoRouter.GET("/titles/series", titleInfoController.GetSeries)
		infoRouter.GET("/titles/movies", titleInfoController.GetMovies)

		infoRouter.GET("/video/all", videoController.GetAllVideos)
		infoRouter.GET("/video/:titleId", videoController.GetVideoByTitleId)
		infoRouter.GET("/no-info", videoController.VideosWithNoTitleInfo)

		infoRouter.GET("/all-subtitles", subtitleController.GetAllSubtitles)
		infoRouter.POST("/video-subtitles/:subtId/:videoId", subtitleController.AssignVideoIdToSubtitles)

		infoRouter.GET("/stream/:videoId", videoController.StreamVideo)
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

	log.Printf("Server running on http://localhost%s", config.PORT)

	err = router.Run(config.PORT)

	if err != nil {
		log.Fatal(err)
	}
}

func Ticker(serverEventChan chan string) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			serverEventChan <- "tick, tack msg"
		}
	}
}
