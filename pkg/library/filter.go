package library

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filter struct {
	BookID   primitive.ObjectID
	BookName string
}

func (f *Filter) GenerateQuery() bson.D {
	query := bson.D{}

	if !f.BookID.IsZero() {
		query = append(query, primitive.E{
			Key:   "_id",
			Value: f.BookID,
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
	if _, err := primitive.ObjectIDFromHex(f.BookID.Hex()); err != nil {
		return err
	}

	return nil
}
