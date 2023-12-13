package handler

import (
	"net/http"

	"library-api/pkg/res"
)

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	requestParams := parseQueryParams(r.URL.Query())

	if err := h.librarySvc.DeleteBook(requestParams); err != nil {
		res.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.JSON(w, http.StatusOK, nil)
}
