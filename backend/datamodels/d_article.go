package datamodels

type DArticle struct {
	Id         int    `xorm:"not null pk autoincr INT(10)" form:"id" json:"id"`
	Name       string `xorm:"not null default '' VARCHAR(366)" form:"name" json:"name"`
	Img        string `xorm:"not null default '' VARCHAR(366)" form:"img" json:"img"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)" form:"language" json:"language"`
	State      int    `xorm:"not null default 0 comment('状态0连载中1完结2停更') TINYINT(3)" form:"state" json:"state"`
	Author     string `xorm:"not null default '' comment('作者') VARCHAR(366)" form:"author" json:"author"`
	Describe   string `xorm:"comment('描述') TEXT" form:"describe" json:"describe"`
	SortId     int    `xorm:"not null default 0 comment('分类id') index INT(10)" form:"sort_id" json:"sort_id"`
	Read       int    `xorm:"not null default 0 comment('阅读量') INT(10)" form:"read" json:"read"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)" json:"create_time"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)" json:"-"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)" json:"-"`
}
