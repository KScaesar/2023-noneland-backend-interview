//go:build wireinject
// +build wireinject

package di

import (
	"net/http"

	"github.com/google/wire"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/api"
)

//go:generate wire gen

// 移除具有 side effect 的全域變數
// global variable 是 testing 的萬惡之首
// 所有元件都要用注入的方式建構

func NewServer(cfg *configs.Config) *http.Server {
	panic(
		wire.Build(
			Infrastructure,
			Application,
			HttpAdapter,
			api.NewServer,
		),
	)
}

// 範例: 串接不同的驅動方式
// func NewCommandRoot(cfg *configs.Config) *cobra.Command {
// 	panic(
// 		wire.Build(
// 			Infrastructure,
// 			Application,
// 			CmdAdapter,
// 			cmd.NewCommandRoot,
// 		),
// 	)
// }
