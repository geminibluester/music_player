package main

import (
	"music_player/actions"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", actions.Find)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
