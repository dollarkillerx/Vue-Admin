package datamodels

type DArticleItem struct {
	Id         int    `xorm:"not null pk autoincr INT(10)" form:"id" json:"id"`
	Name       string `xorm:"not null default '' VARCHAR(366)" form:"name" json:"name"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)" form:"language" json:"language"`
	Collection int    `xorm:"not null default 0 comment('集') INT(10)" json:"collection" form:"collection"`
	CartoonId  int    `xorm:"not null default 0 comment('漫画id') INT(10)" json:"cartoon_id" form:"cartoon_id"`
	Content    string `xorm:"comment('各种类型数据，哈哈哈') TEXT" json:"content" form:"content"`
	Read       int    `xorm:"not null default 0 comment('阅读量') INT(10)" json:"read" form:"read"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)" json:"create_time"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)" json:"-"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)" json:"-"`
}
