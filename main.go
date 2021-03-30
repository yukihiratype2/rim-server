package main

import (
	"rim-server/internal/app/rim/api"
	"rim-server/internal/app/rim/event"
	"rim-server/internal/app/rim/imageservice"
	"rim-server/internal/app/rim/model"
	"rim-server/internal/app/rim/s3"

	"github.com/spf13/viper"
)

func main() {
	checkConfig()
	var err error
	err = model.Connect()
	if err != nil {
		panic(err)
	}
	err = s3.Connect()
	if err != nil {
		panic(err)
	}
	event.Start()

	go imageservice.Start()
	api.Start()
}

func checkConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("dsn", "host=localhost user=natsuki dbname=natsuki port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	viper.SetDefault("s3", map[string]string{"endpoint": "127.0.0.1:9000", "accessKeyID": "minioadmin", "secretAccessKey": "minioadmin"})
	viper.ReadInConfig()
}
