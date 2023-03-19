package httpserver

import (
	"music_player/app"
	"music_player/pkg/e"
	"music_player/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	n := ctx.DefaultQuery("n", "鼠")
	v := ctx.DefaultQuery("v", "鼠")
	svr := service.InfoService{}
	err, result := svr.FindByKey(n, v)
	if err != nil {
		appG.Response(http.StatusGone, e.ERROR, err.Error())
	}
	appG.Success(result)

}
func Create(ctx *gin.Context) {}
func Modifan(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is the edit page")
}
func Delete(ctx *gin.Context) {}
func Reload(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	appG.Success("sadjfhj")
	return
}
func NoRouteFound(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_ARTICLE, nil)
}
func ChatWithAi(ctx *gin.Context) {
	
}