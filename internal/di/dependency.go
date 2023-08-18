package di

import (
	"github.com/google/wire"

	"noneland/backend/interview/internal/api"
	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/internal/database"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/pkg"
)

var (
	InfrastructureLayer = wire.NewSet(
		pkg.NewSqliteGorm,
	)
	ApplicationLayer = wire.NewSet(
		database.NewUserRepository,
		wire.Bind(new(entity.UserRepository), new(*database.UserRepository)),
		app.NewUserUseCase,
	)
	HttpAdapterLayer = wire.NewSet(
		api.NewUserHandler,
		wire.Struct(new(api.HandlerGroup), "*"),
	)
)
