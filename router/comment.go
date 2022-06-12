package router

import "github.com/MiniDouyin/controller"

func CommentRouter() {
	// GET
	DouyinRouter.GET("/comment/comment/", controller.GetComment)
	DouyinRouter.GET("/comment/count/", controller.GetCommentFavoriteCount)
	DouyinRouter.GET("/comment/front_list/", controller.GetCommentListFront)
	DouyinRouter.GET("/comment/list", controller.GetCommentList)
	// POST
	DouyinRouter.POST("/comment/delete/", controller.DeleteComment)
	DouyinRouter.POST("/comment/add/", controller.AddComment)
}
