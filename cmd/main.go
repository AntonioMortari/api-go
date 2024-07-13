package main

import (
	"github.com/antoniomortari/api-go/controller"
	"github.com/antoniomortari/api-go/db"
	"github.com/antoniomortari/api-go/repository"
	"github.com/antoniomortari/api-go/routes"
	"github.com/antoniomortari/api-go/service"
	"github.com/gin-gonic/gin"
)

func main() {

	connection, err := db.Connection()
	if err != nil {
		panic(err)
	}

	server := gin.Default()

	// product
	productRepository := repository.NewProductRepository(connection)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	routes.InitProductRoutes(&server.RouterGroup, productController)

	server.Run(":3000")

}
