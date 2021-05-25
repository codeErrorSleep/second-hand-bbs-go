package models

type Product struct {
	Model

	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	ProductType string `json:"type"`
	Content     string `json:"content"`
	CreatedBy   string `json:"created_by"`
	ModifiedBy  string `json:"modified_by"`
	State       int    `json:"state"`
}

// 获取产品列表
func GetProducts(pageNum int, pageSize int, maps interface{}) (products []Product) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&products)
	return
}

// 获取商品总数
func GetProductsTotal(maps interface{}) (count int) {
	db.Model(&Product{}).Where(maps).Count(&count)
	return
}

// 判断商品是否存在
func ExistProductByName(name string) bool {
	var product Product
	db.Select("id").Where("name=?", name).First(&product)
	if product.ID > 0 {
		return true
	}
	return false
}

func AddProduct(productName string, productType string, price int) bool {
	db.Create(&Product{
		ProductName: productName,
		ProductType: productType,
		Price:       price,
	})
	return true
}
