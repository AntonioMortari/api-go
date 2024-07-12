package main

import (
	"github.com/antoniomortari/api-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	routes.InitRoutes(&server.RouterGroup)

	server.Run(":3000")

}