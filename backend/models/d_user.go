package models

type DUser struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Name       string `xorm:"not null default '' VARCHAR(255)"`
	Email      string `xorm:"not null default '' VARCHAR(355)"`
	Password   string `xorm:"not null default '' CHAR(32)"`
	Avatar     string `xorm:"not null default '' comment('头像地址') VARCHAR(355)"`
	IsAdmin    int    `xorm:"not null default 0 comment('0普通用户1管理员') index TINYINT(4)"`
	CreateTime int    `xorm:"not null default 0 INT(10)"`
}
