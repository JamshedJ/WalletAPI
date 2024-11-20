package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config *Config

func Get() *Config {
	return config
}

type Config struct {
	App struct {
		Port        int
		Environment string
		SecretKey   string
		Database    struct {
			Dsn string
		}
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot initialize config: %v", err)
	}
	viper.Unmarshal(&config)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&config); err != nil {
			log.Default().Printf("cannot unmarshal config: %v\n", err)
		}
	})

	return nil
}
