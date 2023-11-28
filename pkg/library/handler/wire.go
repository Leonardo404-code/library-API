//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/google/wire"

	"library-api/internal/database"
	storagePkg "library-api/pkg/google"
	"library-api/pkg/library"
	libRepo "library-api/pkg/library/repository"
	libSvc "library-api/pkg/library/service"
)

func Build() library.Handlers {
	wire.Build(
		Must,
		libSvc.Must,
		libRepo.Must,
		database.Connect,
		storagePkg.Must,
	)
	return nil
}
