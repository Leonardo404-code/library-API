package service

import (
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"library-api/pkg/library"
	"library-api/pkg/library/repository"
)

func Test_service_CreateBook(t *testing.T) {
	type (
		fields struct {
			libraryRepo func() library.Repository
		}

		args struct {
			book     *library.Book
			bookFile multipart.File
		}
	)

	objectID := primitive.NewObjectID()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *library.Book
		wantErr bool
	}{
		{
			name: "should return a title lenght error",
			fields: fields{
				libraryRepo: func() library.Repository {
					m := &repository.Mock{}
					return m
				},
			},
			args: args{
				book: &library.Book{
					Title: "Ti",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a description lenght error",
			fields: fields{
				libraryRepo: func() library.Repository {
					m := &repository.Mock{}
					return m
				},
			},
			args: args{
				book: &library.Book{
					Title:       "Title",
					Description: "Descri",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a time parse error",
			fields: fields{
				libraryRepo: func() library.Repository {
					m := &repository.Mock{}
					return m
				},
			},
			args: args{
				book: &library.Book{
					Title:       "Title",
					Description: "description",
					ReleaseDate: "01-12-2004",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should create a book",
			fields: fields{
				libraryRepo: func() library.Repository {
					m := &repository.Mock{}
					m.On("CreateBook", mock.Anything, mock.Anything, mock.Anything).Return(nil)

					return m
				},
			},
			args: args{
				book: &library.Book{
					ID:          objectID,
					Title:       "Title",
					Description: "description",
					ReleaseDate: "2023-01-01",
				},
			},
			want: &library.Book{
				ID: objectID,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				libraryRepo: tt.fields.libraryRepo(),
			}

			if _, err := s.CreateBook(tt.args.book, tt.args.bookFile); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
