package main

import (
	"rim-server/internal/app/rim/api"
	"rim-server/internal/app/rim/imageservice"
	"rim-server/internal/app/rim/model"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	model, err := model.NewModel()
	if err != nil {
		panic(err)
	}
	server := api.New(model, minioClient)

	is := imageservice.NewImageService(minioClient)
	go is.Start()
	server.Start()
}
