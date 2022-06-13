package storage

import (
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
	Id            int64     `json:"id,omitempty"`
	AuthorID      int64     `json:"-"`
	Author        User      `gorm:"foreignKey:AuthorID;references:ID;"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type Comment struct {
	Id         int64 `json:"id,omitempty"`
	UserID     int64
	User       User   `json:"user" gorm:"foreignKey:UserID;references:ID;"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []VideoResponse `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

//	User 用户信息
type User struct {
	ID            int64  `json:"id,omitempty"`
	UserName      string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// 关注
type RelationResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"`
}

type VideoResponse struct {
	Id            int64  `json:"id,omitempty"`
	AuthorID      int64  `json:"-"`
	Author        User   `gorm:"foreignKey:AuthorID;references:ID;"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
