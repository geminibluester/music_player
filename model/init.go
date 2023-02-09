package model

var Dba DataBase

func init() {
	dsn := "file::memory:?cache=shared"
	tableName := "tbl_shengxiao"
	Dba = DataBase{
		Dsn:       dsn,
		TableName: tableName,
	}
	Dba.init()
}
