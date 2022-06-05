package model

type UserResponse struct {
	ID            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}
