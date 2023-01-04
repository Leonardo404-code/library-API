package library

import (
	"net/http"
)

type Handlers interface {
	GetBooks(http.ResponseWriter, *http.Request)
}
