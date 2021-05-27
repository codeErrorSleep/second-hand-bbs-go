package product_service

import "second-hand-bbs-go/internal/models"

type Product struct {
	ID          int
	ProductName string
	Price       int
	ProductType string
	Content     string
	CreatedBy   string
	ModifiedBy  string
	State       int

	PageNum  int
	PageSize int
}

func (p *Product) ExistById() (bool, error) {
	return models.ExistProductByID(p.ID)
}

func (p *Product) Get() (*models.Product, error) {
	product, err := models.GetProduct(p.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Count() (int, error) {
	count, err := models.GetProductTotal()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *Product) GetAll() ([]*models.Product, error) {
	var products []*models.Product

	products, err := models.GetProducts(p.PageNum, p.PageSize, p.getMaps())
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if p.State != -1 {
		maps["state"] = p.State
	}
	return maps
}

func (p *Product) Add() error {
	product := map[string]interface{}{
		"product_name": p.ProductName,
		"price":        p.Price,
		"type":         p.ProductType,
		"content":      p.Content,
		"created_by":   p.CreatedBy,
		"state":        p.State,
	}
	if err := models.AddProduct(product); err != nil {
		return err
	}
	return nil
}
