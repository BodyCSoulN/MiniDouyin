package model

import (
	"errors"

	"github.com/MiniDouyin/storage"
	"gorm.io/gorm"
)

func changeFollowCount(user_id, to_user_id int64, swap_direction bool) {
	if swap_direction {
		storage.Mysql.Model(&storage.DBUser{}).Where("id = ?", user_id).
			UpdateColumn("followCount", gorm.Expr("followCount + ?", 1))
		storage.Mysql.Model(&storage.DBUser{}).Where("id = ?", to_user_id).
			UpdateColumn("followerCount", gorm.Expr("followerCount + ?", 1))
	} else {
		storage.Mysql.Model(&storage.DBUser{}).Where("id = ?", user_id).
			UpdateColumn("followCount", gorm.Expr("followCount - ?", 1))
		storage.Mysql.Model(&storage.DBUser{}).Where("id = ?", to_user_id).
			UpdateColumn("followerCount", gorm.Expr("followerCount - ?", 1))
	}
}

// Action 关注操作
func Action(user_id, to_user_id int64, action_type string) error {
	var count int64
	storage.Mysql.Model(&storage.Attention{}).Where("watch_id = ? and be_watch_id = ?", user_id, to_user_id).Count(&count)
	if count > 0 && action_type == "1" {
		// 关注但是还想关注
		return errors.New("already pay attention")
	} else if count > 0 && action_type == "2" {
		// 取关
		changeFollowCount(user_id, to_user_id, false)
		return storage.Mysql.Model(&storage.Attention{}).Where("watch_id = ? and be_watch_id = ?", user_id, to_user_id).Delete(&storage.Attention{}).Error
	} else if count == 0 && action_type == "2" {
		// 没关注但是想取关
		return errors.New("you are no pay attention")
	}
	// 正常关注
	toInsert := &storage.Attention{
		WatchID:   user_id,
		BeWatchID: to_user_id,
	}
	changeFollowCount(user_id, to_user_id, true)
	return storage.Mysql.Model(&storage.Attention{}).Create(toInsert).Error
}

func FollowList(user_id int64) (follow_list *[]storage.Attention, err error) {
	err = storage.Mysql.Table("attention").Where("be_watch_id = ?", user_id).Find(&follow_list).Error
	return
}

func FollowerList(user_id int64) (follower_list *[]storage.Attention, err error) {
	err = storage.Mysql.Table("attention").Where("watch_id = ?", user_id).Find(&follower_list).Error
	return
}
