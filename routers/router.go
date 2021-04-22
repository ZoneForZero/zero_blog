package routers

import (
	"os"
	CONTROLLER "zero_blog/controller"
	MIDDLEWARE "zero_blog/middleware"

	"github.com/gin-gonic/gin"
)

// 简单应用，默认一个路由
func AllRouter() *gin.Engine {
	// 创建路由
	router := gin.Default()
	// 加载中间件, 顺序不能改
	router.Use(MIDDLEWARE.Session(os.Getenv("SESSION_SECRET")))
	router.Use(MIDDLEWARE.Cors())
	router.Use(MIDDLEWARE.CurrentUser())
	// 路由分组，testRouter的路由前面默认为/testRouter/
	testRouter := router.Group("/testRouter")
	{
		testRouter.GET("ping", CONTROLLER.Ping)
		// 用户注册
		testRouter.POST("user/register", CONTROLLER.UserRegister)
		// 用户登录
		testRouter.POST("user/login", CONTROLLER.UserLogin)
		// 需要登录保护的
		auth := testRouter.Group("")
		auth.Use(MIDDLEWARE.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", CONTROLLER.UserMe)
			auth.DELETE("user/logout", CONTROLLER.UserLogout)
		}
	}
	return router
}
