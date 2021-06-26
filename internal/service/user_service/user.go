package user_service

import (
	"second-hand-bbs-go/internal/models"
	"second-hand-bbs-go/internal/models/request"
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
	err := models.InsetUser(user)
	if err != nil {
		return err
	}
	return nil
}

// 验证密码是否正确
func GetUserByName(user *request.ChangePasswordStruct) (models.User, error) {
	u, err := models.GetUserByName(user.Username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func ChangePassword(user *models.User) error {
	user.UpdatedAt = time.Now()
	err := models.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}
<<<<<<< HEAD

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
=======
>>>>>>> parent of fee293e (登录生成token,后续添加验证)
