package controllers

import (
	"log"

	"github.com/AntonyChR/orus-media-server/internal/infrastructure"
	gin "github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

func NewServerLogsController(logSSEManager *infrastructure.LogSSEManager) *ServerLogsController {
	return &ServerLogsController{
		LogSSEManager: logSSEManager,
	}
}

type ServerLogsController struct {
	LogSSEManager *infrastructure.LogSSEManager
}

func (s *ServerLogsController) ServerEvents(c *gin.Context) {

	ctx := c.Request.Context()
	clientId := uuid.New().String()

	s.LogSSEManager.Register(c, clientId)
	log.Println("Client connected: ", clientId)

	for {
		<-ctx.Done()
		log.Println("Client disconnected: ", clientId)
		s.LogSSEManager.Unregister(clientId)
		return
	}
}
