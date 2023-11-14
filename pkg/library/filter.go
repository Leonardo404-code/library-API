package library

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filter struct {
	BookID   string
	BookName string
}

func (f *Filter) GenerateQuery() (bson.D, error) {
	query := bson.D{}

	if f.BookID != "" {
		if !primitive.IsValidObjectID(f.BookID) {
			return nil, errors.New("invalid ID")
		}
		query = append(query, primitive.E{
			Key:   "_id",
			Value: fmt.Sprintf("ObjectId(%s)", f.BookID),
		})
	}

	if f.BookName != "" {
		query = append(query, primitive.E{
			Key:   "name",
			Value: f.BookName,
		})
	}

	return query, nil
}
