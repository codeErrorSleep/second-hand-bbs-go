package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"math/rand"
	"time"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * AppSetting.PageSize
	}

	return result
}

// 获取唯一id
func GetOnlyId() uint {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(999999)
	return uint(result)
}
