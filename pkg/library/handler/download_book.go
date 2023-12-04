package handler

import (
	"net/http"

	"library-api/pkg/res"
)

// @Summary		Download Book
// @Description	Download your books by passing the necessary parameters
// @Tags			Cloud
// @Param			book_id	query	string	true	"ID of the book you want to download"
// @Router			/books/download [post]
// @Success		200
func (h *handler) DownloadBook(w http.ResponseWriter, r *http.Request) {
	requestParams := parseQueryParams(r.URL.Query())

	if err := h.librarySvc.DownloadBook(requestParams); err != nil {
		res.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.JSON(w, http.StatusOK, "")
}
