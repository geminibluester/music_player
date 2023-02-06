package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DATA_URL = "http://aseests.quhuitu.com/shengxiaoshuju.json"

type Result struct {
	ID           int    `gorm:"column:id;primaryKey;not null" json:"id"`
	NanShengXiao string `gorm:"column:nan_shengxiao;not null" json:"nan"`
	NvShengXiao  string `gorm:"column:nv_shengxiao;not null" json:"nv"`
	ZhiShu       int    `gorm:"column:zhishu;not null" json:"zhishu"`
	JieGuo       string `gorm:"column:jieguo;not null" json:"jieguo"`
	PingShu      string `gorm:"column:pingshu;not null" json:"pingshu"`
}
type ApiResult struct {
	NanShengXiao string `gorm:"column:nan_shengxiao;not null" json:"nan"`
	NvShengXiao  string `gorm:"column:nv_shengxiao;not null" json:"nv"`
	ZhiShu       string `gorm:"column:zhishu;not null" json:"zhishu"`
	JieGuo       string `gorm:"column:jieguo;not null" json:"jieguo"`
	PingShu      string `gorm:"column:pingshu;not null" json:"pingshu"`
}
type ResponseBody struct {
	Id           int    `json:"id"`
	NanShengXiao string `json:"nan_shengxiao"`
	NvShengXiao  string `json:"nv_shengxiao"`
	ZhiShu       int    `json:"zhishu"`
	JieGuo       string `json:"jieguo"`
	PingShu      string `json:"pingshu"`
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	tableName := "tbl_shengxiao"
	fmt.Println("开始创建数据表", tableName)
	err = db.Table(tableName).AutoMigrate(&Result{})
	if err != nil {
		fmt.Println("创建数据表失败", tableName)
	}
	rsp := []ResponseBody{}
	err = gout.GET(DATA_URL).BindJSON(&rsp).Do()
	if err != nil {
		fmt.Println("获取远程数据失败", err.Error())
	}
	for k, value := range rsp {
		record := Result{
			ID:           value.Id,
			NanShengXiao: value.NanShengXiao,
			NvShengXiao:  value.NvShengXiao,
			ZhiShu:       value.ZhiShu,
			JieGuo:       value.JieGuo,
			PingShu:      value.PingShu,
		}
		db.Table(tableName).Create(&record)
		fmt.Printf("正在入库第%d条数据\n", k)
	}
	fmt.Println("远程数据入库成功")
	var result ApiResult
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(ctx *gin.Context) {
		n := ctx.DefaultQuery("n", "鼠")
		v := ctx.DefaultQuery("v", "鼠")
		if err = db.Table(tableName).Where("nan_shengxiao =? and nv_shengxiao=?", n, v).Scan(&result).Error; err != nil {
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
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
