package model

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

// 测试用秘钥
var jwtKey = []byte("Serein404")

type MyClaims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// Setting 颁发token
func Setting(name string, id int64) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &MyClaims{
		Id:       id,
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //开始时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		//fmt.Println(err)
		log.Println(err)
		return "", err
	}
	return tokenStr, err
}

//Getting 解析token
func Getting(tokenStr string) (*MyClaims, error) {
	claims, err := parseToken(tokenStr)
	if err != nil {
		//fmt.Println(err)
		//log.Fatal(err)
		return nil, err
	}
	return claims, nil
}

func parseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		//fmt.Println("解析失败", err)
		//log.Fatal(err)
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// Test 测试jwt是否解析正确
func Test() {
	name := "serein"
	id := int64(1)
	token, _ := Setting(name, id)
	fmt.Println("!", token)
	//fmt.Println("#", any(Getting(token)))
}
