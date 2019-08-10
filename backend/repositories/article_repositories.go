package repositories

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/datasource/mysql_conn"
	"Vue-Admin/backend/easyutils"
	"github.com/pkg/errors"
)

// 添加 article
func ArticleAdd(article *datamodels.DArticle) error {
	article.CreateTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.Insert(article)
	if e != nil || i == 0 {
		return errors.New("insert article error")
	}
	return nil
}


// 修改 article
func ArticleModify(article *datamodels.DArticle) error {
	article.UpdateTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.Where("id = ?", article.Id).Update(article)
	if e != nil || i == 0 {
		return errors.New("update nav error")
	}
	return nil
}

// 删除 article
func ArticleDel(id int) error{
	data := datamodels.DArticle{}
	data.DeleteTime = easyutils.TimeGetNowTime()
	i, e := mysql_conn.MysqlEngine.Where("id = ?", id).Cols("update_time").Update(&data)
	if e != nil || i == 0 {
		return errors.New("del nav error")
	}
	return nil
}

// 更具id查询 article
func ArticleGetById(id int) (*datamodels.DArticle,error) {
	data := &datamodels.DArticle{}
	b, e := mysql_conn.MysqlEngine.Where("id = ?", id).Where("delete_time = ?", 0).Get(data)
	if e != nil || b == false {
		return nil,errors.New("article not data")
	}
	return data,nil
}


