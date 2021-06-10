package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_NOT_EXIST_PRODUCT: "该商品不存在",
	ERROR_GET_PRODUCT_FAIL:  "获取商品失败",
	ERROR_GET_PRODUCTS_FAIL: "获取商品列表失败",
	ERROR_ADD_PRODUCT_FAIL:  "添加商品失败",

	USER_NAME_EXIST:     "用户名已经存在",
	USER_CREATE_FALSE:   "用户注册失败",
	USER_PASSWORD_WRONG: "密码错误",
	USER_UPDATE_FAIL:    "用户信息更新错误",
	USER_LOGIN_FAIL:     "登录失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
