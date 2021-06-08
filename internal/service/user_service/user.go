package user_service

import (
	"second-hand-bbs-go/internal/models"
	"second-hand-bbs-go/utils"
)

// 判断用户是否存在
func IsUserExistByName(name string) (bool, error) {
	count, err := models.IsUserExist(name)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 注册用户
func Register(user models.User) error {
	user.ID = utils.GetOnlyId()
	err := models.InsetUser(user)
	if err != nil {
		return err
	}
	return nil
}
