package service

import (
	"fmt"

	"github.com/antoniomortari/api-go/models"
	"github.com/antoniomortari/api-go/repository"
	"github.com/antoniomortari/api-go/rest_err"
)

type ProductService struct {
	repository repository.ProductRepository
}

func (ps *ProductService) GetAll() ([]models.Product, *rest_err.RestError) {
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

func (ps *ProductService) GetById(id int) (*models.Product, *rest_err.RestError){
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