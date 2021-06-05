package models

import "github.com/jinzhu/gorm"

type User struct {
	Model

	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 获取名字是否存在
func IsUserExist(name string) (bool, error) {
	var count int
	err := db.Where("user_name=?", name).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if count > 0 {
		return false, nil
	} else {
		return true, nil
	}

}
