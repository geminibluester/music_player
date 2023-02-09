package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/guonaihong/gout"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataBase struct {
	Dsn       string
	Db        *gorm.DB
	TableName string
}

// init init dabase handel and insert remote data into table
func (d *DataBase) init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	d.Db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	d.initTableData()
}

// initTableData insert remote data into table
func (d *DataBase) initTableData() bool {
	fmt.Println("开始创建数据表", d.TableName)
	err := d.Db.Table(d.TableName).AutoMigrate(&Result{})
	if err != nil {
		fmt.Println("创建数据表失败", d.TableName)
	}
	d.getDataFromJson(d.Db, d.TableName)
	return true
}

// getDataFromJson get remote json data by DATA_URL
func (d *DataBase) getDataFromJson(db *gorm.DB, tableName string) {
	rsp := []ResponseBody{}
	err := gout.GET(DATA_URL).BindJSON(&rsp).Do()
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
		d.Db.Table(d.TableName).Create(&record)
		fmt.Printf("正在入库第%d条数据\n", k)
	}
	fmt.Println("远程数据入库成功")
}

// findByKey find the record in table by n and v
func (d *DataBase) findByKey(n string, v string) (error, ApiResult) {
	var result ApiResult
	if err := d.Db.Table(d.TableName).Where("nan_shengxiao =? and nv_shengxiao=?", n, v).Scan(&result).Error; err != nil {
		return err, result
	}
	return nil, result
}
