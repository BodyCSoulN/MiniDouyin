package model

import "github.com/MiniDouyin/storage"

func GetVideoByID(video_id int64) (video storage.VideoResponse, err error) {
	err = storage.Mysql.Table("video").Where("id = ?", video_id).Find(&video).Error
	return
}
