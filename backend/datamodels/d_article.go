package datamodels

type DArticle struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Name       string `xorm:"not null default '' VARCHAR(366)"`
	Img        string `xorm:"not null default '' VARCHAR(366)"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)"`
	State      int    `xorm:"not null default 0 comment('状态0连载中1完结2停更') TINYINT(3)"`
	Author     string `xorm:"not null default '' comment('作者') VARCHAR(366)"`
	Describe   string `xorm:"comment('描述') TEXT"`
	SortId     int    `xorm:"not null default 0 comment('分类id') index INT(10)"`
	Read       int    `xorm:"not null default 0 comment('阅读量') INT(10)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)"`
}
