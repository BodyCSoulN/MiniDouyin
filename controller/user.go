package controller

import (
	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Register
// 获取username,password;
// 根据username判断用户是否存在;
// 生成User结构体,生成token序列;
// 更新token结束时间;
// 更新userLoginInfo,token映射User{};
// 新用户信息添加进数据库中
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if model.IsExist(username) { // 查询用户是否存在
		c.JSON(http.StatusOK, storage.UserLoginResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}
	// 检查密码长度合法性
	usernameOK := model.CheckUsername(username, c)
	if !usernameOK {
		return
	}
	passwordOK := model.CheckPassword(password, c)
	if !passwordOK {
		return
	}
	// 开始注册，生成新ID
	newUserId, err := model.GetNewIdByName(username)
	if err != nil {
		log.Println("新用户id生成失败", err)
		c.JSON(http.StatusOK, storage.UserLoginResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: "注册失败"},
		})
		return
	}

	newUser := storage.User{
		ID:       newUserId,
		Name:     username,
		IsFollow: true,
	}
	// 生成token，更新token映射信息
	token, _ := model.Setting(username, newUser.ID) // 给新用户生成token序列
	storage.TokenEndTime[username] = storage.UsrToken{
		Token:   token,
		EndTime: time.Now().Add(storage.TokenValidTime),
	}
	storage.UsersLoginInfo[token] = newUser // 更新UserLoginInfo

	dbUser := storage.DBUser{
		ID:            newUser.ID,
		Username:      newUser.Name,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      true,
		Online:        false,
	}
	//fmt.Println("Register: id=", dbUser.ID)
	model.AddUserIntoDB(dbUser) // 将用户信息存入数据库

	c.JSON(http.StatusOK, storage.UserLoginResponse{
		Response: storage.Response{StatusCode: 0},
		UserId:   newUserId,
		Token:    token,
	})
}

// Login 获取username, password,
// 然后判断用户是否存在，同时获取用户信息
// 判断用户是否存在 or 密码是否正确
// 判断用户是否已经在线
// 更新token
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	exist, ok, dbUser := model.CheckLogin(username, password)

	if !exist || !ok { // 用户不存在 or 密码错误
		var respond string
		if exist == false {
			respond = "User doesn't exist"
		} else if ok == false {
			respond = "Password error"
		}
		c.JSON(http.StatusOK, storage.UserLoginResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: respond},
		})
		return
	}
	if dbUser.Online { //用户已在线
		c.JSON(http.StatusOK, storage.UserLoginResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: "user is online, repeat login"},
		})
		return
	}
	// 用户存在 && 密码正确
	token, err := UPDToken(dbUser.Username, dbUser.ID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, storage.UserLoginResponse{
		Response: storage.Response{StatusCode: 0},
		UserId:   dbUser.ID,
		Token:    token,
	})
}

// UserInfo
//首先获取token,根据token从内存找出用户的信息
// 然后设置用户为在线状态，最后将用户信息返回给客户端展示出来
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	DYUser, exist := storage.UsersLoginInfo[token]
	//fmt.Println("UserInfo: id=", DYUser.ID)
	err := model.SetOnline(DYUser.ID, true)
	if err != nil {
		log.Println("UserInfo:更新在线状态失败: ", err)
		c.JSON(http.StatusOK, storage.UserResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: "服务器错误"},
		})
		return
	}
	if exist { // 用户存在
		c.JSON(http.StatusOK, storage.UserResponse{
			Response: storage.Response{StatusCode: 0},
			User:     DYUser,
		})
	} else {
		c.JSON(http.StatusOK, storage.UserResponse{
			Response: storage.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

// UPDToken
// 为用户重新生成token
// 将用户原token映射信息从userInfoLogin删除
// 更改token映射的用户信息, 更新token结束时间
func UPDToken(username string, id int64) (string, error) {
	newToken, _ := model.Setting(username, id) //为用户重新生成token
	oldToken := storage.TokenEndTime[username].Token
	userInfo := storage.UsersLoginInfo[oldToken]
	// fmt.Println("UPDToken: id=", userInfo.ID)
	delete(storage.UsersLoginInfo, oldToken)

	storage.UsersLoginInfo[newToken] = userInfo
	storage.TokenEndTime[username] = storage.UsrToken{
		Token:   newToken,
		EndTime: time.Now().Add(storage.TokenValidTime),
	}
	return newToken, nil
}
