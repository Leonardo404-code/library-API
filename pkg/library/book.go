package library

import (
	"time"
)

type Book struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Writer      string    `json:"writer" bson:"writer"`
	ReleaseDate string    `json:"release_date" bson:"release_date"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" bson:"deleted_at"`
}
