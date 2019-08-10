package repositories

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/datasource/mysql_conn"
	"github.com/pkg/errors"
)

// 添加导航
func NavAdd(nav *datamodels.DNav) error {
	i, e := mysql_conn.MysqlEngine.InsertOne(nav)
	if e != nil || i== 0 {
		return errors.New("insert nav error")
	}
	return nil
}

// 修改导航
func NavModify(nav *datamodels.DNav) error {
	i, e := mysql_conn.MysqlEngine.Where("id = ?", nav.Id).Update(nav)
	if e != nil || i == 0 {
		return errors.New("update nav error")
	}
	return nil
}

// 删除导航
func NavDel(id int) error {
	i, e := mysql_conn.MysqlEngine.Where("id = ?", id).Delete(datamodels.DNav{})
	if e != nil || i == 0 {
		return errors.New("del nav error")
	}
	return nil
}

// Id查询导航
func NavSelectById(id int) (*datamodels.DNav,error) {
	data := &datamodels.DNav{}
	b, e := mysql_conn.MysqlEngine.Where("id = ?", id).Get(data)
	if e != nil || b == false {
		return nil,errors.New("not data")
	}

	return data,nil
}
