package router

import "github.com/MiniDouyin/controller"

func FavoriteRouter() {
	// GET
	DouyinRouter.GET("/favorite/list/", controller.FavoriteList)
	// POST
	DouyinRouter.POST("/favorite/action/", controller.FavoriteAction)
}
