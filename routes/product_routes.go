package routes

import (
	"github.com/antoniomortari/api-go/controller"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(router *gin.RouterGroup, productController controller.ProductController) {

	router.GET("/products", productController.GetAll)

}
