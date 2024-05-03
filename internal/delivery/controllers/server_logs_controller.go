package controllers

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

func NewServerLogsController() *ServerLogsController {
	return &ServerLogsController{}
}

type ServerLogsController struct {
}

func (s *ServerLogsController) ServerEvents(c *gin.Context, msgChann chan string) {

	ctx := c.Request.Context()
	for {
		select {
		case msg := <-msgChann:
			id := uuid.New().String()
			data := fmt.Sprintf("{\"id\":\"%s\", \"content\": \"%s\"}", id, msg)
			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()
		case <-ctx.Done():
			return
		}
	}
}
