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
	requestParams := parseQueryParams(r.URL.Query())

	if requestParams.BookID != "" {
		if err := requestParams.ValidateParams(); err != nil {
			res.JSON(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	books, err := h.libraryRepo.Search(requestParams)
	if err != nil {
		res.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(books) < 1 {
		res.JSON(w, http.StatusNotFound, fmt.Errorf("book not found").Error())
		return
	}

	res.JSON(w, http.StatusOK, books)
}
