package v1

import (
	"github.com/gin-gonic/gin"
	"second-hand-bbs-go/internal/models"
	"second-hand-bbs-go/internal/models/request"
	"second-hand-bbs-go/internal/service/user_service"
	"second-hand-bbs-go/utils"
	"second-hand-bbs-go/utils/app"
	"second-hand-bbs-go/utils/e"
)

// 用户注册
func Register(c *gin.Context) {
	appG := app.Gin{c}
	// 绑定对应的参数
	var user models.User
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.LoginVerify); err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	// 判断对应的参数格式是否正确
	isExist, err := user_service.IsUserExistByName(user.Username)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	if isExist {
		appG.Response(e.USER_NAME_EXIST, nil)
		return
	}
	// 插入
	err = user_service.Register(&user)
	if err != nil {
		appG.Response(e.USER_CREATE_FALSE, nil)
		return
	}
	appG.Response(e.SUCCESS, nil)
}

// 更新用户密码
func ChangeUserPassword(c *gin.Context) {
	appG := app.Gin{c}
	// 绑定对应的参数
	var user request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	// 判断原密码是否正确
	modelUser, err := user_service.GetUserByName(&user)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	if modelUser.Password == "" || modelUser.Password != user.Password {
		appG.Response(e.USER_PASSWORD_WRONG, nil)
		return
	}
	// 更新用户信息
	modelUser.Password = user.NewPassword
	err = user_service.ChangePassword(&modelUser)
	if err != nil {
		appG.Response(e.USER_PASSWORD_WRONG, err.Error())
		return
	}
	appG.Response(e.SUCCESS, nil)
}
