package main

import (
	"rim-server/internal/app/rim/api"
	"rim-server/internal/app/rim/imageservice"
	"rim-server/internal/app/rim/model"
)

func main() {
	var err error
	err = model.Connect()
	err = imageservice.Connect()
	if err != nil {
		panic(err)
	}

	go imageservice.Start()
	api.Start()
}
