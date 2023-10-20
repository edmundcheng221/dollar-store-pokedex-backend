package main

import (
	"log"
	"net/http"
	"os"

	"github.com/edmundcheng221/dollar-store-pokedex-backend/router"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router.InitializeRoutes()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
