package service

import (
	"github.com/antoniomortari/api-go/models"
	"github.com/antoniomortari/api-go/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func (ps *ProductService) GetAll() ([]models.Product, error) {
	return ps.repository.GetAll()
}

func NewProductService(r repository.ProductRepository) ProductService {
	return ProductService{
		repository: r,
	}
}