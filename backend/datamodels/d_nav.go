package datamodels

type DNav struct {
	Id       int    `xorm:"not null pk autoincr INT(10)"`
	Language string `xorm:"not null default '' comment('语言编号缩写') index VARCHAR(255)"`
	Name     string `xorm:"not null default '' comment('导航名称') VARCHAR(255)"`
	Url      string `xorm:"not null default '' comment('导航地址') VARCHAR(366)"`
	Weight   int    `xorm:"not null default 0 comment('权重') INT(11)"`
}
