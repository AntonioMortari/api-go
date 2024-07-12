package routes

import (
	"github.com/antoniomortari/api-go/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup){

	productController := controller.NewProductController()

	router.GET("/products", productController.GetAll)

}