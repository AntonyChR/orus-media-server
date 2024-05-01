package controllers

import (
	"log"
	"net/http"

	config "github.com/AntonyChR/orus-media-server/config"
	domain "github.com/AntonyChR/orus-media-server/internal/domain"
	gin "github.com/gin-gonic/gin"
)

type ConfigController struct {
	Config            *config.Config
	TitleInfoProvider domain.TitleInfoProvider
}

func NewConfigController(config *config.Config, titleInfoProvider domain.TitleInfoProvider) *ConfigController {
	return &ConfigController{
		Config:            config,
		TitleInfoProvider: titleInfoProvider,
	}
}

func (c *ConfigController) SetApiKey(ctx *gin.Context) {
	apiKey := ctx.Query("apiKey")
	if apiKey == "" {
		ctx.String(http.StatusOK, "API key is required")
		return
	}

	// test api key
	log.Println("Testing API key: ", apiKey)
	// save old api key 
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
