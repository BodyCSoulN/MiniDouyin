package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/service"
	"github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	var lastest_time time.Time
	latest_time_q := c.Query("latest_time")
	if latest_time_q == "" {
		lastest_time = time.Now()
	} else {
		lastest_time_t, err := strconv.ParseInt(latest_time_q, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, storage.Response{
				StatusCode: -1,
				StatusMsg:  "请求时间戳有误",
			})
			return
		}
		lastest_time = time.Unix(lastest_time_t/1000, 0)
	}
	token := c.Query("token")
	var user_id int64
	claims, err := model.Getting(token)
	if err == nil {
		user_id = claims.Id
	}
	feed_resp, err := service.Feed(user_id, lastest_time)
	if err != nil {
		c.JSON(http.StatusOK, storage.Response{
			StatusCode: -1,
			StatusMsg:  "服务器错误",
		})
	}
	c.JSON(http.StatusOK, feed_resp)
}
