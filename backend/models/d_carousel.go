package models

type DCarousel struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Img        string `xorm:"not null default '' comment('轮播图img') VARCHAR(266)"`
	Url        string `xorm:"not null default '' comment('推荐位 地址') VARCHAR(266)"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') VARCHAR(255)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)"`
}
