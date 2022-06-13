package controller

import (
	"net/http"
	"strconv"

	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

//	AddComment 新增评论
func AddComment(c *gin.Context) {
	var data storage.CommentActionRequest
	_ = c.ShouldBindJSON(&data)
	claim, _ := model.Getting(data.Token)
	toData := &storage.Comment{
		Id: claim.Id,
	}
	code := model.AddComment(toData)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": model.GetErrMsg(code),
	})
}

//	GetComment 获取评论信息
func GetComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetComment(int64(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": model.GetErrMsg(code),
	})
}

//	GetCommentFavoriteCount 获取评论点赞数
func GetCommentFavoriteCount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	total := model.GetCommentFavoriteCount(int64(id))
	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

//	DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteComment(int64(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": model.GetErrMsg(code),
	})
}

// GetCommentList 后台查询评论列表
func GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetCommentList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": model.GetErrMsg(code),
	})

}

//	GetCommentListFront 展示页面显示评论列表
func GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetCommentListFront(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": model.GetErrMsg(code),
	})
}
