package repository

import (
	"database/sql"
	"fmt"

	"github.com/antoniomortari/api-go/models"
	"github.com/antoniomortari/api-go/rest_err"
)

type ProductRepository struct {
	connection *sql.DB
}

func (pr *ProductRepository) GetAll() ([]models.Product, *rest_err.RestError) {
	query := "SELECT * FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		restErr := rest_err.NewInternalServerError()
		return nil, restErr
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
			restErr := rest_err.NewInternalServerError()
			return nil, restErr
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) Create(product models.Product) (int, *rest_err.RestError) {
	query, err := pr.connection.Prepare("INSERT INTO products" + " (name, price)" + " VALUES($1, $2)" + " RETURNING ID")
	if err != nil {
		restErr := rest_err.NewInternalServerError()
		return 0, restErr
	}

	var id int

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		restErr := rest_err.NewInternalServerError()
		return 0, restErr
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetById(id int) (*models.Product, *rest_err.RestError) {
	query, err := pr.connection.Prepare("SELECT * FROM products" + " WHERE id = $1")
	if err != nil {
		restErr := rest_err.NewInternalServerError()
		return nil, restErr
	}

	var product models.Product

	err = query.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		if(err == sql.ErrNoRows){
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("Produto de id %d não encontrado", id))
		}
		restErr := rest_err.NewInternalServerError()
		return nil, restErr
	}

	query.Close()

	fmt.Printf(product.Name)

	return &product, nil
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}
