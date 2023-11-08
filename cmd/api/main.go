package main

import (
	"github.com/spf13/viper"
	"go-learning-demo/config"
	"go-learning-demo/server"
	"log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Println(err.Error())
	}
}
