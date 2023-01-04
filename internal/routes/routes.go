package routes

import (
	handlers "library-api/pkg/library/handler"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	library := handlers.Build()

	r.Get("/books", library.GetBooks)

	return r
}
