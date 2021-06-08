package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:user_name"`
	Password string `json:"password" `
}

// 获取名字是否存在
func IsUserExist(name string) (int, error) {
	var count int
	err := db.Table("user").Where("user_name=?", name).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

// 插入新用户
func InsetUser(user User) error {
	user.CreatedAt = time.Now()
	err := db.Table("user").Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
