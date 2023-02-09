package model

import "testing"

func TestDb(t *testing.T) {
	dsn := "file::memory:?cache=shared"
	tableName := "tbl_shengxiao"
	Db := DataBase{
		Dsn:       dsn,
		TableName: tableName,
	}
	Db.init()
	Db.initTableData()
	err, result := Db.FindByKey("鼠", "鼠")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if result.ZhiShu != "70" {
		t.Fail()
	}
}
