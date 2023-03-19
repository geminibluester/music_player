package httpserver

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer(addr string, mode string) {
	r := setupRouter(mode)
	log.Printf("服务启动，监听于%s", addr)
	err := r.Run(addr) // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		log.Fatalln("服务启动失败", err.Error())
	}
}
func setupRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(NoRouteFound)
	r.GET("/", Find)
	r.GET("/favicon.ico", func(ctx *gin.Context) { ctx.String(http.StatusOK, "") })
	r.POST("/ai", ChatWithAi)
	manage := r.Group("/manage")
	manage.POST("/create", Create)
	manage.GET("/edit", Modifan)
	manage.GET("/reload", Reload)
	return r
}
