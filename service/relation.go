package service

import (
	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/storage"
)

// Action 关注操作
func Action(user_id, to_user_id int64, action_type string) error {
	return model.Action(user_id, to_user_id, action_type)
}

// FollowList 关注列表
func FollowList(user_id int64) (follow_list []storage.User, err error) {
	idList, err := model.FollowList(user_id)
	if err != nil {
		return
	}
	for _, s := range *idList {
		temp, err := model.GetUserInfoByIDR(s.BeWatchID)
		if err != nil {
			continue
		}
		follow_list = append(follow_list, temp)
	}
	return
}

// FollowerList 粉丝列表
func FollowerList(user_id int64) (follower_list []storage.User, err error) {
	idList, err := model.FollowerList(user_id)
	if err != nil {
		return
	}
	for _, s := range *idList {
		temp, err := model.GetUserInfoByIDR(s.WatchID)
		if err != nil {
			continue
		}
		follower_list = append(follower_list, temp)
	}
	return
}
