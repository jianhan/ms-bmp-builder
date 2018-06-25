package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/micro/go-micro"
	"github.com/spf13/viper"
	"time"
)

func main() {
	// log .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
		panic(err)
	}

	// initialize new service
	srv := micro.NewService(
		micro.Name(viper.GetString("SERVICE_NAME")),
		micro.RegisterTTL(time.Duration(viper.GetInt("REGISTER_TTL_DURATION"))*time.Second),
		micro.RegisterInterval(time.Duration(viper.GetInt("REGISTER_INTERVAL"))*10),
		micro.Version(viper.GetString("SERVICE_VERSION")),
	)

	// init service
	srv.Init()
}
