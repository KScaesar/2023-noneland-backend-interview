package configs

import (
	"embed"

	"noneland/backend/interview/pkg"
)

//go:embed *
var directory embed.FS

type Config struct {
	Port        string `configs:"port"`
	LogLevel    string `configs:"log_level"`
	EnableHttp2 bool   `configs:"enable_http2"`

	DebugHttp     bool `configs:"debug_http"`
	DebugDatabase bool `configs:"debug_database"`

	ExchangeUrl string `configs:"exchange_url"`
}

// NewConfig 為了在寫測試的時候, 可以方便抽換設定值
// 原本的寫法寫死相對路徑, 當進行測試時 呼叫此函數的時候會失敗
func NewConfig(filename string) *Config {
	return pkg.NewConfigFromFilename[Config](directory, filename)
}
