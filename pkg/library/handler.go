package library

import (
	"net/http"
)

type Handlers interface {
	GetBooks(http.ResponseWriter, *http.Request)

	SendBook(http.ResponseWriter, *http.Request)

	DownloadBook(http.ResponseWriter, *http.Request)
}
