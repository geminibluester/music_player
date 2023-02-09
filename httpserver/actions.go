package httpserver

import (
	"music_player/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	n := ctx.DefaultQuery("n", "鼠")
	v := ctx.DefaultQuery("v", "鼠")
	err, result := model.Dba.FindByKey(n, v)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusGone,
			"message": "failed",
			"data":    err.Error(),
		})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    result,
	})
}
