package controller

import (
	"net/http"

	"github.com/antoniomortari/api-go/models"
	"github.com/antoniomortari/api-go/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) ProductController {
	return ProductController{
		service: s,
	}
}

func (p *ProductController) GetAll(ctx *gin.Context) {
	products := []models.Product{
		{
			Id: 1,
			Name: "Computador",
			Price: 2000,
		},
	}

	ctx.JSON(http.StatusOK, products)
}