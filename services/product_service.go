package services

import (
	"eCommerceAPI/models"
	"eCommerceAPI/repository"
)

type ProductService struct {
	Repo *repository.ProductRepo
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) CreateProduct(p *models.Product) error {
	// In a real app, you would add business logic here
	// e.g., if p.Price <= 0 {return errors.New("price must be positive")}
	return s.Repo.Create(p)
}
