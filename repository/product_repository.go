package repository

import (
	"database/sql"
	"fmt"

	"github.com/antoniomortari/api-go/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func (pr *ProductRepository) GetAll() ([]models.Product, error) {
	query := "SELECT * FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var productList = []models.Product{}
	var productObj = models.Product{}

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) Create(product models.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO products" + " (name, price)" + " VALUES($1, $2)" + " RETURNING ID")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var id int

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id,nil
}

func (pr *ProductRepository) GetById(id int) (*models.Product, error){
	query, err := pr.connection.Prepare("SELECT * FROM products" + " WHERE id = $1")
	if(err != nil){
		fmt.Println(err)
		return nil, err
	}

	var product models.Product

	err = query.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price)
	if(err != nil){
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &product, nil
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}
