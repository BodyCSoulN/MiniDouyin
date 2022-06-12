package router

import (
	"github.com/MiniDouyin/controller"
)

func PublishRouter() {
	publishRouter := Router.Group("/douyin")

	// POST
	publishRouter.POST("/publish/action", controller.PublishVideo)
	// GET
	publishRouter.GET("/publish/list", controller.GetPublishList)
}
