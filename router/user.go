package router

import "github.com/MiniDouyin/controller"

func UserRouter() {
	// GET
	DouyinRouter.GET("/user/", controller.UserInfo)
	// POST
	DouyinRouter.POST("/user/register/", controller.Register)
	// Login功能存在bug 一旦后端服务重启，已注册账号将无法登录
	DouyinRouter.POST("/user/login/", controller.Login)
}
