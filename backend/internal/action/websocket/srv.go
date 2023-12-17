package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
)

func NewUpgrader(logger *slog.Logger) *UpgraderAction {
	hub := newHub()
	go hub.run()
	return &UpgraderAction{
		logger: logger,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		hub: hub,
	}
}

type UpgraderAction struct {
	logger   *slog.Logger
	upgrader websocket.Upgrader
	hub      *Hub
}

func (s *UpgraderAction) Upgrade(c *gin.Context) {
	ws, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		s.logger.Error("websocket upgrade error:" + err.Error())
		return
	}

	name := c.Request.URL.Query().Get("name")
	client := &Client{hub: s.hub, conn: ws, send: make(chan string, 256), log: s.logger, name: name}
	client.hub.register <- client
	client.hub.broadcast <- client.name + " joined "

	go client.write()
	go client.read()
}
