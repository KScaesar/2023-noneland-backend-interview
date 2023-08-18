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
	"noneland/backend/interview/pkg"
)

// Injectors from wire.go:

func NewServer(cfg *configs.Config) *http.Server {
	db := pkg.NewSqliteGorm()
	userRepository := database.NewUserRepository(db, cfg)
	userUseCase := app.NewUserUseCase(userRepository)
	userHandler := api.NewUserHandler(userUseCase)
	handlerGroup := api.HandlerGroup{
		UserH: userHandler,
	}
	server := api.NewServer(cfg, handlerGroup)
	return server
}
