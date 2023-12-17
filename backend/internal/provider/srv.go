package provider

import (
	"cchart/internal/kernel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewSrv(config *kernel.Config, engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: engine,
	}
}
