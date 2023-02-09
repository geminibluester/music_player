package model

const DATA_URL = "http://aseests.quhuitu.com/shengxiaoshuju.json"

type ResponseBody struct {
	Id           int    `json:"id"`
	NanShengXiao string `json:"nan_shengxiao"`
	NvShengXiao  string `json:"nv_shengxiao"`
	ZhiShu       int    `json:"zhishu"`
	JieGuo       string `json:"jieguo"`
	PingShu      string `json:"pingshu"`
}

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
