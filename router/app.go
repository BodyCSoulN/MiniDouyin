package router

import (
	"github.com/MiniDouyin/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "E:\\codelife\\Goland_Project\\MiniDouyin\\static")

	douyinRouter := r.Group("/douyin")

	douyinRouter.POST("/publish/action", controller.PublishVideo)
	douyinRouter.GET("/publish/list", controller.GetPublishList)

	douyinRouter.GET("/comment/comment", controller.GetComment)
	douyinRouter.GET("/comment/count", controller.GetCommentFavoriteCount)
	douyinRouter.GET("/comment/front_list", controller.GetCommentListFront)
	douyinRouter.GET("/comment/list", controller.GetCommentList)
	douyinRouter.POST("/comment/delete", controller.DeleteComment)
	douyinRouter.POST("/comment/add", controller.AddComment)

	return r

}
