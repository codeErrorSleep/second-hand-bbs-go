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

func GetProductTotal() (int, error) {
	var count int
	if err := db.Model(&Product{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 获取产品列表
func GetProducts(pageNum int, pageSize int, maps interface{}) ([]*Product, error) {
	var products []*Product
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
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

func AddProduct(data map[string]interface{}) error {
	product := map[string]interface{}{
		"ProductName": data["product_name"].(string),
		"Price":       data["price"].(int),
		"ProductType": data["type"].(string),
		"Content":     data["content"].(string),
		"State":       data["state"].(int),
		"CreatedBy":   data["created_by"].(string),
		"ModifiedBy":  data["modified_by"].(string),
	}
	if err := db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}
