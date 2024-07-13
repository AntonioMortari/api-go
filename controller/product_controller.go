package controller

import (
	"net/http"
	"strconv"

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

func (pc *ProductController) GetAll(ctx *gin.Context) {
	products, err := pc.service.GetAll()
	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func(pc *ProductController) Create(ctx *gin.Context){
	var product = models.Product{}

	err := ctx.BindJSON(&product)
	if(err != nil){
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := pc.service.Create(product)
	if(err != nil){
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (pc *ProductController) GetById(ctx *gin.Context){
	param := ctx.Param("id")
	
	id, err := strconv.Atoi(param)
	if(err != nil){
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := pc.service.GetById(id)
	if(err != nil){
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}