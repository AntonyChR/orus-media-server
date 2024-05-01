package controllers

import (
	"net/http"
	"strconv"

	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
	gin "github.com/gin-gonic/gin"
)

func NewSubtitleController(subtitleService *services.SubtitleService) *SubtitleController {
	return &SubtitleController{
		SubtitleService: subtitleService,
	}
}

type SubtitleController struct {
	SubtitleService *services.SubtitleService
}

func (s *SubtitleController) GetAllSubtitles(ctx *gin.Context) {
	data, err := s.SubtitleService.GetAll()
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (s *SubtitleController) GetSubtitlesByVideoId(ctx *gin.Context) {
	videoIdStr := ctx.Param("videoId")
	if videoIdStr == "" {
		ctx.String(http.StatusBadRequest, "Video id is required: /api/media/video-subtitles/:videoId")
		return
	}

	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid video id")
		return
	}

	data, err := s.SubtitleService.GetByVideoId(uint(videoId))
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (s *SubtitleController) AssignVideoIdToSubtitles(ctx *gin.Context) {
	videoIdStr := ctx.Param("videoId")
	if videoIdStr == "" {
		ctx.String(http.StatusBadRequest, "Video id is required: /api/media/assign-video-id/:videoId")
		return
	}

	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid video id")
		return
	}

	subtIdStr := ctx.Param("subtId")
	if subtIdStr == "" {
		ctx.String(http.StatusBadRequest, "Subtitle id is required: /api/media/assign-video-id/:videoId/:subtId")
		return
	}

	subtId, err := strconv.Atoi(subtIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid subtitle id")
		return
	}

	err = s.SubtitleService.SetVideoId(uint(subtId), uint(videoId))
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.String(http.StatusOK, "Video id assigned to subtitles")
}
