package model

import (
	"MiniDouyin/storage"
	"log"
	"time"
)

func GetUserInfoByName(username string) (storage.DBUser, error) {
	var dbUser storage.DBUser
	err := storage.Mysql.Where("username=?", username).Take(&dbUser).Error
	if err != nil {
		log.Println("GetUserInfoByName: ", err)
	}
	//fmt.Println(dbUser)
	return dbUser, err
}

// CheckLogin 返回值：用户是否存在 密码是否正确 用户信息
func CheckLogin(username, password string) (bool, bool, storage.DBUser) {
	dbUser, err := GetUserInfoByName(username)
	if err != nil {
		return false, false, storage.DBUser{}
	}
	if dbUser.Password == password {
		return true, true, dbUser
	}
	return true, false, storage.DBUser{}
}

// InitUserLoginInfo
// 将数据库所有用户导入内存，存入controller.UsersLoginInfo;
// controller.UsersLoginInfo是一个字典结构，map[string] controller.User;
// 其中 key= token, value= 用户信息
func InitUserLoginInfo() {
	userList, _ := GetAllUserInfo()
	for _, user := range userList {
		DYUser := storage.User{
			ID:            user.ID,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		token, err := Setting(user.Username, user.ID)
		if err != nil {
			log.Println("initUserLoginInfo:", err)
			continue
		}
		storage.UsersLoginInfo[token] = DYUser
		storage.TokenEndTime[user.Username] = storage.UsrToken{
			Token:   token,
			EndTime: time.Now().Add(storage.TokenValidTime),
		}
	}
}
