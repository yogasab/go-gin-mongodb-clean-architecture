package main

import (
	"go-gin-mongodb-clean-architecture/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitializeRoutes(router)

	router.Run(":5000")
}
