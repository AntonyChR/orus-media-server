package controllers

import (
	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
)

func NewMediaInfoController(
	videoService *services.VideoService,
	titleInfoService *services.TitleInfoService,
	mediaInfoSync *services.MediaInfoSyncService,
	subtitleService *services.SubtitleService,
) *MediaInfoController {
	return &MediaInfoController{
		TitleInfoService: titleInfoService,
		VideoService:     videoService,
		MediaInfoSync:    mediaInfoSync,
		SubtitleService:  subtitleService,
	}
}

type MediaInfoController struct {
	TitleInfoService *services.TitleInfoService
	VideoService     *services.VideoService
	MediaInfoSync    *services.MediaInfoSyncService
	SubtitleService  *services.SubtitleService
}
