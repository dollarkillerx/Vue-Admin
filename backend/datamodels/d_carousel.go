package datamodels

type DCarousel struct {
	Id         int    `xorm:"not null pk autoincr INT(10)" form:"id" json:"id"`
	Img        string `xorm:"not null default '' comment('轮播图img') VARCHAR(266)" json:"img" form:"img"`
	Url        string `xorm:"not null default '' comment('推荐位 地址') VARCHAR(266)" json:"url" form:"url"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') VARCHAR(255)" json:"language" form:"language"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)" json:"create_time"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)" json:"-"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)" json:"-"`
}
