package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var Client *minio.Client

func Connect() (err error) {
	endpoint := viper.GetString("s3.endpoint")
	accessKeyID := viper.GetString("s3.accessKeyID")
	secretAccessKey := viper.GetString("s3.secretAccessKey")

	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	return
}
