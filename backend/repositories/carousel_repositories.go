package repositories

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/datasource/mysql_conn"
	"Vue-Admin/backend/easyutils"
	"github.com/pkg/errors"
)

// 添加轮播图
func CarouselAdd(data *datamodels.DCarousel) error {
	data.CreateTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.InsertOne(data)
	if e != nil || i == 0 {
		return errors.New("carousel add data error")
	}

	return nil
}

// 修改轮播图
func CarouselModify(data *datamodels.DCarousel) error {
	data.UpdateTime = easyutils.TimeGetNowTime()

	i, e := mysql_conn.MysqlEngine.Where("id = ?", data.Id).Update(data)
	if e != nil || i == 0 {
		return errors.New("update carousel error")
	}
	return nil
}

// 删除轮播图
func CarouselDel(id int) error {
	data := datamodels.DCarousel{}
	data.DeleteTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.Where("id = ?", id).Cols("delete_time").Update(&data)
	if e != nil || i == 0 {
		return errors.New("carousel del error")
	}
	return nil
}

// 更具id 查询
func CarouseSelectById(id int) (*datamodels.DCarousel,error) {
	data := datamodels.DCarousel{}
	b, e := mysql_conn.MysqlEngine.Where("id = ?", id).Where("delete_time = ?", 0).Get(&data)
	if e != nil || b == false {
		return nil,errors.New("not data")

	}
	return &data,nil
}

