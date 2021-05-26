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
