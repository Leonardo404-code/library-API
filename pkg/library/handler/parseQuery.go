package handler

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"library-api/pkg/library"
)

func parseQueryParams(requestQuery url.Values) *library.Filter {
	filter := new(library.Filter)

	for key, value := range requestQuery {
		switch key {
		case "book_name":
			filter.BookName = value[0]
		case "book_id":
			objectID, err := primitive.ObjectIDFromHex(value[0])
			if err != nil {
				return &library.Filter{
					BookID:   primitive.ObjectID{},
					BookName: "",
				}
			}
			filter.BookID = objectID
		default:
			return &library.Filter{
				BookID:   primitive.ObjectID{},
				BookName: "",
			}
		}
	}

	return filter
}
