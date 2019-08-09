package models

type DArticleItem struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Name       string `xorm:"not null default '' VARCHAR(366)"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)"`
	Collection int    `xorm:"not null default 0 comment('集') INT(10)"`
	CartoonId  int    `xorm:"not null default 0 comment('漫画id') INT(10)"`
	Content    string `xorm:"comment('各种类型数据，哈哈哈') TEXT"`
	Read       int    `xorm:"not null default 0 comment('阅读量') INT(10)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)"`
}
