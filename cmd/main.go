package main

import (
	"github.com/JamshedJ/WalletAPI/config"
	"github.com/JamshedJ/WalletAPI/delivery/api"
	"github.com/JamshedJ/WalletAPI/infrastructure/glog"
	"github.com/JamshedJ/WalletAPI/infrastructure/repository/gormRepo"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic("cannot initialize config: " + err.Error())
	}
}

func main() {
	logger := glog.NewLogger()

	err := gormRepo.InitDatabase(config.Get().App.Database.Dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot initialize database")
	}

	err = gormRepo.AutoMigrate()
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot migrate database")
	}

	api.Run(nil, config.Get().App.Port)
}
