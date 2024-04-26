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
) *MediaInfoController {
	return &MediaInfoController{
		TitleInfoService: titleInfoService,
		VideoService:     videoService,
		MediaInfoSync:    mediaInfoSync,
	}
}

type MediaInfoController struct {
	TitleInfoService *services.TitleInfoService
	VideoService     *services.VideoService
	MediaInfoSync    *services.MediaInfoSyncService
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

func (m *MediaInfoController) GetVideoByTitleId(ctx *gin.Context) {
	titleIdStr := ctx.Param("titleId")
	if titleIdStr == "" {
		ctx.String(http.StatusBadRequest, "Title id is required: /files/title/:titleId")
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

	ctx.String(http.StatusOK, "")

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
		ctx.Header("Content-Type", "video/mp4")
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
