package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DBConfig struct {
	MinConns int32
	MaxConns int32
	URL      string
}

type Config struct {
	AppEnv string
	DB     DBConfig
}

func NewConfig() (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Overload(".env")
	}

	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("app.env", "development")

	v.SetDefault("db.url", "")
	v.SetDefault("db.minconns", 1)
	v.SetDefault("db.maxconns", 10)

	cfg := &Config{
		AppEnv: v.GetString("app.env"),
		DB: DBConfig{
			MinConns: int32(v.GetInt("db.minconns")),
			MaxConns: int32(v.GetInt("db.maxconns")),
		},
	}

	if urlFromEnv := v.GetString("db.url"); urlFromEnv != "" {
		cfg.DB.URL = urlFromEnv
	} else {
		panic(fmt.Errorf("db.url is not set"))
	}

	return cfg, nil
}
