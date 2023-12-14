package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	customErr "library-api/pkg/errors"
	"library-api/pkg/library"
	"library-api/pkg/res"
)

// @Summary		Create Book
// @Description	Creates the book in database and uploads the file into bucket
// @Tags			Books
// @Param			data	body		handler.BookRequestDoc	true	"Required book information"
// @Param			book	formData	file					true	"Book File"
// @Router			/books [post]
// @Accept			multipart/form-data
// @Produce		json
// @Success		200	{object}	handler.BookResponseDoc
func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		customErr.Error(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("error in parser form: %v", err),
		)
		return
	}

	bookInfo := r.FormValue(bookFormField)
	bookModel := new(library.Book)

	if err := json.Unmarshal([]byte(bookInfo), &bookModel); err != nil {
		customErr.Error(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("error in unmarshal book data: %v", err),
		)
		return
	}

	bookPDF, _, err := r.FormFile(bookPDFField)
	if err != nil {
		customErr.Error(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("error in get field: %v", err),
		)
		return
	}
	defer bookPDF.Close()

	book, err := h.librarySvc.CreateBook(bookModel, bookPDF)
	if err != nil {
		customErr.Error(w, http.StatusInternalServerError, err)
		return
	}

	res.JSON(w, http.StatusOK, book)
}
