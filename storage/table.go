package storage

// 用户表(作者)
//type User struct {
//	ID            int64
//	Name          string
//	FollowCount   int64
//	FollowerCount int64
//	IsFollow      bool `json:"is_follow,omitempty"`
//	gorm.Model
//}

// DBUser 数据库用户表
type DBUser struct {
	ID            int64  `gorm:"column:id; primaryKey"`
	Username      string `gorm:"column:username; unique"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:followCount"`
	FollowerCount int64  `gorm:"column:followerCount"`
	IsFollow      bool   `gorm:"column:isfollow"`
	Online        bool   `gorm:"column:online"`
}

// 视频表
//type Video struct {
//	ID            int64
//	AuthorID      int64 `json:"author_id,omitempty"`
//	Author        User  `gorm:"foreignKey:AuthorID;references:ID;"`
//	PlayUrl       string
//	CoverUrl      string
//	FavoriteCount int64
//	CommentCount  int64
//	Title         string
//	IsFavorite    bool   `json:"is_favorite,omitempty"`
//	gorm.Model
//}

// 评论表
//type Comment struct {
//	ID            int64
//	UserID        int64
//	User          User `gorm:"foreignKey:UserID;references:ID;"`
//	VideoID       int64
//	Video         Video `gorm:"foreignKey:VideoID;references:ID"`
//	Content       string
//	FavoriteCount int64
//	gorm.Model
//}

// 关注表
type Attention struct {
	ID        int64
	WatchID   int64
	Watch     User `gorm:"foreignKey:WatchID;references:ID;"`
	BeWatchID int64
	BeWatch   User `gorm:"foreignKey:BeWatchID;references:ID;"`
}
