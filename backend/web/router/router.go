package router

import (
	"Vue-Admin/backend/web/controller"
	"Vue-Admin/backend/web/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	app := gin.New()

	router(app)

	return app
}

func router(app *gin.Engine) {

	// 这样都是 api/v1 分为 需要 和不需要 权鉴

	// 普通 v1
	v1 := app.Group("api/v1")
	{
		// 用户
		user := v1.Group("/user")
		{
			// 用户注册
			user.POST("/register", controller.UserRegister)
			// 用户登录 返回token
			user.POST("/login", controller.UserLogin)
		}
	}

	// v1 需要 权鉴
	v1_auth := app.Group("api/v1",middleware.Auth)
	{
		// 用户
		user := v1_auth.Group("/user")
		{
			user.GET("/info/:email", controller.UserInfo)
		}

		// 导航管理
		nav := v1_auth.Group("/nav")
		{
			// 添加导航
			nav.POST("/add",controller.NavAdd)
			// 修改导航
			nav.POST("/modify",controller.NavModify)
			// 删除导航
			nav.GET("/del/:id",controller.NavDel)
			// 通过id查询
			nav.GET("/data/:id",controller.NavGetInfoById)
		}

		// 分类管理
		sort := v1_auth.Group("/sort")
		{
			// 添加分类
			sort.POST("/add",controller.SortAdd)
			// 修改分类
			sort.POST("/modify",controller.SortModify)
			// 删除分类
			sort.GET("/del/:id",controller.SortDel)
			// 通过id查询
			sort.GET("/data/:id",controller.SortGetInfoById)
		}

		// carousel 轮播图管理
		carousel := v1_auth.Group("/carousel")
		{
			// 添加轮播图
			carousel.POST("/add",controller.CarouselAdd)
			// 修改轮播图
			carousel.POST("/modify",controller.CarouselModify)
			// 删除分类
			carousel.GET("/del/:id",controller.CartouselDelById)
			// 通过id查询
			carousel.GET("/data/:id",controller.CarouselGetInfoById)
		}


		// article 管理
		article := v1_auth.Group("/article")
		{
			// 添加 article
			article.POST("/add",controller.ArticleAdd)
			// 修改 article
			article.POST("/modify",controller.ArticleModify)
			// 删除 article
			article.GET("/del/:id",controller.ArticleDel)
			// 通过id查询
			article.GET("/data/:id",controller.ArticleGetInfoById)
		}

		// article item 管理
		article_item := v1_auth.Group("/article_item")
		{
			// 添加 article item
			article_item.POST("/add",controller.ArticleItemAdd)
			// 修改 article item
			article_item.POST("/modify",controller.ArticleItemModify)
			// 删除 article item
			article_item.GET("/del/:id",controller.ArticleItemDel)
			// 通过id查询
			article_item.GET("/data/:id",controller.ArticleItemGetInfoById)
		}
	}


}
