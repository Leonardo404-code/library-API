package library

import (
	"time"
)

type Book struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Writer      string    `json:"writer"`
	ReleaseDate string    `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
