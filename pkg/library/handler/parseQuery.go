package handler

import (
	"net/url"

	"library-api/pkg/library"
)

func parseQueryParams(requestQuery url.Values) *library.Filter {
	reqParams := new(library.Filter)

	for key, value := range requestQuery {
		switch key {
		case "book_name":
			reqParams.BookName = value[0]
		case "book_id":
			reqParams.BookID = value[0]
		default:
			return &library.Filter{
				BookID:   "",
				BookName: "",
			}
		}
	}

	return reqParams
}
