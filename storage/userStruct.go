package storage

import (
	"gorm.io/gorm"
	"time"
)

type UsrToken struct {
	Token   string
	EndTime time.Time
}

// TokenValidTime token存活时间
var TokenValidTime = 7 * 24 * time.Hour

// TokenEndTime
// 用户名映射{token,过期时间}
var TokenEndTime = map[string]UsrToken{
	"admin": {
		Token:   "",
		EndTime: time.Now().Add(TokenValidTime),
	},
}

// UsersLoginInfo  use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var UsersLoginInfo = map[string]User{
	//"zhangleidouyin": {
	//	Id:            1,
	//	Name:          "zhanglei",
	//	FollowCount:   10,
	//	FollowerCount: 5,
	//	IsFollow:      true,
	//},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// User 用户信息
type User struct {
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	gorm.Model
}
