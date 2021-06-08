package utils

// 统一声明对应需要的验证
var (
	IdVerify    = Rules{"ID": {NotEmpty()}}
	LoginVerify = Rules{"Username": {NotEmpty(), Ge("3")}, "Password": {NotEmpty()}}
)
