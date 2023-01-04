package main

import (
	"library-api/internal/database"
	"library-api/internal/routes"
	"log"
	"net/http"
)

func init() {
	database.Connect()
}

func main() {
	routes := routes.SetupRoutes()

	log.Println("Server Started at http://localhost:3000")

	if err := http.ListenAndServe(":3000", routes); err != nil {
		log.Fatalf("error in SetupRoutes(): %v", err)
	}
}
