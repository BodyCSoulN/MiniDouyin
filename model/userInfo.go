package model

import (
	"MiniDouyin/storage"
	"log"
)

// SetOnline 根据id设置用户的在线状态，set online = status
func SetOnline(id int64, status bool) error {
	dbUser := storage.DBUser{}
	err := storage.Mysql.Where("id=?", id).Take(&dbUser).Error
	if err != nil {
		log.Println("SetOnline: ", err)
		return err
	}
	//通过User模型的主键id的值作为where条件，更新online字段值
	//用户id是从1开始的
	storage.Mysql.Model(&dbUser).Where("id>?", 0).Update("online", status)
	return nil
}
