package model

import (
	"github.com/MiniDouyin/storage"
	"log"
	"time"
)

func GetAllUserInfo() ([]storage.DBUser, error) {
	var userList []storage.DBUser
	err := storage.Mysql.Find(&userList).Error //select * from `douyinuser`
	if err != nil {
		log.Println("GetAllUserInfo: ", err)
	}
	return userList, nil
}

func GetUserInfoByID(id int64) (storage.DBUser, error) {
	var dbUser storage.DBUser
	err := storage.Mysql.Where("id=?", id).Take(&dbUser).Error
	if err != nil {
		log.Println("GetUserInfoByID: ", err)
	}
	return dbUser, nil
}

func GetPasswordByName(username string) (string, error) {
	pwd := ""
	err := storage.Mysql.Where("username=?", username).Take(&pwd).Error
	if err != nil {
		log.Println("GetPasswordByName: ", err)
	}
	return pwd, nil
}

// ResetOnlineStatus 重置数据库中用户在线状态为下线
func ResetOnlineStatus() {
	err := storage.Mysql.Model(&storage.DBUser{}).Where("id>?", 0).Update("online", 0).Error
	if err != nil {
		log.Println("ResetOnlineStatus: ", err)
	}
}

// TokenIsValid token有效返回true，否则false
func TokenIsValid(token string) bool {
	/*
		此处会有一个潜在问题，如果用户在刷视频过程中突然token失效怎么办，其实在实际生产中，这几乎不会出现，
		但是在这里如果出现了，我并不能解决。原因在于我判断token过期是在进入Feed函数时判断的，但是用户在刷
		视频时突然token过期，也将会在Feed函数里判断，并不能识别出是那种情况，进而无法知道是否要强制用户重
		新登陆。我的解决方案是，将token过期时间设置为很长，也就是说刷视频中token突然过期的情况几乎不能出现
	*/
	claims := Getting(token)
	if time.Now().After(storage.TokenEndTime[claims.Username].EndTime) {
		log.Println("token失效")
		// 用户登录状态改为下线
		err := SetOnline(storage.UsersLoginInfo[token].ID, false)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	return true
}
