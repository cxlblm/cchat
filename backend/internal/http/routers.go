package http

import (
	"cchart/internal/action/websocket"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(websocket.NewUpgrader, wire.Struct(new(Router), "*"))

type Router struct {
	Engine            *gin.Engine
	WebsocketUpgrader *websocket.UpgraderAction
}

func (r *Router) Register() {
	engine := r.Engine
	engine.GET("/ws", r.WebsocketUpgrader.Upgrade)
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"abc": "abc"})
	})
}
