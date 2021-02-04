package imageservice

import (
	"context"
	"image"
	"io"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var S3 *minio.Client

func Connect() (err error) {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"

	S3, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	return
}

func Start() {
	for n := range S3.ListenBucketNotification(context.Background(), "test-img", "", "", []string{"s3:ObjectCreated:*"}) {
		if n.Err != nil {
			panic(n.Err)
		}
		fileName := n.Records[0].S3.Object.Key
		fetchImage(fileName)
	}
}

func fetchImage(objetKey string) {
	object, err := S3.GetObject(context.Background(), "test-img", objetKey, minio.GetObjectOptions{})
	img, err := imaging.Decode(object)
	if err != nil {
		panic(err)
	}
	r, w := io.Pipe()
	go cropImage(img, w)
	saveImage(objetKey, r)
}

func cropImage(img image.Image, destWriter *io.PipeWriter) {
	result := imaging.Resize(img, 128, 128, imaging.Lanczos)
	imaging.Encode(destWriter, result, imaging.JPEG)
	destWriter.Close()
}

func saveImage(name string, r io.Reader) {
	S3.PutObject(context.Background(), "test-img-thumbnail", name, r, -1, minio.PutObjectOptions{})
}
