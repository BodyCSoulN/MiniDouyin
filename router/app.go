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

	return r

}
