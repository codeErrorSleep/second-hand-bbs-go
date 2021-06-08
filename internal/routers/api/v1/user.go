package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"second-hand-bbs-go/internal/models"
	"second-hand-bbs-go/internal/service/user_service"
	"second-hand-bbs-go/utils"
	"second-hand-bbs-go/utils/app"
	"second-hand-bbs-go/utils/e"
)

func Register(c *gin.Context) {
	appG := app.Gin{c}
	// 绑定对应的参数
	var user models.User
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.LoginVerify); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error())
		return
	}
	// 判断对应的参数格式是否正确
	isExist, err := user_service.IsUserExistByName(user.Username)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error())
		return
	}
	if isExist {
		appG.Response(http.StatusOK, e.USER_NAME_EXIST, nil)
		return
	}
	// 插入
	err = user_service.Register(user)
	if err != nil {
		appG.Response(http.StatusOK, e.USER_CREATE_FALSE, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
