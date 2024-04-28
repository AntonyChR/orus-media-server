package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	services "github.com/AntonyChR/orus-media-server/internal/domain/services"
	gin "github.com/gin-gonic/gin"
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

func (m *MediaInfoController) GetSeries(ctx *gin.Context) {
	data, err := m.TitleInfoService.GetSeries()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (m *MediaInfoController) GetMovies(ctx *gin.Context) {
	data, err := m.TitleInfoService.GetMovies()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (m *MediaInfoController) GetAllTitlesInfo(ctx *gin.Context) {
	data, err := m.TitleInfoService.GetAll()
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}
func (m *MediaInfoController) GetAllVideos(ctx *gin.Context) {
	data, err := m.VideoService.GetAll()
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (m *MediaInfoController) GetAllSubtitles(ctx *gin.Context) {
	data, err := m.SubtitleService.GetAll()
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (m *MediaInfoController) AssignVideoIdToSubtitles(ctx *gin.Context) {
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

	err = m.SubtitleService.SetVideoId(uint(subtId), uint(videoId))
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.String(http.StatusOK, "Video id assigned to subtitles")
}

func (m *MediaInfoController) GetSubtitlesByVideoId(ctx *gin.Context) {
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

	data, err := m.SubtitleService.GetByVideoId(uint(videoId))
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (m *MediaInfoController) GetVideoByTitleId(ctx *gin.Context) {
	titleIdStr := ctx.Param("titleId")
	if titleIdStr == "" {
		ctx.String(http.StatusBadRequest, "Title id is required: /api/media/video/:titleId")
		return
	}

	titleId, err := strconv.ParseUint(titleIdStr, 10, 32)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid title id, should be an int number")
		return
	}

	data, err := m.VideoService.GetByTitleId(uint(titleId))
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (m *MediaInfoController) ResetDatabase(ctx *gin.Context) {
	err := m.VideoService.Reset()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	err = m.TitleInfoService.Reset()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadGateway, "Server error")
		return
	}

	err = m.MediaInfoSync.GetTitleInfoAboutAllMediaFiles()

	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	if err := m.SubtitleService.Reset(); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	if err := m.MediaInfoSync.ScanSubtitles(); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadGateway, "Server error")
		return
	}

	ctx.String(http.StatusOK, "Database reseted")

}

func (m *MediaInfoController) StreamVideo(ctx *gin.Context) {
	videoIdStr := ctx.Param("videoId")

	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid video id")
		return
	}

	videoInfo, err := m.VideoService.GetByVideoId(uint(videoId))

	if err != nil {
		ctx.String(http.StatusNotFound, "Video not found")
		return
	}

	file, err := os.Open(videoInfo.Path)

	if err != nil {
		ctx.String(http.StatusNotFound, "Video not found")
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, "Server error")

	}
	fileSize := info.Size()

	rangeHeader := ctx.GetHeader("range")

	if rangeHeader != "" {
		rpl := strings.ReplaceAll(rangeHeader, "bytes=", "")
		rangeParts := strings.Split(rpl, "-")

		startStr := rangeParts[0]
		start, _ := strconv.ParseInt(startStr, 10, 64)

		var end int64
		if len(rangeParts) > 1 && rangeParts[1] != "" {
			end, _ = strconv.ParseInt(rangeParts[1], 10, 64)
		} else {
			end = fileSize - 1
		}
		chunkSize := end - start + 1
		chunkSizeStr := strconv.FormatInt(chunkSize, 10)

		contentRangeHeader := fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize)

		ctx.Header("Content-Range", contentRangeHeader)
		ctx.Header("Accept-Ranges", "bytes")
		ctx.Header("Content-Type", "video/"+videoInfo.Ext)
		ctx.Header("Content-Length", chunkSizeStr)
		ctx.Status(http.StatusPartialContent)

		_, seekErr := file.Seek(start, io.SeekStart)
		if seekErr != nil {
			log.Println("Error seeking to start position:", seekErr)
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		_, copyErr := io.CopyN(ctx.Writer, file, chunkSize)
		if copyErr != nil {
			if !ctx.Writer.Written() {
				log.Println("Connection closed by client")
				return
			}
			log.Println("Error copying bytes to response: ", copyErr)
			ctx.String(http.StatusInternalServerError, "Server error")
		}
	} else {
		fileSizeStr := strconv.FormatInt(fileSize, 10)

		ctx.Header("Content-Type", "video/mp4")
		ctx.Header("Content-Length", fileSizeStr)
		_, copyErr := io.Copy(ctx.Writer, file)

		if copyErr != nil {
			log.Println("Error copying bytes to response: ", copyErr)
			ctx.String(http.StatusInternalServerError, "Server error")
		}
	}

}
