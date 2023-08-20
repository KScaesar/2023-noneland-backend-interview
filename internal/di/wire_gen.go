// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"net/http"
	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/api"
	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/internal/database"
	"noneland/backend/interview/internal/external"
	"noneland/backend/interview/pkg"
)

// Injectors from wire.go:

// NewServer 寫法1
func NewServerV1(cfg *configs.Config) *http.Server {
	applicationGroup := NewApplication(cfg)
	handlerGroup := NewHttpHandler(applicationGroup)
	server := api.NewServer(cfg, handlerGroup)
	return server
}

// NewServer 寫法2
func NewServerV2(cfg *configs.Config) *http.Server {
	db := pkg.NewSqliteGorm()
	gormTransactionBackupRepository := database.NewGormTransactionBackupRepository(db)
	client := pkg.NewHttpClient()
	httpExchangeQryService := external.NewHttpExchangeQryService(client, cfg)
	transactionBackupUseCase := app.NewTransactionBackupUseCase(gormTransactionBackupRepository, httpExchangeQryService)
	userRepository := database.NewUserRepository(db, cfg)
	userUseCase := app.NewUserUseCase(userRepository)
	applicationGroup := &app.ApplicationGroup{
		TransactionBackupUseCase: transactionBackupUseCase,
		ExchangeQryService:       httpExchangeQryService,
		UserUseCase:              userUseCase,
	}
	exchangeHandler := api.NewExchangeHandler(applicationGroup)
	userHandler := api.NewUserHandler(applicationGroup)
	handlerGroup := api.HandlerGroup{
		ExchangeHandler: exchangeHandler,
		UserHandler:     userHandler,
	}
	server := api.NewServer(cfg, handlerGroup)
	return server
}

func NewApplication(cfg *configs.Config) *app.ApplicationGroup {
	db := pkg.NewSqliteGorm()
	gormTransactionBackupRepository := database.NewGormTransactionBackupRepository(db)
	client := pkg.NewHttpClient()
	httpExchangeQryService := external.NewHttpExchangeQryService(client, cfg)
	transactionBackupUseCase := app.NewTransactionBackupUseCase(gormTransactionBackupRepository, httpExchangeQryService)
	userRepository := database.NewUserRepository(db, cfg)
	userUseCase := app.NewUserUseCase(userRepository)
	applicationGroup := &app.ApplicationGroup{
		TransactionBackupUseCase: transactionBackupUseCase,
		ExchangeQryService:       httpExchangeQryService,
		UserUseCase:              userUseCase,
	}
	return applicationGroup
}

func NewHttpHandler(apps *app.ApplicationGroup) api.HandlerGroup {
	exchangeHandler := api.NewExchangeHandler(apps)
	userHandler := api.NewUserHandler(apps)
	handlerGroup := api.HandlerGroup{
		ExchangeHandler: exchangeHandler,
		UserHandler:     userHandler,
	}
	return handlerGroup
}
