//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/pkg"
)

//go:generate wire gen

// 移除具有 side effect 的全域變數
// global variable 是 testing 的萬惡之首
// 所有元件都要用注入的方式建構

func NewGin(cfg *configs.Config) *gin.Engine {
	panic(
		wire.Build(
			InfrastructureLayer,
			ApplicationLayer,
			HttpAdapterLayer,
			pkg.NewGin,
		),
	)
}
