package model

import (
	"errors"
	"log"

	"github.com/MiniDouyin/storage"
	"gorm.io/gorm"
)

func changeFavoriteCount(video_id int64, swap_direction bool) {
	log.Println("video_id", video_id, "swap_direction", swap_direction)
	if swap_direction {
		storage.Mysql.Model(&storage.Video{}).Where("id = ?", video_id).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
	} else {
		storage.Mysql.Model(&storage.Video{}).Where("id = ?", video_id).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
	}
}

func FavoriteAciton(user_id, video_id int64, action_type string) error {
	var count int64
	storage.Mysql.Model(&storage.Favorite{}).Where("user_id = ? and video_id = ?", user_id, video_id).Count(&count)
	if count > 0 && action_type == "1" {
		// 点赞但是还想点赞
		return errors.New("already favorite")
	} else if count > 0 && action_type == "2" {
		// 取消点赞
		changeFavoriteCount(video_id, false)
		return storage.Mysql.Table("favorite").Where("user_id = ? and video_id = ?", user_id, video_id).Delete(&storage.Favorite{}).Error
	} else if count == 0 && action_type == "2" {
		return errors.New("没点赞就想取消点赞啊")
	}
	// 正常点赞
	toInsert := &storage.Favorite{
		UserID:  user_id,
		VideoID: video_id,
	}
	changeFavoriteCount(video_id, true)
	return storage.Mysql.Model(&storage.Favorite{}).Create(toInsert).Error
}

func IsFavorite(user_id, video_id int64) bool {
	var count int64
	storage.Mysql.Table("favorite").Where("user_id = ? and video_id = ?", user_id, video_id).Count(&count)
	return count > 0
}

func FavoriteList(user_id int64) (favorite_list *[]storage.Favorite, err error) {
	err = storage.Mysql.Table("favorite").Where("user_id = ?", user_id).Find(&favorite_list).Error
	return
}
