package model

import (
	"github.com/MiniDouyin/storage"
	"log"
)

type CommentListResponse struct {
	CommentList []storage.Comment `json:"comment_list"`
}

//  定义常量状态码
const (
	SUCCESS = 200
	ERROR   = 500
)

//	状态码对应的信息
var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",
}

//	GetErrMsg 查询状态函数
func GetErrMsg(code int) string {
	return codeMsg[code]
}

//	AddComment 增加评论
func AddComment(data *storage.Comment) int {
	err := storage.Mysql.Create(&data).Error
	if err != err {
		log.Println(err)
		return ERROR
	}
	return SUCCESS
}

//	GetComment 查询单个评论
func GetComment(id int64) (storage.Comment, int) {
	var comment storage.Comment
	err := storage.Mysql.Where("id = ?", id).First(&comment).Error
	if err != nil {
		log.Println(err)
		return comment, ERROR
	}
	return comment, SUCCESS
}

func GetCommentFavoriteCount(id int64) int64 {
	var comment storage.Comment
	var total int64
	storage.Mysql.Find(&comment).Where("favorite_count = ?", id).Count(&total)
	return total
}

// DeleteComment 删除评论
func DeleteComment(id int64) int {
	var comment storage.Comment
	err := storage.Mysql.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		log.Println(err)
		return ERROR
	}
	return SUCCESS
}

//	GetCommentList 后台查询评论列表
func GetCommentList(pageSize int, pageNum int) ([]storage.Comment, int64, int) {
	var commentList []storage.Comment
	var total int64
	storage.Mysql.Find(&commentList).Count(&total)
	err := storage.Mysql.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("Create_At DESC").
		Select("comment.id, user_id, video_id, content, favorite_count, comment.created_at, comment.update_at, comment.deleted_at")
	if err != nil {
		log.Println(err)
		return commentList, 0, ERROR
	}
	return commentList, total, SUCCESS
}

// GetCommentListFront 展示页面获取评论列表
func GetCommentListFront(id int, pageSize int, pageNum int) ([]storage.Comment, int64, int) {
	var commentList []storage.Comment
	var total int64
	storage.Mysql.Find(&storage.Comment{}).Where("id = ?", id).Count(&total)
	err := storage.Mysql.Model(&storage.Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").
		Select("comment.id, video_id, user_id, user.name, comment.content, comment.favorite_count, comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN video ON comment.video_id = video.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").
		Where("video_id = ?", id).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, ERROR
	}
	return commentList, total, SUCCESS
}
