package controller

import (
	"github.com/MiniDouyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPublishList(c *gin.Context) {
	token := c.Query("token")
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, returnFormat(-1, "Parse UserID Error:"+err.Error()))
		return
	}

	if token == "" || userID <= 0 {
		c.JSON(http.StatusOK, returnFormat(-1, "Unsupported Data Format"))
		return
	}

	// TODO: parse token

	videoList, err := service.GetPublishList(userID)

	if err != nil {
		c.JSON(http.StatusOK, returnFormat(-1, "Get Published Video List Error:"+err.Error()))
		return
	}

	tmp := returnFormat(0, "")

	tmp["videoList"] = videoList

	c.JSON(http.StatusOK, tmp)

}

func PublishVideo(c *gin.Context) {
	video, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, returnFormat(-1, "Can't Read Video File"))
		return
	}
	token := c.PostForm("token")
	title := c.PostForm("title")

	if token == "" || title == "" {
		c.JSON(http.StatusOK, returnFormat(-1, "Unsupported Data Format"))
		return
	}
	// TODO: parse token to get userID
	//userID, err := utils.ParseToken(token)
	//if err != nil {
	//	c.JSON(http.StatusOK, returnFormat(-1, "parse Token Error: "+err.Error()))
	//	return
	//}

	var userID int64 = 1

	if err := service.PublishVideo(video, title, userID); err != nil {
		c.JSON(http.StatusOK, returnFormat(-1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, returnFormat(0, "Published Successfully"))

}

func returnFormat(code int, msg string) map[string]interface{} {
	return map[string]interface{}{
		"status_code": code,
		"status_msg":  msg,
	}
}
