package model

import (
	"time"

	"github.com/MiniDouyin/storage"
	"github.com/MiniDouyin/utils"
)

// Feed 视频流数据口接口 单词最大视频数为30
func Feed(latest_time time.Time) (video_list *[]storage.Video, err error) {
	res := storage.Mysql.Table("video").Where("created_at <= ?", utils.RParseDateAndTime(latest_time)).Order("created_at desc").Limit(30).Find(&video_list)
	if res.Error != nil {
		return nil, res.Error
	}
	return
}
