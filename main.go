package main

import (
	"rim-server/internal/app/rim/api"
	"rim-server/internal/app/rim/event"
	"rim-server/internal/app/rim/imageservice"
	"rim-server/internal/app/rim/model"
	"rim-server/internal/app/rim/s3"
)

func main() {
	var err error
	err = model.Connect()
	if err != nil {
		panic(err)
	}
	err = s3.Connect()
	event.Start()
	if err != nil {
		panic(err)
	}

	go imageservice.Start()
	api.Start()
}
