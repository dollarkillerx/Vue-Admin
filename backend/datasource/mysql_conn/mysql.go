package mysql_conn

import (
	"Vue-Admin/backend/config"
	"Vue-Admin/backend/datamodels"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

var (
	MysqlEngine *xorm.Engine
	e           error
)

func init() {
	MysqlEngine, e = xorm.NewEngine("mysql", config.MyConfig.Mysql.Dsn)
	if e != nil {
		panic(e.Error())
	}

	if config.MyConfig.App.Debug {
		MysqlEngine.ShowSQL(true)
		log.Println(config.MyConfig.Mysql.Dsn)
	}

	ping := MysqlEngine.Ping()
	if ping != nil {
		panic(ping.Error())
	}

	if config.MyConfig.Mysql.Cache {
		cacher2 := xorm.NewLRUCacher2(xorm.NewMemoryStore(), time.Hour*4, 1024)
		MysqlEngine.SetDefaultCacher(cacher2)
	}

	// 数据库表映射
	mapping()
}

func mapping() {
	sync2 := MysqlEngine.Sync2(
		new(datamodels.DUser), // 用户
		new(datamodels.DArticleItem), // 子文章
		new(datamodels.DArticle), // 文章
		new(datamodels.DNav), // 导航
		new(datamodels.DCarousel), // 轮播图
		new(datamodels.DSort), // 分类
	)

	if sync2 != nil {
		panic(sync2.Error())
	}
}
