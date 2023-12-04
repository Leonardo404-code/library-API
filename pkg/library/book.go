package library

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `json:"id"           bson:"_id"`
	Title       string             `json:"title"        bson:"title"`
	Description string             `json:"description"  bson:"description"`
	Writer      string             `json:"writer"       bson:"writer"`
	ReleaseDate string             `json:"release_date" bson:"release_date"`
	Gender      string             `json:"gender"       bson:"gender"`
	CreatedAt   time.Time          `json:"created_at"   bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"   bson:"updated_at"`
	DeletedAt   time.Time          `json:"deleted_at"   bson:"deleted_at"`
}
