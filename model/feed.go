package model

import (
	"time"

	"github.com/MiniDouyin/storage"
)

type Video struct {
	storage.Video
}

func (v Video) TableName() string {
	return "video"
}

func Feed(latest_time time.Time) (video_list *[]Video, err error) {
	return nil, nil
}
