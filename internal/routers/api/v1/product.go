package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"second-hand-bbs-go/utils/app"
	"second-hand-bbs-go/utils/e"
)

func GetProducts(c *gin.Context) {
}

// @Summary 获取单个商品
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/product [post]
func GetProduct(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

}

// @Summary 新增商品
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/product [post]
func AddProduct(c *gin.Context) {
	// 获取相应的数据
	productName := c.Query("product_name")
	productType := c.Query("product_type")
	price := c.Query("price")
	state := c.Query("state")
	createdBy := c.Query("created_by")
	// 判断数据正确性
	valid := validation.Validation{}
	valid.Required(productName, "product_name").Message("名称不能为空")
	valid.MaxSize(productName, 100, "product_name").Message("名称最长为50字符")
	valid.Required(productType, "product_type").Message("商品类型不能为空")
	valid.Required(price, "price").Message("价格不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为50字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

}

//修改商品信息
func EditProduct(c *gin.Context) {
}

//删除商品信息
func DeleteProduct(c *gin.Context) {
}
