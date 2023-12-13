package library

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filter struct {
	BookID   string
	BookName string
}

func (f *Filter) GenerateQuery() bson.D {
	query := bson.D{}

	if f.BookID != "" {
		ObjectID, _ := primitive.ObjectIDFromHex(f.BookID)

		query = append(query, primitive.E{
			Key:   "_id",
			Value: ObjectID,
		})
	}

	if f.BookName != "" {
		query = append(query, primitive.E{
			Key:   "name",
			Value: f.BookName,
		})
	}

	return query
}

func (f *Filter) ValidateParams() error {
	if _, err := primitive.ObjectIDFromHex(f.BookID); err != nil {
		return fmt.Errorf("invalid book ID")
	}

	return nil
}
