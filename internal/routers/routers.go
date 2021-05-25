package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "second-hand-bbs-go/docs"
	"second-hand-bbs-go/internal/routers/api/v1"
	"second-hand-bbs-go/utils"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(utils.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	{
		//获取商品列表
		apiv1.GET("/product", v1.GetProducts)
		//新建商品
		apiv1.POST("/product", v1.AddProduct)
		//更新指定商品
		apiv1.POST("/product/update/:id", v1.EditProduct)
		//删除指定商品
		apiv1.POST("/product/delete/:id", v1.DeleteProduct)
	}

	return r
}
