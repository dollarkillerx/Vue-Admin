package datamodels

type DNav struct {
	Id       int    `xorm:"not null pk autoincr INT(10)" json:"id" form:"id"`
	Language string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)" json:"language" form:"language"`
	Name     string `xorm:"not null default '' comment('导航名称') VARCHAR(255)" json:"name" form:"name"`
	Url      string `xorm:"not null default '' comment('导航地址') VARCHAR(366)" json:"name" form:"name"`
	Weight   int    `xorm:"not null default 0 comment('权重') INT(11)" json:"weight" form:"weight"`
}
