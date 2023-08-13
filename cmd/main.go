package main

import (
	"github.com/edmundcheng221/dollar-store-pokedex-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.InitializeRoutes(r)

	r.Run(":8080")
}
