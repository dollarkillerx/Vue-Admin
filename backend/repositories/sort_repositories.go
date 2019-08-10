package repositories

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/datasource/mysql_conn"
	"Vue-Admin/backend/easyutils"
	"github.com/pkg/errors"
)

// 添加分类
func SortAdd(sort *datamodels.DSort) error {
	sort.CreateTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.InsertOne(&sort)
	if e != nil || i == 0 {
		return errors.New("add sort error")
	}
	return nil
}

// 修改分类
func SortModify(sort *datamodels.DSort) error {
	sort.UpdateTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.Where("id = ?", sort.Id).Update(sort)
	if e != nil || i == 0 {
		return errors.New("modify sort error")
	}
	return nil
}

// 更具id查询分类
func SortSelectById(id int) (*datamodels.DSort,error) {
	data := datamodels.DSort{}
	b, e := mysql_conn.MysqlEngine.Where("id = ?", id).Where("delete_time = ?",0).Get(&data)
	if e != nil || b == false {
		return nil,errors.New("sort not data")
	}
	return &data,nil
}

// 删除分类
func SortDelById(id int) error {
	data := datamodels.DSort{}
	data.DeleteTime = easyutils.TimeGetNowTime()

	i, e := mysql_conn.MysqlEngine.Where("id = ?", id).Cols("delete_time").Update(&data)
	if e != nil || i == 0 {
		return errors.New("del data error")
	}
	return nil
}