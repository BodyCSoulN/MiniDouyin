package router

import "github.com/MiniDouyin/controller"

func PublishRouter() {
	// GET
	DouyinRouter.GET("/publish/list/", controller.GetPublishList)
	// POST
	DouyinRouter.POST("/publish/action/", controller.PublishVideo)
}
