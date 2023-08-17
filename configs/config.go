package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
	Mode string
}

// NewConfigFromFilename 為了在寫測試的時候, 可以方便抽換設定值
func NewConfigFromFilename(filename string) *Config {
	err := os.Setenv("TZ", "UTC")
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: set time zone to utc: %w", err))
	}

	viper.AutomaticEnv()
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("讀取設定檔出現錯誤，原因為：%v", err))
	}
	return &Config{
		Port: viper.GetString("port"),
		Mode: viper.GetString("mode"),
	}
}

func NewConfig() *Config {
	// filename := os.Getenv("ENV")
	filename := "template-dev"
	return NewConfigFromFilename(filename)
}
