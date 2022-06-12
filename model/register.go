package model

import (
	"log"
	"net/http"

	"github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
)

func IsExist(username string) bool {
	var dbUser storage.DBUser
	err := storage.Mysql.Where("username=?", username).Take(&dbUser).Error
	if err != nil {
		log.Println("IsExist: ", err)
	}
	return err == nil
}

func CheckUsername(username string, c *gin.Context) bool {
	if len(username) > 32 {
		c.JSON(http.StatusOK, storage.UserLoginResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: "用户名超过32字符"},
		})
		return false
	}
	return true
}

func CheckPassword(password string, c *gin.Context) bool {
	if len(password) > 32 {
		c.JSON(http.StatusOK, storage.UserLoginResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: "密码超过32字符"},
		})
		return false
	}
	return true
}

// GetNewIdByName
// 获取数据库maxId，newId = maxId+1
func GetNewIdByName(username string) (int64, error) {
	var dbUser storage.DBUser
	err := storage.Mysql.Last(&dbUser).Error
	if err != nil {
		log.Println("GetNewIdByName: ", err)
	}
	return dbUser.ID + 1, err
}

// AddUserIntoDB 向数据库中添加一个用户,User是数据库用户，表中数据比返回给客户端的数据更多
func AddUserIntoDB(dbUser storage.DBUser) {
	err := storage.Mysql.Create(&dbUser).Error
	if err != nil {
		log.Println("AddUserIntoDB: ", err)
		return
	}
}
