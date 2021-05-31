package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `validate:"min=6,max=10"`
	Age      uint8  `validate:"gte=1,lte=10"`
	Sex      string `validate:"oneof=female male"`
}

func register(c *gin.Context) {
	validate := validator.New()

}
