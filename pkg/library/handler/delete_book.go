package handler

import (
	"net/http"

	"library-api/pkg/res"
)

// @Summary		Delete Book
// @Description	Delete a book in database and delete a file from bucket
// @Tags			Books
// @Router			/books [delete]
// @Param			book_id	query	string	true	"ID of the bucket you want to delete"
// @Success		200		"OK"
func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	requestParams := parseQueryParams(r.URL.Query())

	if err := h.librarySvc.DeleteBook(requestParams); err != nil {
		res.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.JSON(w, http.StatusOK, nil)
}
