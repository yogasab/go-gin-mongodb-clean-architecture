package main

import (
	"go-gin-mongodb-clean-architecture/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitializeRoutes(router)

	router.Static("/assets/images", "./assets/images")
	router.Run(":5000")
}
