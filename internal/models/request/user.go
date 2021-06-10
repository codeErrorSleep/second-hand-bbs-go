package request

// 修改密码结构
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type UserLoginStruct struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
