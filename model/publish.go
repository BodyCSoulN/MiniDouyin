package model

import (
	"github.com/MiniDouyin/storage"
	"gorm.io/gorm"
)

type VideoResponse struct {
	ID            int64
	AuthorID      int64
	Author        UserResponse `gorm:"foreignKey:AuthorID;references:ID;"`
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
}

type UserResponse struct {
	ID            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

func GetVideoList(UserID int64) *gorm.DB {
	return storage.Mysql.Model(new(storage.Video)).Preload("Author").
		Where("author_id = ?", UserID)
}

func PublishVideo(videoUrl, coverUrl, title string, userID int64) error {
	video := &storage.Video{
		AuthorID:      userID,
		PlayUrl:       videoUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}
	return storage.Mysql.Create(video).Error
}
