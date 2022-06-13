package service

import (
	"time"

	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/storage"
)

// Feed 视频流接口 带next_time
func Feed(user_id int64, latest_time time.Time) (feed_resp *storage.FeedResponse, err error) {
	video_list, err := model.Feed(latest_time)
	if err != nil {
		return nil, err
	}
	video_list_n := make([]storage.VideoResponse, 0)
	for _, video := range *video_list {
		author, err := model.GetUserInfoByIDR(video.AuthorID)
		if err != nil {
			continue
		}
		video_temp := storage.VideoResponse{
			Id:            video.Id,
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
	var next_time time.Time
	video_list_len := len(*video_list)
	if video_list_len == 0 {
		return &storage.FeedResponse{
			Response: storage.Response{
				StatusCode: -1,
				StatusMsg:  "视频已经被你刷完啦",
			},
		}, nil
	}
	last_video := (*video_list)[video_list_len-1]
	next_time = last_video.CreatedAt
	return &storage.FeedResponse{
		Response: storage.Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		VideoList: video_list_n,
		NextTime:  next_time.Unix(),
	}, nil
}
