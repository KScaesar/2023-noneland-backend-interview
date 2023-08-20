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
	Infrastructure = wire.NewSet(
		pkg.NewSqliteGorm,
		pkg.NewHttpClient,
	)
	Application = wire.NewSet(
		wire.Struct(new(app.ApplicationGroup), "*"),

		database.NewGormTransactionBackupRepository,
		wire.Bind(new(entity.TransactionBackupRepository), new(*database.GormTransactionBackupRepository)),
		app.NewTransactionBackupUseCase,

		external.NewHttpExchangeQryService,
		wire.Bind(new(app.ExchangeQryService), new(*external.HttpExchangeQryService)),

		database.NewUserRepository,
		wire.Bind(new(entity.UserRepository), new(*database.UserRepository)),
		app.NewUserUseCase,
	)
	HttpAdapter = wire.NewSet(
		wire.Struct(new(api.HandlerGroup), "*"),

		api.NewExchangeHandler,
		api.NewUserHandler,
	)
)
