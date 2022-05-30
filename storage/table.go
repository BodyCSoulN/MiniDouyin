package storage

import "gorm.io/gorm"

// 用户表(作者)
type User struct {
	ID            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool `json:"is_follow,omitempty"`
	gorm.Model
}

// 视频表
type Video struct {
	ID            int64
	AuthorID      int64 `json:"author_id,omitempty"`
	Author        User  `gorm:"foreignKey:AuthorID;references:ID;"`
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
	gorm.Model
}

// 评论表
type Comment struct {
	ID      int64
	UserID  int64
	User    User `gorm:"foreignKey:UserID;references:ID;"`
	VideoID int64
	Video   Video `gorm:"foreignKey:VideoID;references:ID"`
	Content string
	gorm.Model
}

// 关注表
type Attention struct {
	ID        int64
	WatchID   int64
	Watch     User `gorm:"foreignKey:WatchID;references:ID;"`
	BeWatchID int64
	BeWatch   User `gorm:"foreignKey:BeWatchID;references:ID;"`
}
