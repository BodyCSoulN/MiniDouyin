package router

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() *gin.Engine {
	r := gin.Default()
	Router = r

	r.Static("/static", "../static")

	PublishRouter()

	return r
}
