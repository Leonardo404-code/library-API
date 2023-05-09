package handler

import (
	"encoding/json"
	"io"
	"net/http"

	customErr "library-api/pkg/errors"
	"library-api/pkg/library"
	"library-api/pkg/res"
)

// @Summary		Send Book
// @Description	Creates the book in database and uploads the file into bucket
// @Tags			Books
// @Router			/books [post]
// @Param			data	body		handler.BookRequestDoc	true	"Required book information"
// @Param			book	formData	file					true	"Book File"
// @Success		200
func (h *handler) SendBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookModel := new(library.Book)

	if err := json.Unmarshal(body, &bookModel); err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	book, err := h.librarySvc.SendBook(bookModel)
	if err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	res.JSON(w, http.StatusOK, book)
}
