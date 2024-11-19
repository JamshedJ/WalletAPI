package main

import (
	"github.com/JamshedJ/WalletAPI/config"
	"github.com/JamshedJ/WalletAPI/delivery/api"
	"github.com/JamshedJ/WalletAPI/infrastructure/glog"
)
func init() {
	if err := config.InitConfig(); err != nil {
		panic("cannot initialize config: " + err.Error())
	}
}
func main() {
	logger := glog.NewLogger()
	logger.Info().Msg("starting application")

	api.Run(nil, config.Get().App.Port)
}
