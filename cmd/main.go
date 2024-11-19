package main

import "github.com/JamshedJ/WalletAPI/config"

func init() {
	if err := config.InitConfig(); err != nil {
		panic("cannot initialize config: " + err.Error())
	}
}
func main() {

}
