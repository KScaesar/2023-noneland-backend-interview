package di

import (
	"github.com/google/wire"

	"noneland/backend/interview/internal/api"
	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/internal/pkg"
	"noneland/backend/interview/internal/repo/xGorm"
)

// gorm 不想重複名稱, 避免讀程式碼的時候誤會
// 所以改名 xGorm

var InfrastructureLayer = wire.NewSet(
	pkg.NewSqliteGorm,
)

var ApplicationLayer = wire.NewSet(
	xGorm.NewUserRepository,
	wire.Bind(new(entity.UserRepository), new(*xGorm.UserRepository)),
	app.NewUserUseCase,
)

var HttpAdapterLayer = wire.NewSet(
	api.NewUserHandler,
)
