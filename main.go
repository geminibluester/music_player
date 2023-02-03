package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Result struct {
	ID           int    `gorm:"column:id" json:"id"`
	NanShengXiao string `gorm:"column:nan_shengxiao" json:"nan"`
	NvShengXiao  string `gorm:"column:nv_shengxiao" json:"nv"`
	ZhiShu       string `gorm:"column:zhishu" json:"zhishu"`
	JieGuo       string `gorm:"column:jieguo" json:"jieguo"`
	PingShu      string `gorm:"column:pingshu" json:"pingshu"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("oliyo_app.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	var result Result
	if err = db.Raw("SELECT * FROM tbl_shengxiao").Scan(&result).Error; err != nil {
		panic(err)
	}
	fmt.Println(result)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})

	r.Run("127.0.0.1:8080") // 监听并在 0.0.0.0:8080 上启动服务
}
