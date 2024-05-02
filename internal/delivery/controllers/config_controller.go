package controllers

import (
	"log"
	"net/http"

	config "github.com/AntonyChR/orus-media-server/config"
	domain "github.com/AntonyChR/orus-media-server/internal/domain"
	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
	gin "github.com/gin-gonic/gin"
)

func NewConfigController(
	config *config.Config,
	titleInfoProvider domain.TitleInfoProvider,
	titleInfoService *services.TitleInfoService,
	videoService *services.VideoService,
	mediaInfoSync *services.MediaInfoSyncService,
	subtitleService *services.SubtitleService,
) *ConfigController {
	return &ConfigController{
		Config:            config,
		TitleInfoProvider: titleInfoProvider,
		TitleInfoService:  titleInfoService,
		VideoService:      videoService,
		MediaInfoSync:     mediaInfoSync,
		SubtitleService:   subtitleService,
	}
}

type ConfigController struct {
	Config            *config.Config
	TitleInfoProvider domain.TitleInfoProvider
	TitleInfoService  *services.TitleInfoService
	VideoService      *services.VideoService
	MediaInfoSync     *services.MediaInfoSyncService
	SubtitleService   *services.SubtitleService
}

func (c *ConfigController) SetApiKey(ctx *gin.Context) {
	apiKey := ctx.Query("apiKey")
	if apiKey == "" {
		ctx.String(http.StatusOK, "API key is required")
		return
	}

	// test api key
	log.Println("Testing API key: ", apiKey)
	old := c.Config.API_KEY
	c.Config.API_KEY = apiKey
	if _, err := c.TitleInfoProvider.Search("The Matrix(1999).mp4"); err != nil {
		c.Config.API_KEY = old
		ctx.String(http.StatusBadRequest, "Invalid API key")
		return
	}

	c.Config.API_KEY = apiKey
	ctx.String(200, "API key set successfully")
}

func (c *ConfigController) ResetDatabase(ctx *gin.Context) {
	err := c.VideoService.Reset()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	err = c.TitleInfoService.Reset()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadGateway, "Server error")
		return
	}

	err = c.MediaInfoSync.GetTitleInfoAboutAllMediaFiles()

	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	if err := c.SubtitleService.Reset(); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	if err := c.MediaInfoSync.ScanSubtitles(); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	ctx.String(http.StatusOK, "Database reseted")

}
