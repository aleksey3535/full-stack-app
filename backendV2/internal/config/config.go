package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Env      string `default:"local"`
	Port     string
	Login    string
	Password string
	Db       dbConf
}

type dbConf struct {
	Dbname   string
	Host     string
	Port     string
	Name     string
	Password string
	Sslmode  string
}

func MustLoad() *Config {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("error occurred loading .env file: %s", err.Error()))
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("invalid config path")
	}
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("error occurred reading config file: %s", err.Error()))
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Sprintf("error occurred unmarshaling config file:%s", err.Error()))
	}
	cfg.Login = os.Getenv("ADMIN_LOGIN")
	cfg.Password = os.Getenv("ADMIN_PASSWORD")
	cfg.Db.Password = os.Getenv("DB_PASSWORD")
	return &cfg
}
