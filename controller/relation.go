package controller

import (
	"net/http"
	"strconv"

	"github.com/MiniDouyin/service"
	"github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {
	token := c.Query("token")
	to_user_id_q := c.Query("to_user_id")
	to_user_id, _ := strconv.ParseInt(to_user_id_q, 10, 64)
	action_type := c.Query("action_type")

	DYUser := storage.UsersLoginInfo[token]
	err := service.Action(DYUser.ID, to_user_id, action_type)

	if err != nil {
		if action_type == "1" {
			c.JSON(http.StatusOK, storage.Response{
				StatusCode: -1,
				StatusMsg:  "关注失败",
			})
			return
		}
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "取关失败",
		})
		return
	}
	c.JSON(http.StatusOK, storage.Response{
		StatusCode: 0,
		StatusMsg:  "成功",
	})
}

func FollowList(c *gin.Context) {
	user_id_q := c.Query("user_id")
	user_id, _ := strconv.ParseInt(user_id_q, 10, 64)
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Error Token",
		})
		return
	}
	follow_list, err := service.FollowList(user_id)
	if err != nil {
		c.JSON(http.StatusOK, storage.RelationResponse{
			Response: storage.Response{StatusCode: 0, StatusMsg: "获取失败"},
		})
		return
	}
	c.JSON(http.StatusOK, storage.RelationResponse{
		Response: storage.Response{StatusCode: 0, StatusMsg: "获取成功"},
		UserList: follow_list,
	})
}

func FollowerList(c *gin.Context) {
	user_id_q := c.Query("user_id")
	user_id, _ := strconv.ParseInt(user_id_q, 10, 64)
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Error Token",
		})
		return
	}
	follower_list, err := service.FollowerList(user_id)
	if err != nil {
		c.JSON(http.StatusOK, storage.RelationResponse{
			Response: storage.Response{StatusCode: 0, StatusMsg: "获取失败"},
		})
		return
	}
	c.JSON(http.StatusOK, storage.RelationResponse{
		Response: storage.Response{StatusCode: 0, StatusMsg: "获取成功"},
		UserList: follower_list,
	})
}
