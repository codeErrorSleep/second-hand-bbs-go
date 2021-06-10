package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
		//登录验证权限绑定
		//apiv1.Use(middleware.AuthMiddleware())
		//获取多个商品列表
		apiv1.POST("/products", v1.GetProducts)
		//获取商品列表总数
		apiv1.GET("/products/count", v1.GetProductTotal)
		//获取指定商品
		apiv1.POST("/product/:id", v1.GetProduct)
		//新建商品
		apiv1.POST("/product/add", v1.AddProduct)
		//更新指定商品
		apiv1.POST("/product/update/:id", v1.EditProduct)
		//删除指定商品
		apiv1.POST("/product/delete/:id", v1.DeleteProduct)
	}

	// 用户管理
	auth := r.Group("/api/v1/user")
	{
		// 注册
		auth.POST("/register", v1.Register)
		// 登录
		auth.POST("/login", v1.Login)
		// 修改密码
		auth.POST("/change_password", v1.ChangeUserPassword)
	}

	return r
}
