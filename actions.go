package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func find(ctx *gin.Context) {
	n := ctx.DefaultQuery("n", "鼠")
	v := ctx.DefaultQuery("v", "鼠")
	err, result := dba.findByKey(n, v)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusGone,
			"message": "failed",
			"data":    err.Error(),
			"p1":      n,
			"p2":      v,
		})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    result,
		"p1":      n,
		"p2":      v,
	})
}
