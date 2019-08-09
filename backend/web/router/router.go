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
	v1 := app.Group("api/v1")
	{
		// 用户 不需要验证
		user := v1.Group("/user")
		{
			// 用户注册
			user.POST("/register", controller.UserRegister)
			// 用户登录 返回token
			user.POST("/login", controller.UserLogin)
		}


		// 用户 需要验证
		user_auth := v1.Group("/user", middleware.Auth)
		{
			user_auth.GET("/info/:email", controller.UserInfo)
		}

	}
}
