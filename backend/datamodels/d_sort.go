package datamodels

type DSort struct {
	Id         int    `xorm:"not null pk autoincr INT(10)" json:"id" form:"id"`
	Name       string `xorm:"not null default '' VARCHAR(266)" json:"name" form:"name"`
	Language   string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)" json:"language" form:"language"`
	Describe   string `xorm:"not null default '' comment('分类描述') VARCHAR(366)" json:"describe" form:"describe"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)" json:"create_time"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)" json:"-"`
	DeleteTime int    `xorm:"not null default 0 comment('软删除') index INT(10)" json:"-"`
}
