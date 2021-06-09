package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"second-hand-bbs-go/internal/service/product_service"
	"second-hand-bbs-go/utils"
	"second-hand-bbs-go/utils/app"
	"second-hand-bbs-go/utils/e"
)

func GetProducts(c *gin.Context) {
	appG := app.Gin{c}
	// 验证参数
	valid := validation.Validation{}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	productService := product_service.Product{
		State:    1,
		PageNum:  utils.GetPage(c),
		PageSize: utils.AppSetting.PageSize,
	}
	// 获取商品列表
	products, err := productService.GetAll()
	if err != nil {
		appG.Response(e.ERROR_GET_PRODUCTS_FAIL, nil)
		return
	}
	appG.Response(e.SUCCESS, products)

}

func GetProductTotal(c *gin.Context) {
	appG := app.Gin{c}
	valid := validation.Validation{}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	productService := product_service.Product{}
	// 获取商品总数
	total, err := productService.Count()
	if err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	appG.Response(e.SUCCESS, total)
}

// @Summary 获取单个商品
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/product [post]
func GetProduct(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id)
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	productService := product_service.Product{ID: id}
	exists, err := productService.ExistById()

	if err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	if !exists {
		appG.Response(e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	product, err := productService.Get()
	if err != nil {
		appG.Response(e.ERROR_GET_PRODUCT_FAIL, nil)
		return
	}
	appG.Response(e.SUCCESS, product)
}

type AddProductForm struct {
	ProductName string `form:"product_name" valid:"Required;MaxSize(100)"`
	Price       int    `form:"price" valid:"Required;MaxSize(255)"`
	ProductType string `form:"type" valid:"Required;MaxSize(2)"`
	Content     string `form:"content" valid:"Required;MaxSize(1024)"`
	CreatedBy   string `form:"created_by" valid:"Required;MaxSize(100)"`
	State       int    `form:"state" valid:"Range(0,1)"`
}

// @Summary 新增商品
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/product [post]
func AddProduct(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddProductForm
	)

	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}

	productService := product_service.Product{
		ProductName: form.ProductName,
		Price:       form.Price,
		ProductType: form.ProductType,
		Content:     form.Content,
		CreatedBy:   form.CreatedBy,
		State:       form.State,
	}

	if err := productService.Add(); err != nil {
		appG.Response(e.ERROR_ADD_PRODUCT_FAIL, nil)
		return
	}
	appG.Response(e.SUCCESS, nil)
}

//修改商品信息
func EditProduct(c *gin.Context) {
}

//删除商品信息
func DeleteProduct(c *gin.Context) {
}
