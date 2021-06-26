package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"` // 主键ID
	Username  string    `json:"username" gorm:"column:user_name"`
	Password  string    `json:"password" `
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
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
func InsetUser(user *User) error {
	user.CreatedAt = time.Now()
	err := db.Table("user").Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// 通过用户名获取用户
func GetUserByName(username string) (User, error) {
	var user User
	err := db.Table("user").Where("user_name=?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

// 修改密码
func SaveUser(user *User) error {
	err := db.Table("user").Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
