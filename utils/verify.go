package utils

var (
	IdVerify    = Rules{"ID": {NotEmpty()}}
	LoginVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
)
