package library

import (
	"net/http"
)

type Handlers interface {
	GetBooks(http.ResponseWriter, *http.Request)
	CreateBook(http.ResponseWriter, *http.Request)
	DeleteBook(http.ResponseWriter, *http.Request)
	DownloadBook(http.ResponseWriter, *http.Request)
}
