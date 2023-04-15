package handler

import "net/http"

// @Summary		Download Book
// @Description	Download your books by passing the necessary parameters
// @Tags			Cloud
// @Param			id	query	string	true	"ID of the book you want to download"
// @Router			/book/download [post]
// @Success		200
func (h *handler) DownloadBook(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
