package controller

import (
	"net/http"
	"strconv"

	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/service"
	"github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

// GetPublishList 获取发布列表
func GetPublishList(c *gin.Context) {
	token := c.Query("token")
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Parse UserID Error:" + err.Error(),
		})
		return
	}

	if token == "" || userID <= 0 {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Unsupported Data Format",
		})
		return
	}

	_, err = model.Getting(token)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Parse Token Error: " + err.Error(),
		})
		return
	}

	videoList, err := service.GetPublishList(userID)

	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Get Published Video List Error: " + err.Error(),
		})
		return
	}

	tmp := returnFormat(0, "")

	tmp["video_list"] = videoList

	c.JSON(http.StatusOK, tmp)

}

// PublishVideo 发布视频
func PublishVideo(c *gin.Context) {
	video, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Can't Read Video File",
		})
		return
	}
	token := c.PostForm("token")
	title := c.PostForm("title")

	if token == "" || title == "" {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Unsupported Data Format",
		})
		return
	}

	claims, err := model.Getting(token)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Parse Token Error: " + err.Error(),
		})
		return
	}

	if err := service.PublishVideo(video, title, claims.Id); err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "Publish Video Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, storage.Response{
		StatusCode: 0,
		StatusMsg:  "Published Successfully",
	})

}

func returnFormat(code int, msg string) map[string]interface{} {
	return map[string]interface{}{
		"status_code": code,
		"status_msg":  msg,
	}
}
