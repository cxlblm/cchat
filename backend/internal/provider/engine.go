package provider

import (
	"cchart/internal/kernel"
	"github.com/gin-gonic/gin"
)

func NewEngine(config *kernel.Config) *gin.Engine {
	return newEngine()
}

func newEngine() *gin.Engine {
	r := gin.Default()

	//err := r.SetTrustedProxies(nil)
	//if err != nil {
	//	return nil
	//}
	return r
}
