package di

import (
	"github.com/google/wire"

	"noneland/backend/interview/internal/api"
	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/internal/database"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/internal/external"
	"noneland/backend/interview/pkg"
)

var (
	InfrastructureLayer = wire.NewSet(
		pkg.NewSqliteGorm,
		pkg.NewHttpClient,
	)
	ApplicationLayer = wire.NewSet(
		database.NewGormTransactionBackupRepository,
		wire.Bind(new(entity.TransactionBackupRepository), new(*database.GormTransactionBackupRepository)),
		app.NewTransactionBackupUseCase,

		external.NewHttpExchangeQryService,
		wire.Bind(new(app.ExchangeQryService), new(*external.HttpExchangeQryService)),

		database.NewUserRepository,
		wire.Bind(new(entity.UserRepository), new(*database.UserRepository)),
		app.NewUserUseCase,
	)
	HttpAdapterLayer = wire.NewSet(
		api.NewExchangeHandler,
		api.NewUserHandler,
		wire.Struct(new(api.HandlerGroup), "*"),
	)
)
