package handler

import (
	"net/http"
)

// @Summary		Send Book
// @Description	Creates the book in database and uploads the file into bucket
// @Tags			Books
// @Router			/books [post]
// @Param			data	body		handler.BookRequestDoc	true	"Required book information"
// @Param			book	formData	file					true	"Book File"
// @Success		200
func (h *handler) SendBook(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented")
}
