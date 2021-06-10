package user_service

import (
	"golang.org/x/net/context"
	"second-hand-bbs-go/internal/models"
	"second-hand-bbs-go/internal/models/request"
	"second-hand-bbs-go/pkg/token"
	"second-hand-bbs-go/utils"
	"time"
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
func Register(user *models.User) error {
	user.ID = utils.GetOnlyId()
	if err := user.EncryptionPassword(); err != nil {
		return err
	}
	if err := models.InsetUser(user); err != nil {
		return err
	}
	return nil
}

func GetUserByName(username string) (models.User, error) {
	u, err := models.GetUserByName(username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func ChangePassword(user *models.User) error {
	user.UpdatedAt = time.Now()
	if err := user.EncryptionPassword(); err != nil {
		return err
	}
	if err := models.SaveUser(user); err != nil {
		return err
	}
	return nil
}

func Login(user *request.UserLoginStruct) (string, error) {
	isRight, err := isUserPasswordRight(user)
	if err != nil {
		return "", err
	}
	if !isRight {
		return "", nil
	}

	// 生成token
	tokenContext := token.Context{
		Username:       user.Username,
		ExpirationTime: int64(utils.AppSetting.JwtExpirationTime),
	}
	tokenString, err := token.Sign(context.Background(), tokenContext, utils.AppSetting.JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func isUserPasswordRight(user *request.UserLoginStruct) (bool, error) {
	u, err := GetUserByName(user.Username)
	if err != nil {
		return false, err
	}
	if !u.CheckPassword(user.Password) {
		return false, nil
	}
	return true, nil
}
