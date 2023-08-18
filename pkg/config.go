package pkg

import (
	"bytes"
	"embed"
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func NewConfigFromFilename[T any](dir embed.FS, filename string) *T {
	err := os.Setenv("TZ", "UTC")
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: set time zone to utc: %w", err))
	}

	buf, err := dir.ReadFile(filename + ".env")
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: read config file: %w", err))
	}

	cfg := new(T)

	viper.AutomaticEnv()
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	err = viper.ReadConfig(bytes.NewReader(buf))
	if err != nil {
		panic(fmt.Sprintf("讀取設定檔出現錯誤，原因為： %v", err))
		return nil
	}

	option := func(c *mapstructure.DecoderConfig) { c.TagName = "configs" }
	err = viper.Unmarshal(cfg, option)
	if err != nil {
		panic(fmt.Sprintf("viper Unmarshal failed: %v", err))
		return nil
	}

	return cfg
}
