package main

import (
	"log"
	"net/http"

	"library-api/internal/database"
	"library-api/internal/routes"
)

func init() {
	database.Connect()
}

//	@title			Library API
//	@version		1.0
//	@description	API for storing and download PDFs into a Google Cloud Plattform bucket
//	@contect.name	Leonardo Bispo
//	@contact.email	leonardobispo1000@gmail.com
//
// license.name Apache 2.0
//
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	routes := routes.SetupRoutes()

	log.Println("Server Started at http://localhost:3000")

	if err := http.ListenAndServe(":3000", routes); err != nil {
		log.Fatalf("error in SetupRoutes(): %v", err)
	}
}
