package handler

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type BookResponseDoc struct {
	ID              uuid.UUID `example:"36e4379a-be55-462d-ab9e-da0c532623ea"`
	Title           string    `json:"title" example:"Test title"`
	AuthorName      string    `json:"author_name" example:"Jhonn Doe"`
	Publisher       string    `json:"publisher" example:"Brazil Journal"`
	PublicationYear int32     `json:"publication_year" example:"2023"`
	Genre           string    `json:"genre" example:"notice"`
}

type BookRequestDoc struct {
	Title           string `json:"title" example:"Test title"`
	AuthorName      string `json:"author_name" example:"Jhonn Doe"`
	Publisher       string `json:"publisher" example:"Brazil Journal"`
	PublicationYear int32  `json:"publication_year" example:"2023"`
	Genre           string `json:"genre" example:"notice"`
}
