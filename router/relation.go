package router

import "github.com/MiniDouyin/controller"

func RelationRouter() {
	// GET
	DouyinRouter.GET("/relation/follow/list/", controller.FollowList)
	DouyinRouter.GET("/relation/follower/list/", controller.FollowerList)
	// POST
	DouyinRouter.POST("/relation/action/", controller.Action)
}
