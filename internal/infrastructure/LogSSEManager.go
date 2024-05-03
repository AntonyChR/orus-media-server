package infrastructure

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewLogSSEManager() *LogSSEManager {
	return &LogSSEManager{
		Clients:     make(map[string]*gin.Context),
		LogsChannel: make(chan string),
	}
}

type LogSSEManager struct {
	Clients     map[string]*gin.Context
	LogsChannel chan string
}

func (l *LogSSEManager) Register(contextRequest *gin.Context, clientId string) {
	l.Clients[clientId] = contextRequest
	contextRequest.Set("clientId", clientId)
}

func (l *LogSSEManager) Unregister(clientId string) {
	delete(l.Clients, clientId)
}

func (l *LogSSEManager) Start() {
	for {
		log := <-l.LogsChannel
		l.Broadcast(log)
	}
}

func (l *LogSSEManager) Broadcast(content string) {

	if len(l.Clients) == 0 {
		return
	}
	logId := uuid.New().String()
	data := fmt.Sprintf("{\"content\": \"%s\",\"id\":\"%s\"}", content, logId)

	for _, client := range l.Clients {
		fmt.Fprintf(client.Writer, "data: %s\n\n", data)
		client.Writer.Flush()
	}
}
