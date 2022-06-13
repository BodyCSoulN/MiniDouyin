package service

import (
	"log"

	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/storage"
)

func FavoriteList(user_id int64) (video_list []storage.VideoResponse, err error) {
	log.Println(user_id)
	idList, err := model.FavoriteList(user_id)
	log.Printf("%#v\n", idList)
	if err != nil {
		return
	}
	video_list_n := make([]storage.VideoResponse, 0)
	for _, s := range *idList {
		author, err := model.GetUserInfoByIDR(s.UserID)
		video, err2 := model.GetVideoByID(s.VideoID)
		if err != nil || err2 != nil {
			continue
		}
		video_temp := storage.VideoResponse{
			Id:            s.VideoID,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    model.IsFavorite(user_id, video.Id),
			Title:         video.Title,
		}
		video_list_n = append(video_list_n, video_temp)
	}
	return video_list_n, nil
}

func FavoriteAction(user_id, video_id int64, action_type string) error {
	return model.FavoriteAciton(user_id, video_id, action_type)
}
