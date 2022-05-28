package storage

import (
	"fmt"

	"github.com/MiniDouyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Mysql *gorm.DB

func init() {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Mysql.Username, config.Conf.Mysql.Password,
		config.Conf.Mysql.Host, config.Conf.Mysql.Port,
		config.Conf.Mysql.Database)
	var err error
	Mysql, err = gorm.Open(mysql.Open(uri), &gorm.Config{
		// 迁移时使用单数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Errorf("mysql Connect err: %v", err))
	}
	if err = dataMigrate(); err != nil {
		panic(fmt.Errorf("mysql Migrate err: %v", err))
	}
}

func dataMigrate() error {
	err := Mysql.AutoMigrate(&User{}, &Video{}, &Comment{}, &Attention{})
	return err
}
