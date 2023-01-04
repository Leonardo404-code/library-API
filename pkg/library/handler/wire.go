//go:build wireinject
// +build wireinject

package handler

import (
	"library-api/internal/database"
	"library-api/pkg/library"
	libRepo "library-api/pkg/library/repository"
	libSvc "library-api/pkg/library/service"

	"github.com/google/wire"
)

func Build() library.Handlers {
	wire.Build(
		Must,
		libSvc.Must,
		libRepo.Must,
		database.Connect,
	)
	return nil
}
