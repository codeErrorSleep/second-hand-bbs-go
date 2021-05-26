package models

import "github.com/jinzhu/gorm"

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

// 获取单个商品
func GetProduct(id int) (*Product, error) {
	var product Product
	err := db.Where("id=?", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &product, nil
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

// 判断商品是否存在
func ExistProductByID(id int) (bool, error) {
	var product Product
	err := db.Select("id").Where("id = ?", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if product.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddProduct(productName string, productType string, price int) bool {
	db.Create(&Product{
		ProductName: productName,
		ProductType: productType,
		Price:       price,
	})
	return true
}
