package httpserver

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer(addr string, mode string) {
	r := setupRouter(mode)
	log.Printf("服务即将启动，监听于%s", addr)
	r.Run(addr) // 监听并在 0.0.0.0:8080 上启动服务
}
func setupRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", Find)
	return r
}
