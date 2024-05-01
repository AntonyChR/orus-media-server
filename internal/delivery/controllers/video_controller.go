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

func NewVideoController(videoService *services.VideoService) *VideoController {
	return &VideoController{VideoService: videoService}
}

type VideoController struct {
	VideoService *services.VideoService
}

func (v *VideoController) GetAllVideos(ctx *gin.Context) {
	data, err := v.VideoService.GetAll()
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (v *VideoController) GetVideoByTitleId(ctx *gin.Context) {
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

	data, err := v.VideoService.GetByTitleId(uint(titleId))
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (v *VideoController) VideosWithNoTitleInfo(ctx *gin.Context) {
	data, err := v.VideoService.GetVideosWithNoTitleInfo()
	if err != nil {
		ctx.String(http.StatusNotFound, "Not found")
		return
	}
	ctx.JSON(http.StatusOK, data)

}

func (v *VideoController) StreamVideo(ctx *gin.Context) {
	videoIdStr := ctx.Param("videoId")

	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid video id")
		return
	}

	videoInfo, err := v.VideoService.GetByVideoId(uint(videoId))

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
