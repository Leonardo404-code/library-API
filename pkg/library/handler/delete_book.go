package handler

import (
	"net/http"

	"library-api/pkg/res"
)

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	requestParams := parseQueryParams(r.URL.Query())

	if err := requestParams.ValidateParams(); err != nil {
		res.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.libraryRepo.DeleteBook(requestParams); err != nil {
		res.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.JSON(w, http.StatusOK, nil)
}
