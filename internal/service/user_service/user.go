package user_service

import "second-hand-bbs-go/internal/models"

// 判断用户是否存在
func IsUserExistByName(name string) (bool, error) {
	return models.IsUserExist(name)
}
