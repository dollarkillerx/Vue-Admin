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
			nav.POST("/del/:id",controller.NavDel)
			// 通过id查询
			nav.GET("/:id",controller.NavGetInfoById)
		}

		// 分类管理
		sort := v1_auth.Group("/sort")
		{
			// 添加分类
			sort.POST("/add",controller.SortAdd)
			// 修改分类
			sort.POST("/modify",controller.SortModify)
			// 删除分类
			sort.POST("/del/:id",controller.SortDel)
			// 更过id查询
			sort.GET("/:id",controller.SortGetInfoById)
		}

	}


}
