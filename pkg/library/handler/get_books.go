package handler

import (
	"net/http"
)

//	@Summary		Get Books
//	@Description	Get books matching passed filters
//	@Tags			Books
//	@Router			/books [get]
//	@Param			name	query	string	false	"Name of the book you want to search"
//	@Param			id		query	string	false	"ID of the book you want to search"
//	@Produce		json
//	@Success		200	{array}	handler.BookResponseDoc
func (h *handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
