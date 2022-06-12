package router

import "github.com/MiniDouyin/controller"

func FeedRouter() {
	DouyinRouter.GET("/feed/", controller.Feed)
}
