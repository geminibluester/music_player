package model

import (
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
			Colorful:                  true,        // Disable color
		},
	)
	d.Db, err = gorm.Open(sqlite.Open(d.Dsn), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalln("数据库初始化失败", err.Error())
	}
	if !d.initTableData() {
		log.Fatalln("init table data failed")
	}
}

// initTableData insert remote data into table
func (d *DataBase) initTableData() bool {
	log.Println("开始创建数据表", d.TableName)
	err := d.Db.Table(d.TableName).AutoMigrate(&Result{})
	if err != nil {
		log.Fatalln("创建数据表失败", d.TableName)
		return false
	}
	log.Println("数据表成功生成")
	err = d.getDataFromJson(d.Db, d.TableName)
	return err == nil
}

// getDataFromJson get remote json data by DATA_URL
func (d *DataBase) getDataFromJson(db *gorm.DB, tableName string) error {
	log.Println("开始获取远程数据")
	rsp := []ResponseBody{}
	err := gout.GET(DATA_URL).BindJSON(&rsp).Do()
	if err != nil {
		log.Fatal("获取远程数据失败", err.Error())
		return err
	}
	log.Println("远程数据获取成功，开始插入数据表")
	x := 0
	for _, value := range rsp {
		record := Result{
			ID:           value.Id,
			NanShengXiao: value.NanShengXiao,
			NvShengXiao:  value.NvShengXiao,
			ZhiShu:       value.ZhiShu,
			JieGuo:       value.JieGuo,
			PingShu:      value.PingShu,
		}
		d.Db.Table(d.TableName).Create(&record)
		x = x + 1
		// log.Printf("正在入库第%d条数据\n", k)
	}
	log.Printf("远程数据插入数据表成功,共插入%d条", x)
	return nil
}

// findByKey find the record in table by n and v
func (d *DataBase) FindByKey(n string, v string) (ApiResult, error) {
	var result ApiResult
	if err := d.Db.Table(d.TableName).Where("nan_shengxiao =? and nv_shengxiao=?", n, v).Scan(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
