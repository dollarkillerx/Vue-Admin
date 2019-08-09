package repositories

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/datasource/mysql_conn"
	"Vue-Admin/backend/easyutils"
	"errors"
)

// 用户注册
func UserRegister(user *datamodels.DUser) error {
	user.CreateTime = easyutils.TimeGetNowTime()
	user.Password = easyutils.Md5Encode(user.Password)
	i, e := mysql_conn.MysqlEngine.InsertOne(user)
	if e != nil || i == 0 {
		return errors.New("insert data error")
	}
	return nil
}

// 用户身份验证
func UserCheck(user *datamodels.DUser) error {
	// 查询用户是否存在
	c_user := datamodels.DUser{}
	b, e := mysql_conn.MysqlEngine.Where("email = ?", user.Email).Get(&c_user)
	if e != nil || b == false {
		return errors.New("not data")
	}

	// 用户存在 验证用户密码是否正确
	psw := easyutils.Md5Encode(user.Password)
	if psw == c_user.Password {
		return nil
	}
	return errors.New("password error")
}

// 更具email 获取用户信息
func UserGetInfoByEmail(email string) (*datamodels.DUser, error) {
	data := datamodels.DUser{}
	b, e := mysql_conn.MysqlEngine.Where("email = ?", email).Get(&data)
	if e != nil || b == false {
		return nil, errors.New("not data")
	}
	return &data, nil
}
