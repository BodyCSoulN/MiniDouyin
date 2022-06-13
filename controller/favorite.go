package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/service"
	"github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

func FavoriteList(c *gin.Context) {
	user_id_q := c.Query("user_id")
	user_id, err := strconv.ParseInt(user_id_q, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "user_id error",
		})
		return
	}
	token := c.Query("token")
	_, err = model.Getting(token)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "token error",
		})
		return
	}
	video_list, err := service.FavoriteList(user_id)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "request error",
		})
		return
	}
	c.JSON(http.StatusOK, storage.FavoriteResponse{
		Response: storage.Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		VideoList: video_list,
	})
}

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Error Token",
		})
		return
	}
	var user_id int64
	claims, err := model.Getting(token)
	if err == nil {
		user_id = claims.Id
	}
	video_id_q := c.Query("video_id")
	video_id, err := strconv.ParseInt(video_id_q, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "video_id error",
		})
	}
	action_type := c.Query("action_type")
	if err = service.FavoriteAction(user_id, video_id, action_type); err == nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
	} else {
		log.Println(err)
	}
}
