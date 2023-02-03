package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	ID           int    `gorm:"column:id"`
	NanShengXiao string `gorm:"column:nan_shengxiao" json:"nan"`
	NvShengXiao  string `gorm:"column:nv_shengxiao" json:"nv"`
	ZhiShu       string `gorm:"column:zhishu" json:"zhishu"`
	JieGuo       string `gorm:"column:jieguo" json:"jieguo"`
	PingShu      string `gorm:"column:pingshu" json:"pingshu"`
	Age          int
}

func main() {
	dsn := "root:gemini4094@tcp(dy.quhuitu.com:3306)/oliyo_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	var result Result
	db.Raw("SELECT * FROM tbl_shengxiao").Scan(&result)
	fmt.Println(result)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
