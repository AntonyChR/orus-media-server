package main

import (
	"embed"
	"log"

	"github.com/AntonyChR/orus-media-server/config"
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

	sqliteDb, err := gorm.Open(sqlite.Open(config.DB_PATH), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Instance repositories

	sqliteFileInfoRepo := repositoryImplementations.NewSqliteFileInfoRepo(sqliteDb)
	sqliteTitleInfoRepo := repositoryImplementations.NewSqliteTitleInfoRepo(sqliteDb)

	omdbInfoProvider := repositoryImplementations.NewOmdbProvider("http://www.omdbapi.com", config.API_KEY)

	// services

	fileInfoService := services.NewTitleInfoService(sqliteTitleInfoRepo)
	titleInfoService := services.NewFileInfoService(sqliteFileInfoRepo)

	fileExporer := infrastructure.NewMediaFileExplorer()

	// Initialize services
	mediaInfoSyncService := services.NewMediaInfoSyncService(
		config.MEDIA_DIR,
		fileExporer,
		omdbInfoProvider,
		titleInfoService,

		fileInfoService,
	)

	// watch media file directory events
	watcher := infrastructure.NewMediaDirWatcher(
		config.MEDIA_DIR,
		fileExporer,
		omdbInfoProvider,
		sqliteTitleInfoRepo,
		sqliteFileInfoRepo,
	)

	go watcher.WatchDirectoryEvents()
	go watcher.ListenMediaEvents()

	// API REST
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))
	router.Use(middlewares.Redirect())
	controller := controllers.NewMediaInfoController(
		titleInfoService,
		fileInfoService,
		mediaInfoSyncService)

	manageData := router.Group("/api/manage")
	{
		manageData.GET("/reset", controller.ResetDatabase)
	}

	infoRouter := router.Group("/api/media")

	{
		infoRouter.GET("/titles/all", controller.GetAllTitlesInfo)
		infoRouter.GET("/titles/series", controller.GetSeries)
		infoRouter.GET("/titles/movies", controller.GetMovies)

		infoRouter.GET("/files/all", controller.GetAllFilesInfo)
		infoRouter.GET("/files/:titleId", controller.GetFileInfoByTitleId)

		infoRouter.GET("/video/:videoId", controller.StreamVideo)

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

	err = router.Run(config.PORT)

	if err != nil {
		log.Fatal(err)
	}
}
