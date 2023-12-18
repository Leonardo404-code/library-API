package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"library-api/pkg/library"
	"library-api/pkg/library/repository"
)

func Test_service_DeleteBook(t *testing.T) {
	type (
		fields struct {
			libraryRepo func() library.Repository
		}

		args struct {
			reqParams *library.Filter
		}
	)

	objectID := primitive.NewObjectID()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should return a invalid ID error",
			fields: fields{
				libraryRepo: func() library.Repository {
					m := &repository.Mock{}
					return m
				},
			},
			args: args{
				reqParams: &library.Filter{
					BookID: "12345667",
				},
			},
			wantErr: true,
		},
		{
			name: "should delete a book",
			fields: fields{
				libraryRepo: func() library.Repository {
					m := &repository.Mock{}
					m.On("Search", mock.Anything).Return([]*library.Book{
						{
							ID: objectID,
						},
					}, nil)
					m.On("DeleteBook", mock.Anything, mock.Anything).Return(nil)

					return m
				},
			},
			args: args{
				reqParams: &library.Filter{
					BookID: objectID.Hex(),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				libraryRepo: tt.fields.libraryRepo(),
			}

			if err := s.DeleteBook(tt.args.reqParams); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
