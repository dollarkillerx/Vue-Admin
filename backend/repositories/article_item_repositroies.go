package repositories

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/datasource/mysql_conn"
	"Vue-Admin/backend/easyutils"
	"errors"
)

// 添加
func ArticleItemAdd(data *datamodels.DArticleItem) error {
	data.CreateTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.InsertOne(data)
	if e != nil || i == 0 {
		return errors.New("article add data error")
	}
	return nil
}

// 修改
func ArticleItemModify(data *datamodels.DArticleItem) error {
	data.UpdateTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.Where("id = ?", data.Id).Update(data)
	if e != nil || i == 0 {
		return errors.New("modify data error")
	}
	return nil
}

// 删除
func ArticleItemDel(id int) error {
	data := datamodels.DArticleItem{}
	data.DeleteTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.Where("id = ?", id).Cols("update_time").Update(&data)
	if e != nil || i == 0 {
		return errors.New("del nav error")
	}
	return nil
}

// 更具id查询
func ArticleItemSelectById(id int) (*datamodels.DArticleItem,error) {
	data := &datamodels.DArticleItem{}
	b, e := mysql_conn.MysqlEngine.Where("id = ?", id).Where("delete_time = ?", 0).Get(data)
	if e != nil || b == false {
		return nil,errors.New("article not data")
	}
	return data,nil
}