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
	utils "github.com/AntonyChR/orus-media-server/utils"
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

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	if err := infrastructure.CheckDirectories(cfg.MEDIA_DIR, cfg.SUBTITLE_DIR); err != nil {
		log.Fatal(err)
	}

	sqliteDb, err := gorm.Open(sqlite.Open(cfg.DB_PATH), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	logSSeManager := infrastructure.NewLogSSEManager()
	go logSSeManager.Start()

	// Instance repositories

	sqliteVideoRepo := repositoryImplementations.NewSqliteVideoRepo(sqliteDb)
	sqliteSubtitleRepo := repositoryImplementations.NewSqliteSubtitleRepo(sqliteDb)
	sqliteTitleInfoRepo := repositoryImplementations.NewSqliteTitleInfoRepo(sqliteDb)

	omdbInfoProvider := repositoryImplementations.NewOmdbProvider("http://www.omdbapi.com", &cfg.API_KEY)

	// services

	videoService := services.NewVideoService(sqliteVideoRepo)
	subtitleService := services.NewSubtitleService(sqliteSubtitleRepo)
	titleInfoService := services.NewTitleInfoService(sqliteTitleInfoRepo)

	fileExporer := infrastructure.NewMediaFileExplorer()

	mediaInfoSyncService := services.NewMediaInfoSyncService(
		cfg.MEDIA_DIR,
		cfg.SUBTITLE_DIR,
		fileExporer,
		omdbInfoProvider,
		videoService,
		titleInfoService,
		subtitleService,
	)

	// watch media file directory events
	eventHanlder := infrastructure.NewMediaEventHandlerService(
		cfg.MEDIA_DIR,
		titleInfoService,
		videoService,
		fileExporer,
		omdbInfoProvider,
	)

	watcher := infrastructure.NewMediaDirWatcher(
		cfg.MEDIA_DIR,
		fileExporer,
		omdbInfoProvider,
		eventHanlder,
		logSSeManager.LogsChannel,
	)
	watcher.StartWatching()

	// controllers

	configController := controllers.NewConfigController(
		&cfg,
		omdbInfoProvider,
		titleInfoService,
		videoService,
		mediaInfoSyncService,
		subtitleService,
	)
	videoController := controllers.NewVideoController(videoService)
	subtitleController := controllers.NewSubtitleController(subtitleService)
	titleInfoController := controllers.NewTitlInfoController(titleInfoService)
	serverLogsController := controllers.NewServerLogsController(logSSeManager)

	// API REST
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))
	router.Use(middlewares.RedirectToRoot())

	router.Use(middlewares.HandleReq(logSSeManager.LogsChannel))

	manageData := router.Group("/api/manage")
	{
		manageData.GET("/reset", configController.ResetDatabase)
		manageData.POST("/api-key", configController.SetApiKey)

		manageData.GET("/events", middlewares.SSEHeader(), serverLogsController.ServerEvents)
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
		infoRouter.StaticFS("/subtitles", http.Dir(cfg.SUBTITLE_DIR))
	}

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

	localIp, err := utils.GetLocalIP()

	if err != nil {
		log.Println("Error getting local ip address")
		log.Println(err)
	}

	utils.PrintServerInfo(localIp, cfg.PORT)

	err = router.Run(cfg.PORT)

	if err != nil {
		log.Fatal(err)
	}
}
