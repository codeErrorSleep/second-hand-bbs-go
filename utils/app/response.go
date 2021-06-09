package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"second-hand-bbs-go/utils/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(errCode int, data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}
