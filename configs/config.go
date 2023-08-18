package configs

import (
	"bytes"
	"embed"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//go:embed *
var directory embed.FS

type Config struct {
	Port        string
	LogLevel    string
	EnableHttp2 bool

	DebugHttp     bool
	DebugDatabase bool
}

// NewConfigFromFilename 為了在寫測試的時候, 可以方便抽換設定值
// 原本的寫法寫死相對路徑, 當進行測試時 呼叫此函數的時候會失敗
func NewConfigFromFilename(filename string) *Config {
	err := os.Setenv("TZ", "UTC")
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: set time zone to utc: %w", err))
	}

	buf, err := directory.ReadFile(filename + ".env")
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: read config file: %w", err))
	}

	viper.AutomaticEnv()
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	err = viper.ReadConfig(bytes.NewReader(buf))
	if err != nil {
		panic(fmt.Sprintf("讀取設定檔出現錯誤，原因為：%v", err))
	}
	return &Config{
		Port:          viper.GetString("port"),
		LogLevel:      viper.GetString("log_level"),
		EnableHttp2:   viper.GetBool("enable_http2"),
		DebugHttp:     viper.GetBool("debug_http"),
		DebugDatabase: viper.GetBool("debug_database"),
	}
}
