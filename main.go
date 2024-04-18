package main

import (
	"embed"
	"log"
	"os"

	controllers "github.com/AntonyChR/orus-media-server/internal/controllers"
	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
	infrastructure "github.com/AntonyChR/orus-media-server/internal/infrastructure"
	repositoryImplementations "github.com/AntonyChR/orus-media-server/internal/infrastructure/repositories"
	cors "github.com/gin-contrib/cors"

	static "github.com/gin-contrib/static"
	gin "github.com/gin-gonic/gin"
	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

//go:embed gui/dist/*
var staticContent embed.FS

func main() {
	//TODO: read config

	sqliteDb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	port := ":3002"
	API_KEY := os.Getenv("API_KEY")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))

	// Instance repositories

	sqliteFileInfoRepo := repositoryImplementations.NewSqliteFileInfoRepo(sqliteDb)
	sqliteTitleInfoRepo := repositoryImplementations.NewSqliteTitleInfoRepo(sqliteDb)

	omdbInfoProvider := repositoryImplementations.NewOmdbProvider("http://www.omdbapi.com", API_KEY)

	// services

	fileInfoService := services.NewTitleInfoService(sqliteTitleInfoRepo)
	titleInfoService := services.NewFileInfoService(sqliteFileInfoRepo)

	fileExporer := infrastructure.NewMediaFileExplorer()

	// Initialize services
	mediaInfoSyncService := services.NewMediaInfoSyncService(
		"./temp",
		fileExporer,
		omdbInfoProvider,
		titleInfoService,

		fileInfoService,
	)

	// watch media file directory events
	watcher := infrastructure.NewMediaDirWatcher(
		"./temp",
		fileExporer,
		omdbInfoProvider,
		sqliteTitleInfoRepo,
		sqliteFileInfoRepo,
	)

	go watcher.WatchDirectoryEvents()
	go watcher.ListenMediaEvents()

	// API REST
	controller := controllers.NewMediaInfoController(
		titleInfoService,
		fileInfoService,
		mediaInfoSyncService)

	manageData := router.Group("/api/manage")
	{
		manageData.GET("/reset", controller.ResetDatabase)
	}

	infoRouter := router.Group("/api/info")

	{
		infoRouter.GET("/titles/all", controller.GetAllTitlesInfo)
		infoRouter.GET("/titles/series", controller.GetSeries)
		infoRouter.GET("/titles/movies", controller.GetMovies)

		infoRouter.GET("/files/all", controller.GetAllFilesInfo)
		infoRouter.GET("/files/title/:titleId", controller.GetFileInfoByTitleId)

		infoRouter.GET("/video/:videoId", controller.StreamVideo)

	}

	//staticFiles, _ := fs.Sub(staticContent, "gui/dist")
	//router.StaticFS("/static", http.FS(staticFiles))

	router.Use(static.Serve("/", static.EmbedFolder(staticContent, "gui/dist")))

	err = router.Run(port)

	if err != nil {
		log.Fatal(err)
	}
}
