package datamodels

type DUser struct {
	Id         int    `xorm:"not null pk autoincr INT(10)" json:"id"`
	Name       string `xorm:"not null default '' VARCHAR(255)" form:"name" json:"name"`
	Email      string `xorm:"not null default '' VARCHAR(355)" form:"email" json:"email"`
	Password   string `xorm:"not null default '' CHAR(32)" form:"password" json:"-"`
	Avatar     string `xorm:"not null default '' comment('头像地址') VARCHAR(355)" form:"avatar" json:"avatar"`
	IsAdmin    int    `xorm:"not null default 0 comment('0普通用户1管理员') index TINYINT(4)" json:"-"`
	CreateTime int    `xorm:"not null default 0 INT(10)" json:"create_time"`
}
