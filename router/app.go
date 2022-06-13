package router

import (
	"github.com/gin-gonic/gin"
)

var DouyinRouter *gin.RouterGroup

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")

	DouyinRouter = r.Group("/douyin")

	FeedRouter()
	FavoriteRouter()
	RelationRouter()
	UserRouter()
	CommentRouter()
	PublishRouter()

	return r

}
