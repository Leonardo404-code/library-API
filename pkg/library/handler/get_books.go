package handler

import (
	"fmt"
	"net/http"

	"library-api/pkg/res"
)

// @Summary		Get Books
// @Description	Get books matching passed filters
// @Tags			Books
// @Router			/books [get]
// @Param			name	query	string	false	"Name of the book you want to search"
// @Param			book_id	query	string	false	"ID of the book you want to search"
// @Produce		json
// @Success		200	{array}	handler.BookResponseDoc
func (h *handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	request := parseQueryParams(r.URL.Query())

	books, err := h.libraryRepo.Search(request)
	if err != nil {
		res.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(books) < 1 {
		err = fmt.Errorf("not found: book not found")
		res.JSON(w, http.StatusNotFound, err.Error())
		return
	}

	res.JSON(w, http.StatusOK, books)
}
