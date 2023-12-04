package routes

import (
	"github.com/go-chi/chi/v5"

	handlers "library-api/pkg/library/handler"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	library := handlers.Build()

	r.Get("/books", library.GetBooks)
	r.Post("/books", library.CreateBook)
	r.Post("/books/download", library.DownloadBook)

	return r
}
