package controller

import (
	"net/http"

	"github.com/antoniomortari/api-go/models"
	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetAll(ctx *gin.Context) {
	products := []models.Product{
		{
			Id: 1,
			Name: "Computador",
			Price: 2000,
		},
	}

	ctx.JSON(http.StatusOK, products)
}