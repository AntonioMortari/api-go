package routes

import (
	"net/http"

	"github.com/antoniomortari/api-go/controller"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(router *gin.RouterGroup, productController controller.ProductController) {

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	

	router.GET("/products", productController.GetAll)
	router.GET("/products/:id", productController.GetById)
	router.POST("/products", productController.Create)
	router.DELETE("/products/:id", productController.DeleteById)
}
