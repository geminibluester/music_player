package model

import (
	"music_player/pkg/contant"
	"testing"
)

func TestDb(t *testing.T) {
	dsn := "file::memory:?cache=shared"
	tableName := "tbl_shengxiao"
	Db := DataBase{
		Dsn:       dsn,
		TableName: tableName,
	}
	Db.init()
	Db.initTableData()
	result, err := Db.FindByKey(contant.DEFAULT_CHAR, contant.DEFAULT_CHAR)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if result.ZhiShu != "70" {
		t.Fail()
	}
}
