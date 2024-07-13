package repository

import (
	"database/sql"
	"fmt"

	"github.com/antoniomortari/api-go/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func(pr *ProductRepository) GetAll() ([]models.Product, error){
	query := "SELECT * FROM products"
	rows, err := pr.connection.Query(query)
	if(err != nil){
		fmt.Println(err)
		return nil, err
	}

	var productList = []models.Product{}
	var productObj = models.Product{}

	for rows.Next(){
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price,
		)
		if(err != nil){
			fmt.Println(err)
			return nil, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}