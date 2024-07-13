package service

import (
	"fmt"

	"github.com/antoniomortari/api-go/models"
	"github.com/antoniomortari/api-go/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func (ps *ProductService) GetAll() ([]models.Product, error) {
	return ps.repository.GetAll()
}

func (ps *ProductService) Create(product models.Product) (int, error){

	result, err := ps.repository.Create(product)
	if(err != nil){
		fmt.Println(err)
		return 0,nil
	}

	return result, nil

}

func (ps *ProductService) GetById(id int) (*models.Product, error){
	result, err := ps.repository.GetById(id)
	if(err != nil){
		return nil, err
	}

	return result, nil
}

func NewProductService(r repository.ProductRepository) ProductService {
	return ProductService{
		repository: r,
	}
}