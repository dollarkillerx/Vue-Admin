package models

type DSort struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Name       string `xorm:"not null default '' VARCHAR(266)"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)"`
	Describe   string `xorm:"not null default '' comment('分类描述') VARCHAR(366)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)"`
}
