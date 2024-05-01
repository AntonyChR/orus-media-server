package controllers

import (
	"log"
	"net/http"

	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
	gin "github.com/gin-gonic/gin"
)

func NewTitlInfoController(titleInfoService *services.TitleInfoService) *TitlInfoController {
	return &TitlInfoController{
		TitleInfoService: titleInfoService,
	}
}

type TitlInfoController struct {
	TitleInfoService *services.TitleInfoService
}

func (t *TitlInfoController) GetSeries(ctx *gin.Context) {
	data, err := t.TitleInfoService.GetSeries()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (t *TitlInfoController) GetMovies(ctx *gin.Context) {
	data, err := t.TitleInfoService.GetMovies()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (t *TitlInfoController) GetAllTitlesInfo(ctx *gin.Context) {
	data, err := t.TitleInfoService.GetAll()
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}
